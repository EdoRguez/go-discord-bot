package bot

import (
	"github.com/bwmarrin/discordgo"
)

var (
	sp                    = NewSpecials()
	defaultDiscordChannel = "1175188918791057528"
)

func Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID ||
		m.ChannelID != defaultDiscordChannel {
		return
	}

	if m.Content == "!help" {
		SendHelp(s, m)
	}
}

func SteamSpecials(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID ||
		m.ChannelID != defaultDiscordChannel {
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
