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
	"log"
	"net/url"

	"github.com/Karm/trg/model"
	"github.com/spf13/cobra"
)

// cardCmd represents the card command
var cardCmd = &cobra.Command{
	Use:   "card",
	Short: "Cards manipulation",
	Long:  `Get card by its ID`,
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetString("id")
		if err != nil {
			log.Fatal(err)
		}
		if Verbose {
			fmt.Printf("Loading card ID %s\n", id)
		}
		u, _ := url.ParseRequestURI(Config.APIURL)
		u.RawQuery = fmt.Sprintf("key=%s&token=%s", Config.Key, Config.Token)
		u.Path = fmt.Sprintf("/1/cards/%s", id)
		if Verbose {
			fmt.Printf("Path: %s\n", u.String())
		}
		card := new(model.Card)
		url2Struct(u, card)
		fmt.Println(card.ToString())
	},
}

func init() {
	rootCmd.AddCommand(cardCmd)
	cardCmd.Flags().String("id", "", "Card ID")
	cardCmd.MarkFlagRequired("id")
}
