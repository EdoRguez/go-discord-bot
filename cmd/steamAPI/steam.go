package steamapi

type Steam struct {
	Specials Special
}

type Special struct {
	Games []Game `json:"items"`
}

type Game struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Discounted      bool   `json:"discounted"`
	DiscountPercent int    `json:"discount_percent"`
	OriginalPrice   int    `json:"original_price"`
	FinalPrice      int    `json:"final_price"`
	Currency        string `json:"currency"`
}
