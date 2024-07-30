package models

type Languages struct {
	HTML   string `json:"html"`
	CSS    string `json:"css"`
	NODEJS string `json:"nodejs"`
	JAVA   string `json:"java"`
}

type Skill struct {
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Languages Languages `json:"languages"`
	Cv        string    `json:"cv"`
}
