package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"io/ioutil"
	"net/http"
	"os"
)

func QueryDeHashed(email string, apiKey string, domain string) {

	// logic for handling the request
	url := "https://api.dehashed.com/search?query=" + domain + "&size=10000"
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	req.Header.Add("Accept", "application/json")
	req.SetBasicAuth(email, apiKey)
	resp, err := client.Do(req)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	respData, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	// pretty-ing the json blob
	type Entry struct {
		Email        string
		Username     string
		Password     string
		PasswordHash string
		Phone        string
		Database     string
	}

	type Response struct {
		Balance int
		Total   int
		Success bool
		Entries []Entry `json:"entries"`
	}

	var response Response
	json.Unmarshal(respData, &response)
	for i := range response.Entries {
		fmt.Println(color.GreenString("Email Address: ") + response.Entries[i].Email)
		fmt.Println(color.GreenString("Username: ") + response.Entries[i].Username)
		fmt.Println(color.GreenString("Password: ") + response.Entries[i].Password)
		fmt.Println(color.GreenString("Password Hash: ") + response.Entries[i].PasswordHash)
		fmt.Println(color.GreenString("Phone Number: ") + response.Entries[i].Phone)
		fmt.Println(color.GreenString("Database: ") + response.Entries[i].Database)
		fmt.Println("")
	}
}

func main() {

	email := flag.String("email", "", "Email address for authentication")
	apiKey := flag.String("apiKey", "", "API key for authentication")
	domain := flag.String("domain", "", "Domain to query")

	flag.Parse()

	if *email == "" || *apiKey == "" || *domain == "" {
		fmt.Println("All arguments (email, apiKey, domain) are required.")
		flag.Usage()
		os.Exit(1)
	}

	QueryDeHashed(*email, *apiKey, *domain)
}
