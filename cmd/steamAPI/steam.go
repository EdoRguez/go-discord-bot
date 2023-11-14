package steamapi

type Steam struct {
	Status   int `json: "status"`
	Specials []Special
}

type Special struct {
	Items []Item
}

type Item struct {
	Id   int    `json: "id"`
	Name string `json: "name"`
}
