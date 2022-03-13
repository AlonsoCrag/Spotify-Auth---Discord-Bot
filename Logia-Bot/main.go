package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"Logia-Bot/Spotify"
	"Logia-Bot/Sv"
	// "encoding/json"
)

type Response struct {
	Access_token string
	Token_type string
	Expires_in int
}

var BotId string
var Resp Response

func main() {

	fmt.Println("Application to manage a Discord Bot")
	bot, err := discordgo.New("Bot " + "OTUxMzY5NTY0ODk2MjQzNzMy.Yimd8w.0cd-yj_G-wNZmBfcuUJSEV6TjWI")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	user, err := bot.User("@me")
	if err != nil {
		fmt.Println(err)
	}

	BotId = user.ID

	bot.AddHandler(messageHandler)
	err = bot.Open()

	if err != nil {
		fmt.Println("Error, While trying to open")
		fmt.Println(err.Error())
		return
	}

	// acces_token := Spotify.ReqAuth()

	// bytes := []byte(acces_token)
	// json.Unmarshal(bytes, &Resp)

	// fmt.Println("Response from spotify", Resp)

	go func() {
		Sv.Start()
	}()

	<-make(chan string)
  	return
}

func messageHandler(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == BotId {
		fmt.Println("Message from the bot")
		return
	}

	if (m.Content == "!greet") {
		output := fmt.Sprintf("Hi there, %s", m.Author.Username)
		_, _ = s.ChannelMessageSend(m.ChannelID, output)
	}

	if (m.Content == "!spotify") {
		s.ChannelMessageSend(m.ChannelID, "Getting spotify info")
		fmt.Println("Your acces token", Resp.Access_token)
		// Spotify.GetProfile(Resp.Access_token)
		Spotify.Auhtorize()
	}
}

