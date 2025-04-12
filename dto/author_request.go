package dto

type AuthorQueryParams struct {
	Sort        string
	Order       string
	Name        string
	IsDelete    string // "Y" or "N"
	CreatedFrom string
	CreatedTo   string
	Page        int
	Limit       int
}

type CreateAuthorRequest struct {
	Name       string      `json:"name" validate:"required"`
	AvatarFile interface{} `json:"avatar_file"`
}

type UpdateAuthorRequest struct {
	Id   int    `json:"id" validate:"required"`
	Name string `json:"name" validate:"required"`
}

type UpdateAuthorPayload struct {
	Name           string `json:"name"`
	AvatarURL      string `json:"avatar_url"`
	AvatarFileName string `json:"avatar_file_name"`
}

type DeleteAuthorRequest struct {
	ID int `json:"id" binding:"required"`
}
