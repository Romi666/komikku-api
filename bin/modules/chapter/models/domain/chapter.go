package domain

type ChapterDetail struct {
	Title	string 		`json:"title"`
	Image	[]string	`json:"image"`
}

type Chapter struct {
Name		string 		`json:"name"`
Endpoint	string		`json:"endpoint"`
}