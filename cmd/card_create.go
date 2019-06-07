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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Karm/trg/model"
	"github.com/spf13/cobra"
)

// createCmd represents the create command
var cardCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create card",
	Long: `Creating cards
	
	e.g. -n "API CARD" -m paullodge2,matusmadzin1,karm2 -i "JBQA Backlog"
	     -l OpenShift,OneOff,JWS -d 'This is the body of the card created via API'`,
	Run: func(cmd *cobra.Command, args []string) {

		u, _ := url.ParseRequestURI(Config.APIURL)
		u.RawQuery = fmt.Sprintf("key=%s&token=%s", Config.Key, Config.Token)

		membersArg, err := cmd.Flags().GetString("members")
		if err != nil {
			log.Fatal(err)
		}
		membersU := strings.Split(membersArg, ",")
		u.Path = fmt.Sprintf("/1/boards/%s/members", Config.BoardID)
		members := new([]model.Member)
		if err = url2Struct(u, members); err != nil {
			log.Fatalf("Failed to retrieve members. %s. Quitting.", err)
		}
		memberIDs := make([]string, 0, len(*members))
		if Verbose {
			fmt.Println("Using members: ")
		}
		for _, mu := range membersU {
			for _, m := range *members {
				if strings.Trim(mu, " ") == m.Username {
					memberIDs = append(memberIDs, m.ID)
					if Verbose {
						fmt.Printf("%s\n", m.Username)
					}
				}
			}
		}
		if len(memberIDs) == 0 {
			log.Fatalln("No valid members found. Check your input. Quitting.")
		}
		if len(memberIDs) != len(membersU) {
			log.Println("Some members were not found.")
		}

		labelsArg, err := cmd.Flags().GetString("labels")
		if err != nil {
			log.Fatal(err)
		}
		labelsU := strings.Split(labelsArg, ",")
		u.Path = fmt.Sprintf("/1/boards/%s/labels", Config.BoardID)
		labels := new([]model.Label)
		if err = url2Struct(u, labels); err != nil {
			log.Fatalf("Failed to retrieve labels. %s. Quitting.", err)
		}
		labelIDs := make([]string, 0, len(*labels))
		if Verbose {
			fmt.Println("Using labels: ")
		}
		for _, lu := range labelsU {
			for _, l := range *labels {
				if strings.Trim(lu, " ") == l.Name {
					labelIDs = append(labelIDs, l.ID)
					if Verbose {
						fmt.Printf("%s\n", l.Name)
					}
				}
			}
		}
		if len(labelIDs) != len(labelsU) {
			log.Println("Some label names were not found.")
		}

		listName, err := cmd.Flags().GetString("list")
		if err != nil {
			log.Fatal(err)
		}
		u.Path = fmt.Sprintf("/1/boards/%s/lists/%s", Config.BoardID, "open")
		lists := new([]model.List)
		if err = url2Struct(u, lists); err != nil {
			log.Fatalf("Failed to retrieve lists. %s. Quitting.", err)
		}
		var listID string
		for _, l := range *lists {
			if l.Name == strings.Trim(listName, " ") {
				listID = l.ID
				break
			}
		}
		if listID == "" {
			log.Fatalln("Provided list does not exist. Quitting.")
		}

		name, err := cmd.Flags().GetString("name")
		if err != nil {
			log.Fatal(err)
		}
		if name == "" {
			log.Fatalln("Provided name is empty. Quitting.")
		}
		desc, err := cmd.Flags().GetString("desc")
		if err != nil {
			log.Fatal(err)
		}

		u.Path = "/1/cards"
		payload := url.Values{}
		payload.Set("name", name)
		payload.Set("desc", desc)
		payload.Set("pos", "top")
		payload.Set("idList", listID)
		payload.Set("boardId", Config.BoardID)
		payload.Set("idMembers", strings.Join(memberIDs, ","))
		payload.Set("idLabels", strings.Join(labelIDs, ","))

		req, err := http.NewRequest("POST", u.String(), strings.NewReader(payload.Encode()))
		if err != nil {
			log.Fatal(err)
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		r, err := myClient.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer r.Body.Close()
		card := new(model.Card)
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			if Verbose {
				fmt.Println(string(body))
			}
			err = json.Unmarshal(body, &card)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(err)
		}
		fmt.Printf("Created card. ID: %s\n", card.ID)
		if Verbose {
			fmt.Println(card.ToString())
		}
	},
}

func init() {
	cardCmd.AddCommand(cardCreateCmd)
	cardCreateCmd.Flags().StringP("name", "n", "", "Name for the card")
	cardCreateCmd.Flags().StringP("labels", "l", "", "Comma separated list of label names")
	cardCreateCmd.Flags().StringP("members", "m", "", "Comma separated list of usernames")
	cardCreateCmd.Flags().StringP("list", "i", "", "Name of the list the card will appear in")
	cardCreateCmd.Flags().StringP("desc", "d", "", "Description, the full text body of the card")
	cardCreateCmd.MarkFlagRequired("name")
	cardCreateCmd.MarkFlagRequired("list")
}
