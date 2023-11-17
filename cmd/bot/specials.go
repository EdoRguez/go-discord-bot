package bot

import (
	"fmt"

	steamapi "github.com/EdoRguez/go-discord-bot/cmd/steamAPI"
	"github.com/EdoRguez/go-discord-bot/pkg/db"
	"github.com/EdoRguez/go-discord-bot/pkg/util"
	"github.com/bwmarrin/discordgo"
)

var (
	rc = db.NewRedisClient()
)

func sendSpecials(s *discordgo.Session, m *discordgo.MessageCreate) {
	data, err := steamapi.GetSpecialGames()
	if err != nil {
		fmt.Println(err)
		s.ChannelMessageSend(m.ChannelID, err.Error())
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

func startDailySpecials(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "> Daily Steam specials started")

	error := discordCronJob.StartCronJob("@every 10s", func() {

		s.ChannelMessageSend(m.ChannelID, "Hola")
		// data, err := steamapi.GetSpecialGames()
		// if err != nil {
		// 	fmt.Println(err)
		// 	return
		// }

		// storage := storage.NewStorage(rc)
		// for _, game := range data.Specials.Games {

		// 	if _, err := storage.GetRedis(strconv.Itoa(game.Id)); err != nil {

		// 		message := util.FormatGameDiscountMessage(game)
		// 		_, err = s.ChannelMessageSend(m.ChannelID, message)
		// 		if err != nil {
		// 			fmt.Println(err)
		// 		}

		// 		storage.SaveRedis(strconv.Itoa(game.Id), game.Name)
		// 	}
		// }
	})
	if error != nil {
		s.ChannelMessageSend(m.ChannelID, "> Daily Steam Specials is already running")
	}
}

func stopDailySpecials(s *discordgo.Session, m *discordgo.MessageCreate) {
	s.ChannelMessageSend(m.ChannelID, "> Daily Steam specials stopped")
	discordCronJob.StopCronJobs()
}
