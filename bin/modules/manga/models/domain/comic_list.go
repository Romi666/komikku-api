package domain

type (
	Comic struct {
		Title		string	`json:"title"`
		Image		string	`json:"image"`
		Endpoint	string	`json:"endpoint"`
	}

	ComicDetail struct {
		Title				string	`json:"title"`
		FirstChapter		string 	`json:"first_chapter"`
		UrlFirstChapter		string 	`json:"url_first_chapter"`
		LatestChapter		string 	`json:"latest_chapter"`
		UrlLatestChapter 	string 	`json:"url_latest_chapter"`
		Description			string `json:"description"`
		Thumbnail			string `json:"thumbnail"`

	}

	ComicInfo struct {
		Thumbnail	string 		`json:"thumbnail"`
		Title		string 		`json:"title"`
		Type		string 		`json:"type"`
		Author		string 		`json:"author"`
		Status		string 		`json:"status"`
		Rating		string 		`json:"rating"`
		Genre		[]string 	`json:"genre"`
		ChapterList []Chapter 	`json:"chapter_list"`
	}

	Chapter struct {
		Name		string 		`json:"name"`
		Endpoint	string		`json:"endpoint"`
	}

)