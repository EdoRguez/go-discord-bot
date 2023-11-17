package bot

import (
	"fmt"

	"github.com/EdoRguez/go-discord-bot/internal/cronjob"
	"github.com/bwmarrin/discordgo"
)

var (
	discordCronJob = cronjob.CronJob{
		CronJob: cronjob.NewCronJob(),
	}
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

	switch m.Content {
	case "!specials":
		sendSpecials(s, m)
	case "!specials-start":
		startDailySpecials(s, m)
	case "!specials-stop":
		stopDailySpecials(s, m)
	}
}
