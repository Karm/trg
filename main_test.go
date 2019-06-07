package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
	"time"
)

var token string
var key string
var boardID string

func init() {
	token = os.Getenv("TRG_TOKEN")
	if len(token) < 24 {
		log.Fatalln("TRG_TOKEN env var expected to run the testsuite.")
	}
	key = os.Getenv("TRG_KEY")
	if len(key) < 24 {
		log.Fatalln("TRG_KEY env var expected to run the testsuite.")
	}
	boardID = os.Getenv("TRG_BOARDID")
	if len(boardID) < 4 {
		log.Fatalln("TRG_BOARDID env var expected to run the testsuite.")
	}
}

// This is just a stub
func TestMain(m *testing.M) {
	stdOutBak := os.Stdout
	stdErrBak := os.Stderr
	stdInBak := os.Stdin

	rOut, wOut, _ := os.Pipe()
	rIn, wIn, _ := os.Pipe()
	rErr, wErr, _ := os.Pipe()

	os.Stdout = wOut
	os.Stderr = wErr
	os.Stdin = rIn

	go func(wOut *os.File, wErr *os.File, rIn *os.File) {
		os.Args[1] = "install"
		os.Args[2] = "--config=./test.trg.json"
		os.Stdout = wOut
		os.Stderr = wErr
		os.Stdin = rIn
		main()
	}(wOut, wErr, rIn)
	time.Sleep(100 * time.Millisecond)

	wOut.Close()
	wErr.Close()
	out, _ := ioutil.ReadAll(rOut)
	err, _ := ioutil.ReadAll(rErr)
	os.Stdout = stdOutBak
	os.Stderr = stdErrBak
	fmt.Printf("%s", err)
	fmt.Printf("%s", out)

	rOut, wOut, _ = os.Pipe()
	rErr, wErr, _ = os.Pipe()
	os.Stdout = wOut
	os.Stderr = wErr

	if strings.Contains(string(out), "Begin installation and write to ./test.trg.json ? y/n [n]:") {
		wIn.WriteString("y\n")
	} else {
		fmt.Fprintf(stdErrBak, "Unexpected sentence in instalaltion wizzard: %s", string(out))
		os.Exit(2)
	}

	time.Sleep(100 * time.Millisecond)

	wOut.Close()
	wErr.Close()
	out, _ = ioutil.ReadAll(rOut)
	err, _ = ioutil.ReadAll(rErr)
	os.Stdout = stdOutBak
	os.Stderr = stdErrBak
	fmt.Printf("%s", err)
	fmt.Printf("%s", out)
	os.Stdin = stdInBak

	/*
		// Begin installation and write to ./test.trg.json ? y/n [n]:
		win.WriteString("y\n")
		//win.Sync()
		//time.Sleep(100 * time.Millisecond)

			// Do you have Trello API key and token? y/n [n]:
			win.WriteString("y\n")
		//	win.Sync()
		//	time.Sleep(100 * time.Millisecond)

			win.WriteString(token + "\n")
		//	win.Sync()
		//	time.Sleep(100 * time.Millisecond)

			win.WriteString(key + "\n")
		//	win.Sync()
		//	time.Sleep(100 * time.Millisecond)

			win.WriteString("\n")
		//	win.Sync()
		//	time.Sleep(100 * time.Millisecond)

			win.WriteString(boardID + "\n")
		//	win.Sync()
		//	time.Sleep(100 * time.Millisecond)

			win.WriteString("\n")
		//	win.Sync()
		//	time.Sleep(100 * time.Millisecond)

			win.WriteString("\n")
			win.Sync()
			time.Sleep(100 * time.Millisecond)

			w.Close()
			win.Close()

			out, _ = ioutil.ReadAll(r)

			os.Stdout = stdOutBak
			os.Stderr = stdErrBak
			os.Stdin = stdInBak

			fmt.Printf("%s", out)
			os.Exit(m.Run())
	*/
}

func TestCard(t *testing.T) {
	stdOutBak := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	os.Args[1] = "list"
	os.Args[2] = "--config=./test.trg.json"
	main()
	time.Sleep(100 * time.Millisecond)
	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = stdOutBak
	fmt.Printf("Captured: %s", out)
}
