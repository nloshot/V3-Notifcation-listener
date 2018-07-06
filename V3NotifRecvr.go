package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"os"
	"strings"
	"time"
)

func main() {

	http.HandleFunc("/es2", es2handler)
	http.HandleFunc("/es4", es4handler)
	http.ListenAndServe(":8080", nil)

}

func es2handler(w http.ResponseWriter, r *http.Request) {
	//	Display something in the browser
	io.WriteString(w, "ES2 Acknowledge Response here...")

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	write2file(string(requestDump))

}

func es4handler(w http.ResponseWriter, r *http.Request) {
	//	Display something in the browser
	io.WriteString(w, "Put the ES4 Acknowledge Response here...")

	// Save a copy of this request for debugging.
	requestDump, err := httputil.DumpRequest(r, true)
	if err != nil {
		fmt.Println(err)
	}

	write2file(string(requestDump))
}

func write2file(dmp string) {

	t := time.Now()
	//formatt := t.Format(time.RFC1123)
	formatt := t.Format(time.RFC822)

	//Make it readable by inserting newlines
	dmp = strings.Replace(dmp, "><", "> \n <", -1)

	f, err := os.OpenFile("V3notiflog.txt", os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString("<<------------- " + formatt + " ------------->>\n" + dmp + "\n\n\n"); err != nil {
		panic(err)
	}

}
