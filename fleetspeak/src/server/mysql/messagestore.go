// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mysql

import (
	"context"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	log "github.com/golang/glog"
	tspb "google.golang.org/protobuf/types/known/timestamppb"

	"github.com/google/fleetspeak/fleetspeak/src/common"
	"github.com/google/fleetspeak/fleetspeak/src/server/db"

	fspb "github.com/google/fleetspeak/fleetspeak/src/common/proto/fleetspeak"
	"google.golang.org/protobuf/proto"
	anypb "google.golang.org/protobuf/types/known/anypb"
)

// dbMessage matches the schema of the messages table, optionally joined to the
// pending_messages table.
type dbMessage struct {
	messageID              []byte
	sourceClientID         []byte
	sourceServiceName      string
	sourceMessageID        []byte
	destinationClientID    []byte
	destinationServiceName string
	messageType            string
	creationTimeSeconds    int64
	creationTimeNanos      int32
	processedTimeSeconds   sql.NullInt64
	processedTimeNanos     sql.NullInt64
	validationInfo         []byte
	failed                 sql.NullBool
	failedReason           sql.NullString
	retryCount             uint32
	dataTypeURL            sql.NullString
	dataValue              []byte
	annotations            []byte
}

func toMicro(t time.Time) int64 {
	return t.UnixNano() / 1000
}

func (d *Datastore) SetMessageResult(ctx context.Context, dest common.ClientID, id common.MessageID, res *fspb.MessageResult) error {
	return d.runInTx(ctx, false, func(tx *sql.Tx) error { return d.trySetMessageResult(ctx, tx, dest.IsNil(), id, res) })
}

func (d *Datastore) trySetMessageResult(ctx context.Context, tx *sql.Tx, forServer bool, id common.MessageID, res *fspb.MessageResult) error {
	dbm := dbMessage{
		messageID:            id.Bytes(),
		processedTimeSeconds: sql.NullInt64{Valid: true, Int64: res.ProcessedTime.Seconds},
		processedTimeNanos:   sql.NullInt64{Valid: true, Int64: int64(res.ProcessedTime.Nanos)},
	}
	if res.Failed {
		dbm.failed = sql.NullBool{Valid: true, Bool: true}
		dbm.failedReason = sql.NullString{Valid: true, String: res.FailedReason}
	}
	_, err := tx.ExecContext(ctx, "UPDATE messages SET failed=?, failed_reason=?, processed_time_seconds=?, processed_time_nanos=? WHERE message_id=?",
		dbm.failed, dbm.failedReason, dbm.processedTimeSeconds, dbm.processedTimeNanos, dbm.messageID)
	if err != nil {
		return err
	}
	var fs int
	if forServer {
		fs = 1
	}
	_, err = tx.ExecContext(ctx, "DELETE FROM pending_messages WHERE for_server=? AND message_id=?", fs, dbm.messageID)
	return err
}

func fromNULLString(s sql.NullString) string {
	if !s.Valid {
		return ""
	}
	return s.String
}

func fromMessageProto(m *fspb.Message) (*dbMessage, error) {
	id, err := common.BytesToMessageID(m.MessageId)
	if err != nil {
		return nil, err
	}
	dbm := &dbMessage{
		messageID:   id.Bytes(),
		messageType: m.MessageType,
	}
	if m.Source != nil {
		dbm.sourceClientID = m.Source.ClientId
		dbm.sourceServiceName = m.Source.ServiceName
	}
	if m.Destination != nil {
		dbm.destinationClientID = m.Destination.ClientId
		dbm.destinationServiceName = m.Destination.ServiceName
	}
	if len(m.SourceMessageId) != 0 {
		dbm.sourceMessageID = m.SourceMessageId
	}
	if m.CreationTime != nil {
		dbm.creationTimeSeconds = m.CreationTime.Seconds
		dbm.creationTimeNanos = m.CreationTime.Nanos
	}
	if m.Result != nil {
		r := m.Result
		if r.ProcessedTime != nil {
			dbm.processedTimeSeconds = sql.NullInt64{Int64: r.ProcessedTime.Seconds, Valid: true}
			dbm.processedTimeNanos = sql.NullInt64{Int64: int64(r.ProcessedTime.Nanos), Valid: true}
		}
		if r.Failed {
			dbm.failed = sql.NullBool{Bool: true, Valid: true}
			dbm.failedReason = sql.NullString{String: r.FailedReason, Valid: true}
		}
	}
	if m.Data != nil {
		dbm.dataTypeURL = sql.NullString{String: m.Data.TypeUrl, Valid: true}
		dbm.dataValue = m.Data.Value
	}
	if m.ValidationInfo != nil {
		b, err := proto.Marshal(m.ValidationInfo)
		if err != nil {
			return nil, err
		}
		dbm.validationInfo = b
	}
	if m.Annotations != nil {
		b, err := proto.Marshal(m.Annotations)
		if err != nil {
			return nil, err
		}
		dbm.annotations = b
	}
	return dbm, nil
}

