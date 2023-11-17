package util

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	steamapi "github.com/EdoRguez/go-discord-bot/cmd/steamAPI"
)

const baseSteamStoreURL = "https://store.steampowered.com/app/"

func FormatGameDiscountMessage(game steamapi.Game) string {
	var result strings.Builder

	// Game Title
	result.WriteString(fmt.Sprintf("> # %s \n", game.Name))

	// Game New Price
	newPrice := float32(game.FinalPrice) / float32(100)
	result.WriteString(fmt.Sprintf("> - New Price         ⟶  **%.2f USD** \n", newPrice))

	// Game Original Price
	originalPrice := float32(game.OriginalPrice) / float32(100)
	result.WriteString(fmt.Sprintf("> - Original Price  ⟶  ~~%.2f USD~~ \n", originalPrice))

	// Game Discount Percentage
	result.WriteString(fmt.Sprintf("> - Discount           ⟶  **%v%%** \n", game.DiscountPercent))

	// Game Discount Expiration
	result.WriteString(fmt.Sprintf("> - End Date          ⟶  %v \n\n", TimeUnixtToDate(game.DiscountExpirationTimeStamp)))

	// Game URL
	result.WriteString(fmt.Sprintf("%s \n\n", baseSteamStoreURL+strconv.Itoa(game.Id)))

	return result.String()
}

func TimeUnixtToDate(timeUnit int64) string {
	t := time.Unix(int64(timeUnit), 0).UTC()
	return t.Format("02/01/2006")
}
