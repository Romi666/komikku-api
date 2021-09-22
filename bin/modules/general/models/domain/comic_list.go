package domain

type (
	Comic struct {
		Title		string	`json:"title"`
		Genre 		string	`json:"genre"`
		Image		string	`json:"image"`
		Status		string	`json:"status"`
		Endpoint	string	`json:"endpoint"`
		LastUpdate	string	`json:"last_update"`
	}
)