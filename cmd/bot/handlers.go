package bot

import (
	"fmt"

	steamapi "github.com/EdoRguez/go-discord-bot/cmd/steamAPI"
	"github.com/EdoRguez/go-discord-bot/pkg/util"
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

		for _, game := range data.Specials.Games {
			message := util.FormatGameDiscountMessage(game)
			_, err = s.ChannelMessageSend(m.ChannelID, message)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
}
