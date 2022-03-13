package Spotify

import (
	"fmt"
	"net/http"
	"bytes"
)

func GetProfile(acces_token string) {

	var ReqUri = "https://api.spotify.com/v1/tracks/2TpxZ7JUBn3uw46aR7qd6V"
	
	req, _ := http.NewRequest("GET", ReqUri, bytes.NewBuffer([]byte( `{ "username":"nothing" }` )))
	strToken := fmt.Sprintf("Bearer %s", acces_token)
	req.Header.Set("Authorization", strToken)
	req.Header.Set("Content-Type", "application/json")

	var client = &http.Client{}
	resp, _ := client.Do(req)

	fmt.Println("Response", resp)

}