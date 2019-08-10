package spotify

type Followers struct {
	HREF  string `json:"href"`
	Total int    `json:"total"`
}

type PrivateUser struct {
	User
	Country string `json:"country"`
	Email   string `json:"email"`
	Product string `json:"product"`
}

type User struct {
	External
	DisplayName string    `json:"display_name"`
	Followers   Followers `json:"followers"`
	Images      []Image   `json:"images"`
}

type Me struct {
	External
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Images      []struct {
		URL    string `json:"url"`
		Width  int    `json:"width"`
		Height int    `json:"height"`
	} `json:"images"`
	Product string `json:"product"`
}