func toMessageResultProto(m *dbMessage) *fspb.MessageResult {
	if !m.processedTimeSeconds.Valid {
		return nil
	}

	ret := &fspb.MessageResult{
		ProcessedTime: &tspb.Timestamp{
			Seconds: m.processedTimeSeconds.Int64,
			Nanos:   int32(m.processedTimeNanos.Int64)},
		Failed: m.failed.Valid && m.failed.Bool,
	}

	if m.failedReason.Valid {
		ret.FailedReason = m.failedReason.String
	}
	return ret
}

func toMessageProto(m *dbMessage) (*fspb.Message, error) {
	mid, err := common.BytesToMessageID(m.messageID)
	if err != nil {
		return nil, err
	}
	pm := &fspb.Message{
		MessageId: mid.Bytes(),
		Source: &fspb.Address{
			ClientId:    m.sourceClientID,
			ServiceName: m.sourceServiceName,
		},
		SourceMessageId: m.sourceMessageID,
		Destination: &fspb.Address{
			ClientId:    m.destinationClientID,
			ServiceName: m.destinationServiceName,
		},
		MessageType: m.messageType,
		CreationTime: &tspb.Timestamp{
			Seconds: m.creationTimeSeconds,
			Nanos:   m.creationTimeNanos,
		},
		Result: toMessageResultProto(m),
	}
	if m.dataTypeURL.Valid {
		pm.Data = &anypb.Any{
			TypeUrl: m.dataTypeURL.String,
			Value:   m.dataValue,
		}
	}
	if len(m.validationInfo) > 0 {
		v := &fspb.ValidationInfo{}
		if err := proto.Unmarshal(m.validationInfo, v); err != nil {
			return nil, err
		}
		pm.ValidationInfo = v
	}
	if len(m.annotations) > 0 {
		a := &fspb.Annotations{}
		if err := proto.Unmarshal(m.annotations, a); err != nil {
			return nil, err
		}
		pm.Annotations = a
	}
	return pm, nil
}

func genPlaceholders(num int) string {
	es := make([]string, num)
	for i := range es {
		es[i] = "?"
	}
	return strings.Join(es, ", ")
}

func (d *Datastore) getPendingMessageRawIds(ctx context.Context, tx *sql.Tx, ids []common.ClientID, offset uint64, limit uint64, forUpdate bool) ([][]byte, error) {
	squery := fmt.Sprintf("SELECT "+
		"m.message_id AS message_id "+
		"FROM messages AS m, pending_messages AS pm "+
		"WHERE m.destination_client_id IN (%s) AND m.message_id=pm.message_id "+
		"ORDER BY message_id",
		genPlaceholders((len(ids))))

	if offset != 0 && limit == 0 {
		return nil, fmt.Errorf("if offset is provided, a limit must be provided as well")
	}

	if limit != 0 {
		squery += " LIMIT ?"
	}

	if offset != 0 {
		squery += " OFFSET ?"
	}

	if forUpdate {
		squery += " FOR UPDATE"
	}

	args := make([]interface{}, len(ids), len(ids)+2)
	for i, v := range ids {
		args[i] = v.Bytes()
	}

	if limit != 0 {
		args = append(args, limit)
	}

	if offset != 0 {
		args = append(args, offset)
	}

	idsToProc := make([][]byte, 0)

	rs, err := tx.QueryContext(ctx, squery, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch the list of pending messages: %v", err)
	}
	defer rs.Close()
	for rs.Next() {
		var id []byte
		if err := rs.Scan(&id); err != nil {
			return nil, err
		}
		idsToProc = append(idsToProc, id)
	}

	return idsToProc, nil
}

