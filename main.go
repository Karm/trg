package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

/*
Payload

   apiUrl := "https://api.com"
    resource := "/user/"
    data := url.Values{}
    data.Set("name", "foo")
    data.Set("surname", "bar")

    u, _ := url.ParseRequestURI(apiUrl)
    u.Path = resource
    urlStr := u.String() // "https://api.com/user/"

    client := &http.Client{}
    r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
    r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
    r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
    r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

    resp, _ := client.Do(r)
    fmt.Println(resp.Status)v
*/

var myClient = &http.Client{Timeout: 10 * time.Second}

func getStruct(u *url.URL, t interface{}) error {
	r, err := myClient.Get(u.String())
	if err != nil {
		return err
	}
	defer r.Body.Close()
	//return json.NewDecoder(r.Body).Decode(t)
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		fmt.Println(string(body))
		return json.Unmarshal(body, &t)
	}
	return err
}

func main() {
	argsWithProg := os.Args
	boardID := argsWithProg[1]
	key := argsWithProg[2]
	token := argsWithProg[3]

	apiURL := "https://api.trello.com/"
	u, _ := url.ParseRequestURI(apiURL)
	u.RawQuery = fmt.Sprintf("key=%s&token=%s", key, token)

	//println(u.String())

	//client := &http.Client{{Timeout: 10 * time.Second}
	//  r, _ := http.NewRequest("POST", urlStr, strings.NewReader(data.Encode())) // URL-encoded payload
	//r.Header.Add("Authorization", "auth_token=\"XXXXXXX\"")
	//r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	//r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))

	//resp, _ := client.Do(r)
	//fmt.Println(resp.Status)

	u.Path = fmt.Sprintf("/1/boards/%s/members", boardID)
	members := new([]Member)
	getStruct(u, members)

	u.Path = fmt.Sprintf("/1/boards/%s/labels", boardID)
	labels := new([]Label)
	getStruct(u, labels)

	u.Path = fmt.Sprintf("/1/boards/%s/lists/%s", boardID, "open")
	lists := new([]List)
	getStruct(u, lists)
}
