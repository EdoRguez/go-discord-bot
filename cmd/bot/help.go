package bot

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func SendHelp(s *discordgo.Session, m *discordgo.MessageCreate) {
	var message strings.Builder

	message.WriteString("> ## List of Commands \n")
	message.WriteString("> - **!specials**               ⟶  Gets all current Steam games specials. \n")
	message.WriteString("> - **!specials-start**   ⟶  Starts daily process to get current Steam games specials at 5am, this action only displays games that haven't been sent before through chat in order to avoid spam. \n")
	message.WriteString("> - **!specials-stop**    ⟶  Stops daily process.  \n")

	s.ChannelMessageSend(m.ChannelID, message.String())
}