func (d *Datastore) GetPendingMessageCount(ctx context.Context, ids []common.ClientID) (uint64, error) {
	var result uint64

	err := d.runInTx(ctx, true, func(tx *sql.Tx) error {
		squery := fmt.Sprintf("SELECT "+
			"COUNT(*) "+
			"FROM messages AS m, pending_messages AS pm "+
			"WHERE m.destination_client_id IN (%s) AND m.message_id=pm.message_id ",
			genPlaceholders((len(ids))))

		args := make([]interface{}, len(ids))
		for i, v := range ids {
			args[i] = v.Bytes()
		}
		rs, err := tx.QueryContext(ctx, squery, args...)
		if err != nil {
			return fmt.Errorf("Failed to fetch the pending message count: %v", err)
		}
		defer rs.Close()
		if !rs.Next() {
			return fmt.Errorf("Got empty result")
		}
		err = rs.Scan(&result)
		if err != nil {
			return fmt.Errorf("Failed to scan result: %v", err)
		}
		return nil
	})

	return result, err
}

func (d *Datastore) GetPendingMessages(ctx context.Context, ids []common.ClientID, offset uint64, count uint64, wantData bool) ([]*fspb.Message, error) {
	var res []*fspb.Message
	err := d.runInTx(ctx, true, func(tx *sql.Tx) error {
		messageIdsRaw, err := d.getPendingMessageRawIds(ctx, tx, ids, offset, count, false)
		if err != nil {
			return err
		}
		var messageIds []common.MessageID
		for _, idRaw := range messageIdsRaw {
			messageID, err := common.BytesToMessageID(idRaw)
			if err != nil {
				return err
			}
			messageIds = append(messageIds, messageID)
		}
		res, err = d.getMessages(ctx, tx, messageIds, wantData)
		return err
	})
	return res, err
}

func (d *Datastore) DeletePendingMessages(ctx context.Context, ids []common.ClientID) error {
	return d.runInTx(ctx, false, func(tx *sql.Tx) error {
		messageIds, err := d.getPendingMessageRawIds(ctx, tx, ids, 0, 0, true)
		if err != nil {
			return err
		}

		idsToProc := make([]interface{}, len(messageIds))
		for i, id := range messageIds {
			idsToProc[i] = id
		}

		// If there are no messages to be deleted, just bail out.
		if len(idsToProc) == 0 {
			return nil
		}

		now := db.NowProto()
		ptimeSecs := sql.NullInt64{Valid: true, Int64: now.Seconds}
		ptimeNanoSecs := sql.NullInt64{Valid: true, Int64: int64(now.Nanos)}
		failed := sql.NullBool{Valid: true, Bool: true}
		failedReason := sql.NullString{Valid: true, String: "Removed by admin action."}

		ps := genPlaceholders(len(idsToProc))
		uquery := fmt.Sprintf("UPDATE messages SET failed=?, failed_reason=?, processed_time_seconds=?, processed_time_nanos=? WHERE message_id IN (%s)", ps)
		_, err = tx.ExecContext(ctx, uquery, append([]interface{}{failed, failedReason, ptimeSecs, ptimeNanoSecs}, idsToProc...)...)
		if err != nil {
			return err
		}

		dquery := fmt.Sprintf("DELETE FROM pending_messages WHERE message_id IN (%s)", ps)
		_, err = tx.ExecContext(ctx, dquery, idsToProc...)

		return err
	})
}

