package link

type LinkCreateRequest struct {
	Url string `json:"url" validate:"required,url"`
}

type LinkEditRequest struct {
	Url  string `json:"url" validate:"required,url"`
	Hash string `json:"hash"`
}

type FindAllResponse struct {
	Links []Link `json:"links"`
	Count int64  `json:"count"`
}
