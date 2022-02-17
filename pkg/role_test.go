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

package pkg_test

import (
	"testing"

	"github.com/clavinjune/bjora-project-golang/pkg"
	"github.com/stretchr/testify/assert"
)

func TestRoleFrom(t *testing.T) {
	tt := []struct {
		name string
		in   string
		want pkg.Role
	}{
		{
			name: "admin",
			in:   " admin ",
			want: pkg.RoleAdmin,
		},
		{
			name: "member",
			in:   " member ",
			want: pkg.RoleMember,
		},
		{
			name: "undefined",
			in:   "qwds",
			want: pkg.RoleUndefined,
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := pkg.RoleFrom(tc.in)

			assert.Equal(t, tc.want, got)
		})
	}
}

func TestRole_String(t *testing.T) {
	tt := []struct {
		name string
		in   pkg.Role
		want string
	}{
		{
			name: "admin",
			in:   pkg.RoleAdmin,
			want: "Admin",
		},
		{
			name: "member",
			in:   pkg.RoleMember,
			want: "Member",
		},
		{
			name: "undefined",
			in:   pkg.RoleUndefined,
			want: "Undefined",
		},
		{
			name: "random",
			in:   pkg.Role(16),
			want: "Role(16)",
		},
	}

	for i := range tt {
		tc := tt[i]

		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			got := tc.in.String()

			assert.Equal(t, tc.want, got)
		})
	}
}