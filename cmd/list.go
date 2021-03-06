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

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists manipulation",
	Long:  `List lists of the board`,
	Run: func(cmd *cobra.Command, args []string) {
		u, _ := url.ParseRequestURI(Config.APIURL)
		u.RawQuery = fmt.Sprintf("key=%s&token=%s", Config.Key, Config.Token)
		u.Path = fmt.Sprintf("/1/boards/%s/lists", Config.BoardID)
		lists := new([]model.List)
		url2Struct(u, lists)
		for _, l := range *lists {
			if Verbose {
				fmt.Println(l.ToString())
			} else {
				fmt.Println(l.Name)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
