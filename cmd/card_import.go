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
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/Karm/trg/model"
	"github.com/spf13/cobra"
)

// importCmd represents the import command
var cardImportCmd = &cobra.Command{
	Use:   "import",
	Short: "Import cards from CSV file",
	Long: `Importing cards from CSV requires a CSV file of this specific format, e.g.
    +-----------------------+--------------------------------------------+-------------+----------+
    |          Name         |                 Description                |  Assignee   |  Labels  |
    +-----------------------+--------------------------------------------+-------------+----------+
    | RPM coverage for X    | https://issues.jboss.org/browse/JBQA-13283 | janonderka1 | RPM      |
    | TLS 1.3 httpd, tomcat | https://issues.jboss.org/browse/JBQA-14119 | karm2       | JWS,JBCS |
    +-----------------------+--------------------------------------------+-------------+----------+
    Summary MUST NOT be empty.
	Issue key MUST NOT be empty.
	Assignee MIGHT be emtpy.
	Labes is a comma separated list of labels and MIGHT be empty.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		path, err := cmd.Flags().GetString("path")
		if err != nil {
			log.Fatal(err)
		}

		csvFile, err := os.Open(path)
		if err != nil {
			log.Fatal(err)
		}
		defer csvFile.Close()

		reader := csv.NewReader(bufio.NewReader(csvFile))

		delayms, err := cmd.Flags().GetInt("delayms")
		if err != nil {
			log.Fatal(err)
		}

		u, _ := url.ParseRequestURI(Config.APIURL)
		u.RawQuery = fmt.Sprintf("key=%s&token=%s", Config.Key, Config.Token)

		u.Path = fmt.Sprintf("/1/boards/%s/members", Config.BoardID)
		members := new([]model.Member)
		if err = url2Struct(u, members); err != nil {
			log.Fatalf("Failed to retrieve members. %s. Quitting.", err)
		}

		u.Path = fmt.Sprintf("/1/boards/%s/labels", Config.BoardID)
		labels := new([]model.Label)
		if err = url2Struct(u, labels); err != nil {
			log.Fatalf("Failed to retrieve labels. %s. Quitting.", err)
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

		first := true
		rowNum := 1
		fmt.Println("\"Row number\",\"Card id\",\"Name\",\"Description\",\"Assignee\",\"Labels\"")
		for {
			line, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				log.Fatal(err)
			}
			if first {
				first = false
				continue
			}

			var memberID string
			for _, m := range *members {
				if strings.Trim(line[2], " ") == m.Username {
					memberID = m.ID
					break
				}
			}
			if memberID == "" {
				log.Printf("Provided assignee %s does not exist. Skipping this row.\n", strings.Trim(line[2], " "))
				rowNum++
				continue
			}

			labelsU := strings.Split(line[3], ",")
			labelIDs := make([]string, 0, len(*labels))
			for _, lu := range labelsU {
				for _, l := range *labels {
					if strings.Trim(lu, " ") == l.Name {
						labelIDs = append(labelIDs, l.ID)
					}
				}
			}
			if len(labelIDs) != len(labelsU) {
				log.Println("Some provided label names not found. Skipping this row.")
				rowNum++
				continue
			}

			if strings.Trim(line[0], " ") == "" {
				log.Println("Name, the first col of this row is empty. Skipping this row.")
				rowNum++
				continue
			}

			u.Path = "/1/cards"
			payload := url.Values{}
			payload.Set("name", line[0])
			payload.Set("desc", line[1])
			payload.Set("pos", "top")
			payload.Set("idList", listID)
			payload.Set("boardId", Config.BoardID)
			payload.Set("idMembers", memberID)
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
			fmt.Printf("\"%d\",\"%s\",\"%s\",\"%s\",\"%s\",\"%s\"\n", rowNum, card.ID, line[0], line[1], line[2], line[3])
			if Verbose {
				fmt.Println(card.ToString())
			}
			time.Sleep(time.Duration(delayms) * time.Second)
			rowNum++
		}
	},
}

func init() {
	cardCmd.AddCommand(cardImportCmd)
	cardImportCmd.Flags().StringP("path", "p", "", "Path to CSV file")
	cardImportCmd.Flags().StringP("list", "i", "", "Name of the list cards will appear in")
	cardImportCmd.Flags().IntP("delayms", "d", 0, "Optional delay between rows in ms to workaround rate limiting")
	cardImportCmd.MarkFlagRequired("path")
	cardImportCmd.MarkFlagRequired("list")
}
