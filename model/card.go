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

// Card struct
type Card struct {
	ID        string         `json:"id"`
	Closed    bool           `json:"closed"`
	Desc      string         `json:"desc"`
	IDBoard   string         `json:"idBoard"`
	IDLabels  []string `json:"idLabels"`
	IDList    string         `json:"idList"`
	IDMembers []string `json:"idMembers"`
	Name      string         `json:"name"`
	ShortURL  string         `json:"shortUrl"`
}

// ToString spits out the format we need...
func (c *Card) ToString() string {
	return fmt.Sprintf(
		`ID:   %s
Name: %s
URL:  %s
Desc: %s`, c.ID, c.Name, c.ShortURL, c.Desc)
}
