package bot

import (
	"fmt"

	steamapi "github.com/EdoRguez/go-discord-bot/cmd/steamAPI"
	"github.com/bwmarrin/discordgo"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!help" {
		_, err := s.ChannelMessageSend(m.ChannelID, "")
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SteamSpecials(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "!specials" {

		data, err := steamapi.GetSpecialGames()
		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = s.ChannelMessageSend(m.ChannelID, fmt.Sprintf("%#v", data.Specials.Games[1].Name))
		if err != nil {
			fmt.Println(err)
		}
	}
}
