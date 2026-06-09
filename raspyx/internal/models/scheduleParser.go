package models

type ParseGroupScheduleData struct {
	Group     string
	IsSession int
}

type Lesson struct {
	Sbj      string `json:"sbj"`
	Teachers []struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
	} `json:"teachers"`
	Dts        string `json:"dts"`
	Df         string `json:"df"`
	Dt         string `json:"dt"`
	Auditories []struct {
		Title string `json:"title"`
		Color string `json:"color"`
	} `json:"auditories"`
	ShortRooms []string `json:"shortRooms"`
	Location   string   `json:"location"`
	Type       string   `json:"type"`
	Week       string   `json:"week"`
	Align      string   `json:"align"`
	ELink      any      `json:"e_link"`
}

type ParseScheduleResponse struct {
	Status    string                         `json:"status"`
	Message   string                         `json:"message"`
	Grid      map[string]map[string][]Lesson `json:"grid"`
	IsSession bool                           `json:"isSession"`
}
