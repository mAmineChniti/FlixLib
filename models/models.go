package models

type Movie struct {
	TitleText struct {
		Text string `json:"text"`
	} `json:"titleText"`
	PrimaryImage struct {
		URL string `json:"url"`
	} `json:"primaryImage"`
}

type APIResponse struct {
	Results []Movie `json:"results"`
}

type MovieData struct {
	Title    string
	ImageURL string
}
