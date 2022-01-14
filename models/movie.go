package models

type Movies struct {
	Results []Movie `json:"results"`
}

// Movie holds a movie details
type Movie struct {
	EpisodeId    int    `json:"episode_id"`
	Title        string `json:"title"`
	OpeningCrawl string `json:"opening_crawl"`
	CommentCount int64  `json:"comment_count"`
	ReleaseDate  string `json:"release_date"`
}

type TypeName interface{}

type Response struct {
	Status   string `json:"status"`
	TypeName `json:"data"`
}
