package models

type Solution struct {
	Design     string `json:"design"`
	Responsive string `json:"responsive"`
	Sitemap    string `json:"sitemap"`
	Feature    string `json:"feature"`
	Feedback   string `json:"feedback"`
}

type Detail struct {
	ProjectId string   `json:"projectId"`
	Client    string   `json:"client"`
	Objective string   `json:"objective"`
	Industry  string   `json:"industry"`
	Duration  string   `json:"duration"`
	Role      string   `json:"role"`
	Workflow  string   `json:"workflow"`
	Solution  Solution `json:"solution"`
}

type Project struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Image  string `json:"image"`
	Detail Detail `json:"detail"`
}
