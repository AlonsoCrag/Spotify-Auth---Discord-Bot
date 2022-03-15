package Spotify

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
	"github.com/toqueteos/webbrowser"
	"encoding/json"
)

type AccesToken struct {
	Access_token string
	Token_type string
	Expires_in int
	Refresh_token string
}

var AccessT AccesToken

func ReqAuthCode(code string) []byte {
	var reqAuthUri = "https://accounts.spotify.com/api/token"

	var body = []byte("grant_type=authorization_code" + "&" + "code=" + code + "&" + "redirect_uri=" + "http://localhost:3000/access/")
	fmt.Println("Final body of the requests",  bytes.NewBuffer(body))

	encodedString := base64.StdEncoding.EncodeToString([]byte("application_id_string" + ":" + "token_id_string"))
	encoded := fmt.Sprintf("Basic %s", encodedString)

	req, _ := http.NewRequest("POST", reqAuthUri, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded") 
	// The way you set ContentType will differe in how you send data to the request
	req.Header.Set("Authorization", encoded)


	client := &http.Client{}
	response, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return []byte("Error")
	}
	
	fmt.Println(response.Status)
	respBody, _ := ioutil.ReadAll(response.Body)
	fmt.Println("Body Response", string(respBody))

	return respBody
}

func Auhtorize() {
	params := url.Values{}
	params.Add("client_id", "client_id")
	params.Add("response_type", "code")
	params.Add("redirect_uri", "http://localhost:3000/access/")

	link := "https://accounts.spotify.com/authorize?" + params.Encode()
	webbrowser.Open(link)
	fmt.Println("Spotify UI has been closed")
}


func SpotifyAuthReady(AuthCode string) {
	fmt.Println("A new spotify access token was found....", AuthCode)
	AccessTokenBytes := ReqAuthCode(AuthCode)

	json.Unmarshal(AccessTokenBytes, &AccessT)
	fmt.Println("Response Token", AccessT)
}
