// Copyright 2017 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
)

func TestMain(t *testing.T) {
	fmt.Println("hello")
	// don't launch etcd server when invoked via go test
	// Note: module name has /v3 now
	if strings.HasSuffix(os.Args[0], "v3.test") {
		t.Skip("skip launching etcd server when invoked via go test")
	}
	if len(os.Args) > 1 && strings.HasPrefix(os.Args[1], "-test.") {
		t.Skip("skip launching etcd server when invoked via go test")
	}

	notifier := make(chan os.Signal, 1)
	signal.Notify(notifier, syscall.SIGINT, syscall.SIGTERM)
	go main()
	<-notifier
}
