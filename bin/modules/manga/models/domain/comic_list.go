package domain

type (
	Comic struct {
		Title		string	`json:"title"`
		Image		string	`json:"image"`
		Endpoint	string	`json:"endpoint"`
	}
)