func (d *Datastore) StoreMessages(ctx context.Context, msgs []*fspb.Message, contact db.ContactID) error {
	return d.runInTx(ctx, false, func(tx *sql.Tx) error {
		ids := make([][]byte, 0, len(msgs))
		for _, m := range msgs {
			dbm, err := fromMessageProto(m)
			if err != nil {
				return err
			}
			// If it is already processed, we don't want to save m.Data.
			if m.Result != nil {
				dbm.dataTypeURL = sql.NullString{Valid: false}
				dbm.dataValue = nil
			}
			ids = append(ids, dbm.messageID)
			if m.Result != nil && !m.Result.Failed {
				if err := d.tryStoreMessage(ctx, tx, dbm, false); err != nil {
					return err
				}
				if m.Result != nil {
					mid, _ := common.BytesToMessageID(m.MessageId)
					if err := d.trySetMessageResult(ctx, tx, m.Destination.ClientId == nil, mid, m.Result); err != nil {
						return err
					}
				}
				continue
			}
			var processedTime sql.NullInt64
			var failed sql.NullBool
			e := tx.QueryRowContext(ctx, "SELECT processed_time_seconds, failed FROM messages where message_id=?", dbm.messageID).Scan(&processedTime, &failed)
			switch {
			case e == sql.ErrNoRows:
				// Common case. Message not yet present, store as normal.
				if err := d.tryStoreMessage(ctx, tx, dbm, false); err != nil {
					return err
				}
			case e != nil:
				return e
			case processedTime.Valid && (!failed.Valid || !failed.Bool):
				// Message previously successfully processed, ignore this reprocessing.
			case m.Result != nil && (!processedTime.Valid || !m.Result.Failed):
				mid, err := common.BytesToMessageID(m.MessageId)
				if err != nil {
					return err
				}
				// Message not previously successfully processed, but this try succeeded. Mark as processed.
				if err := d.trySetMessageResult(ctx, tx, m.Destination.ClientId == nil, mid, m.Result); err != nil {
					return err
				}
			default:
				// The message is already present, but unprocessed/failed, and this
				// processing didn't succeed or is ongoing. Nothing to do.
			}
		}

		if contact == "" {
			return nil
		}

		c, err := strconv.ParseUint(string(contact), 16, 64)
		if err != nil {
			e := fmt.Errorf("unable to parse ContactID [%v]: %v", contact, err)
			log.Error(e)
			return e
		}
		for _, id := range ids {
			if _, err := tx.ExecContext(ctx, "INSERT IGNORE INTO client_contact_messages(client_contact_id, message_id) VALUES (?, ?)", c, id); err != nil {
				return err
			}
		}
		return nil
	})
}

