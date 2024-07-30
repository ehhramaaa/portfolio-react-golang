package models

type SocialMedia struct {
	Github   string `json:"github"`
	Twitter  string `json:"twitter"`
	LinkedIn string `json:"linkedin"`
}

type Home struct {
	Desc        string      `json:"desc"`
	Experience  string      `json:"experience"`
	SocialMedia SocialMedia `json:"socialMedia"`
}
