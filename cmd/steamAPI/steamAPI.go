package steamapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const SteamAPIURL = "https://store.steampowered.com/api/featuredcategories/?l=english"

func GetSpecialGames() (*Steam, error) {
	response, err := http.Get(SteamAPIURL)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if response.StatusCode == http.StatusOK {
		body, err := io.ReadAll(response.Body)
		if err != nil {
			fmt.Println(err)
		}

		var data *Steam
		err = json.Unmarshal(body, &data)
		if err != nil {
			fmt.Println(err)
		}

		return data, nil

	} else {
		fmt.Println("Error: Couldn't retrieve Steam games list")
		return nil, fmt.Errorf("error getting steam data: %d", response.StatusCode)
	}
}