func (d *Datastore) tryStoreMessage(ctx context.Context, tx *sql.Tx, dbm *dbMessage, isBroadcast bool) error {
	if dbm.creationTimeSeconds == 0 {
		return errors.New("message CreationTime must be set")
	}
	res, err := tx.ExecContext(ctx, "INSERT IGNORE INTO messages("+
		"message_id, "+
		"source_client_id, "+
		"source_service_name, "+
		"source_message_id, "+
		"destination_client_id, "+
		"destination_service_name, "+
		"message_type, "+
		"creation_time_seconds, "+
		"creation_time_nanos, "+
		"processed_time_seconds, "+
		"processed_time_nanos, "+
		"failed,"+
		"failed_reason,"+
		"validation_info,"+
		"annotations) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		dbm.messageID,
		dbm.sourceClientID,
		dbm.sourceServiceName,
		dbm.sourceMessageID,
		dbm.destinationClientID,
		dbm.destinationServiceName,
		dbm.messageType,
		dbm.creationTimeSeconds,
		dbm.creationTimeNanos,
		dbm.processedTimeSeconds,
		dbm.processedTimeNanos,
		dbm.failed,
		dbm.failedReason,
		dbm.validationInfo,
		dbm.annotations)
	if err != nil {
		return err
	}
	cnt, err := res.RowsAffected()
	if err != nil {
		return err
	}
	inserted := cnt == 1
	if inserted && !dbm.processedTimeSeconds.Valid {
		var due int64
		if dbm.destinationClientID == nil {
			due = toMicro(db.ServerRetryTime(0))
		} else {
			// If this is being created in response to a broadcast, then we about to
			// hand it to the client and should wait before providing in through
			// ClientMessagesForProcessing. Otherwise, we should give it to the client
			// on next contact.
			if isBroadcast {
				due = toMicro(db.ClientRetryTime())
			} else {
				due = toMicro(db.Now())
			}
		}
		fs := 0
		if dbm.destinationClientID == nil {
			fs = 1
		}
		_, err = tx.ExecContext(ctx, "INSERT INTO pending_messages("+
			"for_server, "+
			"message_id, "+
			"retry_count, "+
			"scheduled_time, "+
			"data_type_url, "+
			"data_value) VALUES(?, ?, ?, ?, ?, ?)",
			fs,
			dbm.messageID,
			0,
			due,
			dbm.dataTypeURL,
			dbm.dataValue)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Datastore) getMessages(ctx context.Context, tx *sql.Tx, ids []common.MessageID, wantData bool) ([]*fspb.Message, error) {
	res := make([]*fspb.Message, 0, len(ids))
	stmt1, err := tx.Prepare("SELECT " +
		"message_id, " +
		"source_client_id, " +
		"source_service_name, " +
		"source_message_id, " +
		"destination_client_id, " +
		"destination_service_name, " +
		"message_type, " +
		"creation_time_seconds, " +
		"creation_time_nanos, " +
		"processed_time_seconds, " +
		"processed_time_nanos, " +
		"validation_info, " +
		"annotations " +
		"FROM messages WHERE message_id=?")
	var stmt2 *sql.Stmt
	if wantData {
		stmt2, err = tx.Prepare("SELECT data_type_url, data_value FROM pending_messages WHERE message_id=?")
		if err != nil {
			return nil, err
		}
	}
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		row := stmt1.QueryRowContext(ctx, id.Bytes())
		var dbm dbMessage
		err := row.Scan(
			&dbm.messageID,
			&dbm.sourceClientID,
			&dbm.sourceServiceName,
			&dbm.sourceMessageID,
			&dbm.destinationClientID,
			&dbm.destinationServiceName,
			&dbm.messageType,
			&dbm.creationTimeSeconds,
			&dbm.creationTimeNanos,
			&dbm.processedTimeSeconds,
			&dbm.processedTimeNanos,
			&dbm.validationInfo,
			&dbm.annotations)
		if err != nil {
			return nil, err
		}
		if wantData {
			row := stmt2.QueryRowContext(ctx, id.Bytes())
			err := row.Scan(&dbm.dataTypeURL, &dbm.dataValue)
			if err != nil && err != sql.ErrNoRows {
				return nil, err
			}
		}
		m, err := toMessageProto(&dbm)
		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return res, nil
}

func (d *Datastore) GetMessages(ctx context.Context, ids []common.MessageID, wantData bool) ([]*fspb.Message, error) {
	var res []*fspb.Message
	err := d.runInTx(ctx, true, func(tx *sql.Tx) error {
		var err error
		res, err = d.getMessages(ctx, tx, ids, wantData)
		return err
	})
	return res, err
}

func (d *Datastore) GetMessageResult(ctx context.Context, id common.MessageID) (*fspb.MessageResult, error) {
	var ret *fspb.MessageResult

	err := d.runInTx(ctx, true, func(tx *sql.Tx) error {
		row := tx.QueryRowContext(ctx, "SELECT "+
			"creation_time_seconds, "+
			"creation_time_nanos, "+
			"processed_time_seconds, "+
			"processed_time_nanos, "+
			"failed, "+
			"failed_reason "+
			"FROM messages WHERE message_id=?", id.Bytes())

		var dbm dbMessage
		if err := row.Scan(
			&dbm.creationTimeSeconds,
			&dbm.creationTimeNanos,
			&dbm.processedTimeSeconds,
			&dbm.processedTimeNanos,
			&dbm.failed,
			&dbm.failedReason,
		); err == sql.ErrNoRows {
			return nil
		} else if err != nil {
			return err
		}

		ret = toMessageResultProto(&dbm)
		return nil
	})

	return ret, err
}

func (d *Datastore) ClientMessagesForProcessing(ctx context.Context, id common.ClientID, lim uint64, serviceLimits map[string]uint64) ([]*fspb.Message, error) {
	if id == (common.ClientID{}) {
		return nil, errors.New("a client is required")
	}
	return d.internalClientMessagesForProcessing(ctx, id, lim, serviceLimits)
}

type pendingUpdate struct {
	id  []byte
	nc  uint32
	due int64
}

func (d *Datastore) internalClientMessagesForProcessing(ctx context.Context, id common.ClientID, lim uint64, serviceLimits map[string]uint64) ([]*fspb.Message, error) {

	var read map[string]uint64
	if serviceLimits != nil {
		read = make(map[string]uint64)
	}
	res := make([]*fspb.Message, 0, lim)

	if err := d.runInTx(ctx, false, func(tx *sql.Tx) error {
		var updates []*pendingUpdate

		// As an internal addition to the MessageStore interface, this
		// also gets server messages when id=ClientID{}.
		rs, err := tx.QueryContext(ctx, "SELECT "+
			"m.message_id, "+
			"m.source_client_id, "+
			"m.source_service_name, "+
			"m.source_message_id, "+
			"m.destination_client_id, "+
			"m.destination_service_name, "+
			"m.message_type, "+
			"m.creation_time_seconds, "+
			"m.creation_time_nanos,"+
			"m.validation_info,"+
			"m.annotations,"+
			"pm.retry_count, "+
			"pm.data_type_url, "+
			"pm.data_value "+
			"FROM messages AS m, pending_messages AS pm "+
			"WHERE m.destination_client_id = ? AND m.message_id=pm.message_id AND pm.scheduled_time < ? "+
			"FOR UPDATE",
			id.Bytes(), toMicro(db.Now()))
		if err != nil {
			return err
		}
		defer rs.Close()
		for rs.Next() {
			var dbm dbMessage
			if err := rs.Scan(
				&dbm.messageID,
				&dbm.sourceClientID,
				&dbm.sourceServiceName,
				&dbm.sourceMessageID,
				&dbm.destinationClientID,
				&dbm.destinationServiceName,
				&dbm.messageType,
				&dbm.creationTimeSeconds,
				&dbm.creationTimeNanos,
				&dbm.validationInfo,
				&dbm.annotations,
				&dbm.retryCount,
				&dbm.dataTypeURL,
				&dbm.dataValue,
			); err != nil {
				return err
			}
			if serviceLimits != nil {
				if read[dbm.destinationServiceName] >= serviceLimits[dbm.destinationServiceName] {
					continue
				} else {
					read[dbm.destinationServiceName]++
				}
			}

			if len(res) >= int(lim) {
				continue
			}

			updates = append(updates, &pendingUpdate{
				id:  dbm.messageID,
				nc:  dbm.retryCount + 1,
				due: toMicro(db.ClientRetryTime())})
			m, err := toMessageProto(&dbm)
			if err != nil {
				return err
			}
			res = append(res, m)
		}
		if err := rs.Err(); err != nil {
			return err
		}
		for _, u := range updates {
			if _, err := tx.ExecContext(ctx, "UPDATE pending_messages SET retry_count=?, scheduled_time=? WHERE for_server=0 AND message_id=?", u.nc, u.due, u.id); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	return res, nil
}

func (d *Datastore) internalServerMessagesForProcessing(ctx context.Context, lim int) ([]*fspb.Message, error) {

	var dbms []*dbMessage
	// First a full transaction to read and update the pending_messages table.
	if err := d.runInTx(ctx, false, func(tx *sql.Tx) error {

		// Use two random bytes to decide where to start reading (and locking). This
		// reduces the odds of trying to grab the same records as another server
		// when there is a large backlog.
		b := [2]byte{}
		rand.Read(b[:])
		start := hex.EncodeToString(b[:])

		readPending := func(q string, l int) error {
			rs, err := tx.QueryContext(ctx, q, start, toMicro(db.Now()), l)
			if err != nil {
				return err
			}
			defer rs.Close()
			for rs.Next() {
				var dbm dbMessage
				if err := rs.Scan(
					&dbm.messageID,
					&dbm.retryCount,
					&dbm.dataTypeURL,
					&dbm.dataValue,
				); err != nil {
					return err
				}
				dbms = append(dbms, &dbm)
			}
			return rs.Err()
		}

		// First query reads from "start" to keyspace end.
		if err := readPending(
			"SELECT "+
				"message_id, "+
				"retry_count, "+
				"data_type_url, "+
				"data_value "+
				"FROM pending_messages "+
				"WHERE for_server=1 "+
				"AND message_id > ? "+
				"AND scheduled_time < ? "+
				"ORDER BY message_id "+
				"LIMIT ? "+
				"FOR UPDATE", lim); err != nil {
			return err
		}
		// If needed, second reads from keyspace begin to "start".
		if len(dbms) < lim {
			if err := readPending(
				"SELECT "+
					"message_id, "+
					"retry_count, "+
					"data_type_url, "+
					"data_value "+
					"FROM pending_messages "+
					"WHERE for_server=1 "+
					"AND message_id < ? "+
					"AND scheduled_time < ? "+
					"ORDER BY message_id "+
					"LIMIT ? "+
					"FOR UPDATE", lim-len(dbms)); err != nil {
				return err
			}
		}
		stmt, err := tx.PrepareContext(ctx, "UPDATE pending_messages SET retry_count=?, scheduled_time=? WHERE for_server=1 AND message_id=?")
		if err != nil {
			return err
		}
		for _, dbm := range dbms {
			if _, err := stmt.ExecContext(ctx, dbm.retryCount+1, toMicro(db.ServerRetryTime(dbm.retryCount+1)), dbm.messageID); err != nil {
				return err
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}
	if dbms == nil {
		return nil, nil
	}
	// Now we read the rest of the message metadata. In a perfect world, we'd
	// rollback the transaction above if anything here fails, however we do this
	// outside of the transaction in order to release any locks held by the
	// transaction asap.
	res := make([]*fspb.Message, 0, len(dbms))
	stmt, err := d.db.PrepareContext(ctx, "SELECT "+
		"source_client_id, "+
		"source_service_name, "+
		"source_message_id, "+
		"destination_client_id, "+
		"destination_service_name, "+
		"message_type, "+
		"creation_time_seconds, "+
		"creation_time_nanos, "+
		"validation_info, "+
		"annotations "+
		"FROM messages "+
		"WHERE message_id = ?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	for _, dbm := range dbms {
		r := stmt.QueryRowContext(ctx, dbm.messageID)
		if err := r.Scan(
			&dbm.sourceClientID,
			&dbm.sourceServiceName,
			&dbm.sourceMessageID,
			&dbm.destinationClientID,
			&dbm.destinationServiceName,
			&dbm.messageType,
			&dbm.creationTimeSeconds,
			&dbm.creationTimeNanos,
			&dbm.validationInfo,
			&dbm.annotations,
		); err != nil {
			return nil, err
		}
		m, err := toMessageProto(dbm)
		if err != nil {
			return nil, err
		}
		res = append(res, m)
	}
	return res, nil
}

type messageLooper struct {
	d *Datastore

	mp               db.MessageProcessor
	processingTicker *time.Ticker
	stopCalled       chan struct{}
	loopDone         chan struct{}
}

func (d *Datastore) RegisterMessageProcessor(mp db.MessageProcessor) {
	if d.looper != nil {
		log.Warning("Attempt to register a second MessageProcessor.")
		d.looper.stop()
	}
	d.looper = &messageLooper{
		d:                d,
		mp:               mp,
		processingTicker: time.NewTicker(300 * time.Millisecond),
		stopCalled:       make(chan struct{}),
		loopDone:         make(chan struct{}),
	}
	go d.looper.messageProcessingLoop()
}

func (d *Datastore) StopMessageProcessor() {
	if d.looper != nil {
		d.looper.stop()
	}
	d.looper = nil
}

// messageProcessingLoop reads messages that should be processed on the server
// from the datastore and delivers them to the registered MessageProcessor.
func (l *messageLooper) messageProcessingLoop() {
	defer close(l.loopDone)
	for {
		select {
		case <-l.stopCalled:
			return
		case <-l.processingTicker.C:
			l.processMessages()
		}
	}
}

func (l *messageLooper) stop() {
	l.processingTicker.Stop()
	close(l.stopCalled)
	<-l.loopDone
}

func (l *messageLooper) processMessages() {
	for {
		msgs, err := l.d.internalServerMessagesForProcessing(context.Background(), 10)
		if err != nil {
			log.Errorf("Failed to read server messages for processing: %v", err)
			continue
		}
		l.mp.ProcessMessages(msgs)
		if len(msgs) == 0 {
			return
		}
	}
}
