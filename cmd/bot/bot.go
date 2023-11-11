package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

type Bot struct {
	Status int `json: "status"`
}

func (b *Bot) MessageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Ignore all messages created by the bot itself
	// This isn't required in this specific example but it's a good practice.
	if m.Author.ID == s.State.User.ID {
		return
	}

	// if m.Content == "!gopher" {
	// 	fmt.Println("entro mano")
	// 	s.ChannelMessageSend(m.ChannelID, "hjajajajajaja")
	// 	//Call the KuteGo API and retrieve our cute Dr Who Gopher
	// 	response, err := http.Get(SteamAPIURL)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	defer response.Body.Close()

	// 	if response.StatusCode == 200 {
	// 		// _, err = s.ChannelFileSend(m.ChannelID, "dr-who.png", response.Body)
	// 		// if err != nil {
	// 		// 	fmt.Println(err)
	// 		// }
	// 		s.ChannelMessageSend(m.ChannelID, response.Body)
	// 	} else {
	// 		fmt.Println("Error: Can't get dr-who Gopher! :-(")
	// 	}
	// }

	// if m.Content == "!random" {

	// 	//Call the KuteGo API and retrieve a random Gopher
	// 	response, err := http.Get(SteamAPIURL)
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	defer response.Body.Close()

	// 	if response.StatusCode == 200 {
	// 		_, err = s.ChannelFileSend(m.ChannelID, "random-gopher.png", response.Body)
	// 		if err != nil {
	// 			fmt.Println(err)
	// 		}
	// 	} else {
	// 		fmt.Println("Error: Can't get random Gopher! :-(")
	// 	}
	// }

	if m.Content == "!gophers" {

		//Call the KuteGo API and display the list of available Gophers
		response, err := http.Get(SteamAPIURL)
		if err != nil {
			fmt.Println(err)
		}
		defer response.Body.Close()

		if response.StatusCode == 200 {
			// Transform our response to a []byte
			body, err := ioutil.ReadAll(response.Body)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(body)

			// Put only needed informations of the JSON document in our array of Gopher
			var data Steam
			err = json.Unmarshal(body, &data)
			if err != nil {
				fmt.Println(err)
			}

			// Create a string with all of the Gopher's name and a blank line as separator
			//var gophers strings.Builder
			// for _, gopher := range data {
			// 	gophers.WriteString("Status = " + gopher.Status + "\n")
			// }

			fmt.Println(data)

			// Send a text message with the list of Gophers
			_, err = s.ChannelMessageSend(m.ChannelID, strconv.Itoa(data.Status))
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error: Can't get list of Gophers! :-(")
		}
	}
}
