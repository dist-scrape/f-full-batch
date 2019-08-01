package model

type Response struct {
	OEMs []OEM `json:"makeModels"`
}

type OEM struct {
	Title  string  `json:"make"`
	Models []Model `json:"children"`
}

type Model struct {
	Title string `json:"model"`
}
