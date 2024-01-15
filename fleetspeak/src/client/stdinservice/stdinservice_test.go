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

package stdinservice

import (
	"bytes"
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/google/fleetspeak/fleetspeak/src/client/clitesting"
	"github.com/google/fleetspeak/fleetspeak/src/common/anypbtest"

	sspb "github.com/google/fleetspeak/fleetspeak/src/client/stdinservice/proto/fleetspeak_stdinservice"
	fspb "github.com/google/fleetspeak/fleetspeak/src/common/proto/fleetspeak"
)

func TestStdinServiceWithEcho(t *testing.T) {
	s, err := Factory(&fspb.ClientServiceConfig{
		Name: "EchoService",
		Config: anypbtest.New(t, &sspb.Config{
			Cmd: "python",
		}),
	})
	if err != nil {
		t.Fatal(err)
	}

	outChan := make(chan *fspb.Message, 1)
	err = s.Start(&clitesting.MockServiceContext{
		OutChan: outChan,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = s.ProcessMessage(context.Background(),
		&fspb.Message{
			MessageType: "StdinServiceInputMessage",
			Data: anypbtest.New(t, &sspb.InputMessage{
				Args: []string{"-c", `print("foo bar")`},
			}),
		})
	if err != nil {
		t.Fatal(err)
	}

	var output *fspb.Message
	select {
	case output = <-outChan:
	default:
		t.Fatal(".ProcessMessage (/bin/echo foo bar) expected to produce message, but none found")
	}

	om := &sspb.OutputMessage{}
	if err := output.Data.UnmarshalTo(om); err != nil {
		t.Fatal(err)
	}

	wantStdout := []byte("foo bar\n")
	wantStdoutWin := []byte("foo bar\r\n")
	if !bytes.Equal(om.Stdout, wantStdout) &&
		!bytes.Equal(om.Stdout, wantStdoutWin) {
		t.Fatalf("unexpected output; got %q, want %q", om.Stdout, wantStdout)
	}
}

func TestStdinServiceWithCat(t *testing.T) {
	s, err := Factory(&fspb.ClientServiceConfig{
		Name: "CatService",
		Config: anypbtest.New(t, &sspb.Config{
			Cmd: "python",
		}),
	})
	if err != nil {
		t.Fatal(err)
	}

	outChan := make(chan *fspb.Message, 1)
	err = s.Start(&clitesting.MockServiceContext{
		OutChan: outChan,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = s.ProcessMessage(context.Background(),
		&fspb.Message{
			MessageType: "StdinServiceInputMessage",
			Data: anypbtest.New(t, &sspb.InputMessage{
				Args: []string{"-c", `
try:
  my_input = raw_input  # Python2 compat
except NameError:
  my_input = input

try:
  while True:
    print(my_input())
except EOFError:
  pass
		`},
				Input: []byte("foo bar"),
			}),
		})
	if err != nil {
		t.Fatalf("s.ProcessMessage(...) = %q, want success", err)
	}

	var output *fspb.Message
	select {
	case output = <-outChan:
	default:
		t.Fatal(".ProcessMessage (/bin/cat <<< 'foo bar') expected to produce message, but none found")
	}

	om := &sspb.OutputMessage{}
	if err := output.Data.UnmarshalTo(om); err != nil {
		t.Fatal(err)
	}

	wantStdout := []byte("foo bar\n")
	wantStdoutWin := []byte("foo bar\r\n")
	if !bytes.Equal(om.Stdout, wantStdout) &&
		!bytes.Equal(om.Stdout, wantStdoutWin) {
		t.Fatalf("unexpected output; got %q, want %q", om.Stdout, wantStdout)
	}
}

func TestStdinServiceReportsResourceUsage(t *testing.T) {
	s, err := Factory(&fspb.ClientServiceConfig{
		Name: "BashService",
		Config: anypbtest.New(t, &sspb.Config{
			Cmd: "python",
		}),
	})
	if err != nil {
		t.Fatal(err)
	}

	outChan := make(chan *fspb.Message, 1)
	err = s.Start(&clitesting.MockServiceContext{
		OutChan: outChan,
	})
	if err != nil {
		t.Fatal(err)
	}

	err = s.ProcessMessage(context.Background(),
		&fspb.Message{
			MessageType: "StdinServiceInputMessage",
			Data: anypbtest.New(t, &sspb.InputMessage{
				// Generate some system (os.listdir) and user (everything else) execution time...
				Args: []string{"-c", `
import os
import time

t0 = time.time()
while time.time() - t0 < 1.:
  os.listdir(".")
		`},
			}),
		})
	if err != nil {
		t.Fatal(err)
	}

	var output *fspb.Message
	select {
	case output = <-outChan:
	default:
		t.Fatal(".ProcessMessage (/bin/bash ...) expected to produce message, but none found")
	}

	om := &sspb.OutputMessage{}
	if err := output.Data.UnmarshalTo(om); err != nil {
		t.Fatal(err)
	}

	if om.ResourceUsage == nil {
		t.Fatalf("unexpected output; StdinServiceOutputMessage.resource_usage not set: %q", om)
	}

	if om.ResourceUsage.MeanUserCpuRate <= 0 {
		t.Fatalf("unexpected output; StdinServiceOutputMessage.resource_usage.mean_user_cpu_rate not set: %q", om)
	}

	if om.ResourceUsage.MeanSystemCpuRate <= 0 {
		t.Fatalf("unexpected output; StdinServiceOutputMessage.resource_usage.mean_system_cpu_rate not set: %q", om)
	}

	// We don't test for ResourceUsage.MeanResidentMemory because memory is currently not being
	// queried after the process has terminated. It's only queried right after launching the command
	// in which case it can be recorded as "0" which would be indistinguishable from it not being set
	// at all, resulting in a flaky test case. The fact that the other resource usage metrics have
	// been set here is good enough for now.

	if om.Timestamp.Seconds <= 0 {
		t.Fatalf("unexpected output; StdinServiceOutputMessage.timestamp.seconds not set: %q", om)
	}
}

func TestStdinServiceCancellation(t *testing.T) {
	s, err := Factory(&fspb.ClientServiceConfig{
		Name: "SleepService",
		Config: anypbtest.New(t, &sspb.Config{
			Cmd: "python",
		}),
	})
	if err != nil {
		t.Fatal(err)
	}

	outChan := make(chan *fspb.Message, 1)
	err = s.Start(&clitesting.MockServiceContext{
		OutChan: outChan,
	})
	if err != nil {
		t.Fatal(err)
	}

	ctx, c := context.WithCancel(context.Background())
	c()

	if err := s.ProcessMessage(ctx,
		&fspb.Message{
			MessageType: "StdinServiceInputMessage",
			Data: anypbtest.New(t, &sspb.InputMessage{
				Args: []string{"-c", fmt.Sprintf(`
import time

time.sleep(%f)
		`, clitesting.MockCommTimeout.Seconds())},
			}),
		}); err != nil && !strings.HasSuffix(err.Error(), "context canceled") {
		t.Fatal(err)
	} else if err == nil {
		t.Fatal(".ProcessMessage was expected to be cancelled, but returned with no error")
	}
}
