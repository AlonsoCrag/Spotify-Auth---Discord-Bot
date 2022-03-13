package Sv

import (
	"fmt"
	"net/http"
	"Logia-Bot/Spotify"
)

var AuthCode string

func Start() {

	http.HandleFunc("/accesstoken", func (resp http.ResponseWriter, req *http.Request) {
		fmt.Println("Incoming access code from a spotify-user")	
		fmt.Println(req.URL.Query()["code"][0])
		AuthCode = req.URL.Query()["code"][0]
		Spotify.SpotifyAuthReady(AuthCode)
		Spotify.GetUserProfile()
	})

	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		fmt.Println("Server error", err)
	}
}