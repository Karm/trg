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
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"time"

	"github.com/Karm/trg/app"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

// Verbose is a global logging option
var Verbose bool

// Config for the application
var Config *app.Config

var version string
var cfgFileP string
var defaultCfgFileP string

var rootCmd = &cobra.Command{
	Use:   "trg",
	Short: "CLI for a limited scope of specific Trello actions",
	Long: ` _______ _____   _____ 
|__   __|  __ \ / ____|  is a small CLI that manipulates a selected
   | |  | |__) | |  __   Trello board. It focuses just on a very
   | |  |  _  /| | |_ |  limited scope of specific tasks.
   | |  | | \ \| |__| |  Copyright © 2019 Michal Karm Babacek
   |_|  |_|  \_\\_____|  Apache 2.0 License
` + version,
}

func url2Struct(u *url.URL, t interface{}) error {
	r, err := myClient.Get(u.String())
	if err != nil {
		return err
	}
	defer r.Body.Close()
	//return json.NewDecoder(r.Body).Decode(t)
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if Verbose {
			fmt.Println(string(body))
		}
		err = json.Unmarshal(body, &t)
	}
	return err
}

func file2Struct(path string, t interface{}) error {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return json.Unmarshal(content, t)
}

func struct2file(path string, t interface{}) error {
	json, err := json.Marshal(t)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(path, json, 0600)
}

// Execute is Root cmd entrypoint
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	home, err := homedir.Dir()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defaultCfgFileP = filepath.Join(home, ".trg.json")
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfgFileP, "config", "c", defaultCfgFileP, "config file")
	rootCmd.PersistentFlags().BoolVarP(&Verbose, "verbose", "v", false, "verbose output")
}

func initConfig() {
	if cfgFileP == "" {
		if Verbose {
			fmt.Println("Using default " + defaultCfgFileP + " config file.")
		}
		cfgFileP = defaultCfgFileP
	} else {
		if Verbose {
			fmt.Println("Using " + cfgFileP + " config file.")
		}
	}
	Config = &app.Config{}
	err := file2Struct(cfgFileP, Config)
	if err != nil {
		fmt.Println(err)
		fmt.Println("It seems no valid config was provided.")
		installApp()
	}
}

func rsp() string {
	var response string
	_, err := fmt.Scanln(&response)
	if err != nil && err.Error() != "unexpected newline" {
		log.Fatal(err)
	}
	return response
}

func installApp() {
	fmt.Printf("Begin installation and write to %s ? y/n [n]: ", cfgFileP)
	if rsp() != "y" {
		fmt.Println("Bye.")
		os.Exit(0)
	}

	Config = &app.Config{}
	fmt.Print("\nDo you have Trello API key and token? y/n [n]: ")
	if rsp() != "y" {
		fmt.Println("Visit https://trello.com/app-key and click on `generate a Token'")
		os.Exit(0)
	}

	fmt.Print("\nToken: ")
	Config.Token = rsp()
	re := regexp.MustCompile(`[a-z0-9]`)
	l := len(Config.Token)
	if l < 32 || l > 256 || !re.MatchString(Config.Token) {
		fmt.Println("Token seems invalid.")
		os.Exit(1)
	}

	fmt.Print("\nKey: ")
	Config.Key = rsp()
	l = len(Config.Key)
	if l < 31 || l > 256 || !re.MatchString(Config.Key) {
		fmt.Println("Key seems invalid.")
		os.Exit(1)
	}

	fmt.Print("\nAPI URL [https://api.trello.com/]: ")
	Config.APIURL = rsp()
	if len(Config.APIURL) < 4 {
		Config.APIURL = "https://api.trello.com/"
		fmt.Printf("Using default %s\n", Config.APIURL)
	}

	fmt.Print("\nBoard ID: ")
	Config.BoardID = rsp()
	l = len(Config.BoardID)
	if l < 7 || l > 256 || !re.MatchString(Config.BoardID) {
		fmt.Println("Board ID seems invalid.")
		os.Exit(1)
	}
	struct2file(cfgFileP, Config)

	if runtime.GOOS != "windows" {
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		ex, err := os.Executable()
		if err != nil {
			log.Fatal(err)
		}
		dir := path.Dir(ex)
		bashrc := filepath.Join(home, ".bashrc")
		addedPath := fmt.Sprintf("export PATH=%s:$PATH\n", dir)
		fmt.Printf(`
Do you want to have this line appended at the end of %s ?
%sy/n [n]: `, bashrc, addedPath)
		// Add to PATH, write to ~/.bashrc ?
		if rsp() == "y" {
			file, err := os.OpenFile(bashrc, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			_, err = file.WriteString(addedPath)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(".bashrc edited. Call: `source ~/.bashrc'")
		}

		// Add bash completion? write  to ~/.bashrc
		compl := ". <(trg completion)"
		fmt.Printf(`
To enable bash completion, you need to have
%s in your .bashrc. Do you want it appended there? y/n [n]: `, compl)
		if rsp() == "y" {
			file, err := os.OpenFile(bashrc, os.O_WRONLY|os.O_APPEND, os.ModeAppend)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()
			_, err = file.WriteString(compl)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(".bashrc edited. Call: `source ~/.bashrc'")
		}
	}
	fmt.Println("Done.")
	os.Exit(0)
}
