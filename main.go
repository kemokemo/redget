package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"path"
)

type issue struct {
	Body body `json:"issue"`
}

type body struct {
	Description string `json:"description"`
}

var issueNum int

func init() {
	flag.IntVar(&issueNum, "n", 1, "please specify issue number.")
	flag.Parse()
}

func main() {
	os.Exit(run())
}

func run() int {
	redmine := os.Getenv("REDMINE_URL")
	key := os.Getenv("REDMINE_KEY")
	if redmine == "" || key == "" {
		fmt.Println("please set env")
		return 1
	}

	u, err := url.Parse(redmine)
	if err != nil {
		fmt.Printf("failed to parse the url %s:%v\n", redmine, err)
		return 1
	}
	u.Path = path.Join(u.Path, fmt.Sprintf("%d.json", issueNum))
	q := u.Query()
	q.Set("key", key)
	u.RawQuery = q.Encode()

	r, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("failed to get from %s: %v\n", redmine, err)
		return 1
	}
	defer r.Body.Close()

	var issue issue
	err = json.NewDecoder(r.Body).Decode(&issue)
	if err != nil {
		fmt.Println("failed to decode:", err)
		return 1
	}
	fmt.Println(issue.Body.Description)

	return 0
}
