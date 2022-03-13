package Spotify

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func GetUserProfile() string {
	var Uri string =  "https://api.spotify.com/v1/me"

	fmt.Println("Access token to get user profile ->", AccessT.Access_token)

	req, _ := http.NewRequest("GET", Uri, nil)
	req.Header.Set("Authorization", "Bearer " +  AccessT.Access_token)
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		fmt.Println(err)
		return "Error while doing a request in API.go"
	}

	fmt.Println("Status Code", resp.StatusCode)

	response, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Response Data \n", string(response))
	return string(response)
}