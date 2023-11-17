package bot

import (
	"fmt"
	"strconv"

	steamapi "github.com/EdoRguez/go-discord-bot/cmd/steamAPI"
	"github.com/EdoRguez/go-discord-bot/internal/cronjob"
	"github.com/EdoRguez/go-discord-bot/internal/storage"
	"github.com/EdoRguez/go-discord-bot/pkg/db"
	"github.com/EdoRguez/go-discord-bot/pkg/util"
	"github.com/bwmarrin/discordgo"
	"github.com/redis/go-redis/v9"
)

type Specials struct {
	rc             *redis.Client
	discordCronJob cronjob.CronJob
}

func (sp *Specials) SendSpecials(s *discordgo.Session, m *discordgo.MessageCreate) {
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

func (sp *Specials) StartDailySpecials(s *discordgo.Session, m *discordgo.MessageCreate) {

	// Every day at 5am
	error := sp.discordCronJob.StartCronJob("0 5 * * *", func() {
		storage := storage.NewStorage(sp.rc)

		if err := storage.CheckRedisConnection(); err != nil {
			s.ChannelMessageSend(m.ChannelID, "> Redis DB Connection Error")
			return
		}

		data, err := steamapi.GetSpecialGames()
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, game := range data.Specials.Games {

			if _, err := storage.GetRedis(strconv.Itoa(game.Id)); err != nil {

				message := util.FormatGameDiscountMessage(game)
				_, err = s.ChannelMessageSend(m.ChannelID, message)
				if err != nil {
					fmt.Println(err)
				}

				storage.SaveRedis(strconv.Itoa(game.Id), game.Name)
			}
		}
	})
	if error != nil {
		s.ChannelMessageSend(m.ChannelID, "> Daily Steam Specials is **ALREADY RUNNING**")
		return
	}
	s.ChannelMessageSend(m.ChannelID, "> Daily Steam specials **STARTED**")
}

func (sp *Specials) StopDailySpecials(s *discordgo.Session, m *discordgo.MessageCreate) {
	sp.discordCronJob.StopCronJobs()
	s.ChannelMessageSend(m.ChannelID, "> Daily Steam specials **STOPPED**")
}

func NewSpecials() *Specials {
	return &Specials{
		rc: db.NewRedisClient(),
		discordCronJob: cronjob.CronJob{
			CronJob: cronjob.NewCronJob(),
		},
	}
}
