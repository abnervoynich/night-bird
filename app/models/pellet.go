package models

type Pellet struct {
	ID          int             `json:"id"`
	Name        string          `json:"name"`
	Url         string          `json:"url"`
	Request     *RequestPellet  `json:"request"`
	Response    *ResponsePellet `json:"response"`
	HttpMethod  httpMethod      `json:"http-method"`
	ProjectName string          `json:"project-name"`
	IsActive    bool            `json:"is-active" `
}
