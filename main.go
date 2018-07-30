package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
	"net/http/cookiejar"
	"net/url"

	"golang.org/x/net/publicsuffix"
)

type Pad struct {
	ID   string `json:"id"`
	Text string `json:"text"`
	Time int    `json:"time"`
}

type History struct {
	Pads []Pad `json:"history"`
}

func main() {
	baseURL := "https://md.darmstadt.ccc.de"
	user := "fleaz"
	password := "LOLNOPE"

	jar, err := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})
	if err != nil {
		log.Fatal(err)
	}

	client := &http.Client{
		Jar: jar,
	}

	fmt.Println("Logging in")
	authURL := fmt.Sprintf("%s/auth/ldap", baseURL)
	formData := url.Values{}
	formData.Add("username", user)
	formData.Add("password", password)
	if _, err = client.PostForm(authURL, formData); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Logged in")

	histURL := fmt.Sprintf("%s/history", baseURL)
	r, err := client.Get(histURL)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Got Hist")

	var h History
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(&h)
	if err != nil {
		log.Printf("verbose error info: %#v", err)
	}

	for _, p := range h.Pads {
		fmt.Printf("Found pad with the name %q in your history\n", p.Text)
	}

}
