package bot

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var sp = NewSpecials()

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

	switch m.Content {
	case "!specials":
		sp.SendSpecials(s, m)
	case "!specials-start":
		sp.StartDailySpecials(s, m)
	case "!specials-stop":
		sp.StopDailySpecials(s, m)
	}
}
