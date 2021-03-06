// Copyright © 2019 Michal Karm Babacek <karm@redhat.com>
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

package model

import "fmt"

// List struct
type List struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Closed  bool   `json:"closed"`
	IDBoard string `json:"idBoard"`
}

// ToString spits out the format we need...
func (l *List) ToString() string {
	return fmt.Sprintf(
		`ID:      %s
Name:    %s
Closed:  %t
IDBoard: %s`, l.ID, l.Name, l.Closed, l.IDBoard)
}
