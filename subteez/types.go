package subteez

type SearchRequest struct {
	Query    string   `form:"query" json:"query" uri:"query" binding:"required"`
	Language []string `form:"lang" json:"lang" uri:"lang" binding:"required"`
}

type SubtitleDetailsRequest struct {
	ID       string   `form:"id" json:"id" uri:"id" binding:"required"`
	Language []string `form:"lang" json:"lang" uri:"lang" binding:"required"`
}

type SubtitleDownloadRequest struct {
	ID string `form:"id" json:"id" uri:"id" binding:"required"`
}

type SearchResult struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type SubtitleFile struct {
	ID       string `json:"id"`
	Language string `json:"lang"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Comment  string `json:"comment"`
	Title    string `json:"title"`
}

type SubtitleDetails struct {
	Status string         `json:"status"`
	Name   string         `json:"name"`
	Year   string         `json:"year"`
	Banner interface{}    `json:"posterUrl"`
	Files  []SubtitleFile `json:"files"`
}