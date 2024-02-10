package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/EdoRguez/go-discord-bot/cmd/bot"
	"github.com/EdoRguez/go-discord-bot/internal/config"
	"github.com/bwmarrin/discordgo"
)

var (
	Conf config.Config
)

func init() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed loading config", err)
	}

	Conf = c
}

func main() {
	dg, err := discordgo.New("Bot " + Conf.TOKEN)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	h := bot.Handlers{
		Specials:       bot.NewSpecials(),
		DefaultChannel: Conf.DEFAULT_CHANNEL,
	}

	registerBotHandlers(dg, h)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")

	initDefaultHandlers(dg, h)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func registerBotHandlers(dg *discordgo.Session, h bot.Handlers) {
	dg.AddHandler(h.Help)
	dg.AddHandler(h.SteamSpecials)

}

func initDefaultHandlers(dg *discordgo.Session, h bot.Handlers) {
	for _, guild := range dg.State.Guilds {
		if guild.ID == Conf.DISCORD_SERVER_ID {
			h.Specials.StartDailySpecials(dg, nil, Conf.DEFAULT_CHANNEL)
			break
		}
	}
}
