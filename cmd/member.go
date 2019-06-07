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

package cmd

import (
	"fmt"
	"net/url"

	"github.com/Karm/trg/model"
	"github.com/spf13/cobra"
)

// memberCmd represents the member command
var memberCmd = &cobra.Command{
	Use:   "member",
	Short: "Members (users) manipulation",
	Long:  `List members of the board`,
	Run: func(cmd *cobra.Command, args []string) {
		u, _ := url.ParseRequestURI(Config.APIURL)
		u.RawQuery = fmt.Sprintf("key=%s&token=%s", Config.Key, Config.Token)
		u.Path = fmt.Sprintf("/1/boards/%s/members", Config.BoardID)
		members := new([]model.Member)
		url2Struct(u, members)
		for _, m := range *members {
			if Verbose {
				fmt.Println(m.ToString())
			} else {
				fmt.Println(m.Username)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(memberCmd)
}
