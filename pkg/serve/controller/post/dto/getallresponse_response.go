package dto

type GetAllPostsResponse struct {
	ID              int64  `json:"id"`
	Title           string `json:"title"`
	Image           string `json:"image"`
	Visibility      string `json:"visibility"`
	ContentMarkdown string `json:"contentMarkdown"`
	ContentHTML     string `json:"contentHtml"`
}
