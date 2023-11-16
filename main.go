package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/EdoRguez/go-discord-bot/cmd/bot"
	"github.com/EdoRguez/go-discord-bot/pkg/db"
	"github.com/bwmarrin/discordgo"
)

var (
	Token       string
	redisClient = db.NewRedisClient()
)

func init() {
	flag.StringVar(&Token, "t", os.Getenv("DISCORD_BOT_TOKEN"), "Bot Token")
	flag.Parse()
}

func main() {
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	registerBotHandlers(dg)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

func registerBotHandlers(dg *discordgo.Session) {
	dg.AddHandler(bot.Help)
	dg.AddHandler(bot.SteamSpecials)
}
