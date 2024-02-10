package bot

import (
	"github.com/bwmarrin/discordgo"
)

type Handlers struct {
	Specials       *Specials
	DefaultChannel string
}

func (h Handlers) Help(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID ||
		m.ChannelID != h.DefaultChannel {
		return
	}

	if m.Content == "!help" {
		SendHelp(s, m)
	}
}

func (h Handlers) SteamSpecials(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID ||
		m.ChannelID != h.DefaultChannel {
		return
	}

	switch m.Content {
	case "!specials":
		h.Specials.SendSpecials(s, m)
	case "!specials-start":
		h.Specials.StartDailySpecials(s, m, "")
	case "!specials-stop":
		h.Specials.StopDailySpecials(s, m)
	}
}
