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
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var cardDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete card",
	Long:  `Deleting cards by their ID. Operation is not reversible`,
	Run: func(cmd *cobra.Command, args []string) {

		idsF, err := cmd.Flags().GetString("ids")
		if err != nil {
			log.Fatal(err)
		}
		ids := strings.Split(idsF, ",")
		u, _ := url.ParseRequestURI(Config.APIURL)
		u.RawQuery = fmt.Sprintf("key=%s&token=%s", Config.Key, Config.Token)

		for _, id := range ids {
			id = strings.Trim(id, " ")
			if len(id) < 4 {
				fmt.Printf("ID %s looks too short. Skipping it.\n", id)
				continue
			}

			u.Path = fmt.Sprintf("/1/cards/%s", id)
			if Verbose {
				fmt.Printf("Path: %s\n", u.String())
			}

			req, err := http.NewRequest("DELETE", u.String(), nil)
			if err != nil {
				log.Fatal(err)
			}

			r, err := myClient.Do(req)
			if err != nil {
				log.Fatal(err)
			}
			defer r.Body.Close()

			if body, err := ioutil.ReadAll(r.Body); err == nil {
				if Verbose {
					fmt.Println(string(body))
				}
				if r.StatusCode == 200 {
					fmt.Printf("Deleted %s.\n", id)
				} else {
					fmt.Printf("Failed to delete %s. Perhaps wrong id?\n", id)
					continue
				}
			} else {
				fmt.Printf("Failed to delete %s.\n", id)
				continue
			}
		}
	},
}

func init() {
	cardCmd.AddCommand(cardDeleteCmd)
	cardDeleteCmd.Flags().String("ids", "", "Card IDs to be deleted, comma separated")
	cardDeleteCmd.MarkFlagRequired("ids")
}
