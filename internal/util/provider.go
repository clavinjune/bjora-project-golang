// Copyright 2022 ClavinJune/bjora
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

package util

import (
	"sync"

	"github.com/bwmarrin/snowflake"
	"github.com/google/wire"
)

var (
	snowInst *snowflake.Node
	snowOnce sync.Once

	// ProviderSet contains all provider from util package
	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideSnowflake,
	)
)

// ProvideSnowflake produces *snowflake.Node
func ProvideSnowflake() *snowflake.Node {
	snowOnce.Do(func() {
		snowInst, _ = snowflake.NewNode(1)
	})

	return snowInst
}
