// Copyright 2017 Percona LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mongod

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/monotek/mongodb_exporter/testutils"
)

func TestGetReplSetConfDecodesFine(t *testing.T) {
	// setup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client := testutils.MustGetConnectedReplSetClient(ctx, t)
	defer client.Disconnect(ctx)

	// run
	status := GetReplSetConf(client)

	// test
	assert.NotNil(t, status)
}
