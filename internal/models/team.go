package models

type Team struct {
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Lead  string `json:"lead"`
	Users struct {
		Uuid       string `json:"uuid"`
		Name       string `json:"name"`
		Patronymic string `json:"patronymic"`
		Surname    string `json:"surname"`
	} `json:"users"`
}
