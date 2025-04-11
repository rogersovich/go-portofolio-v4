package dto

type AboutQueryParams struct {
	Sort        string
	Order       string
	Title       string
	Description string
	IsDelete    string // "Y" or "N"
	CreatedFrom string
	CreatedTo   string
	Page        int
	Limit       int
}

type CreateAboutRequest struct {
	Title           string      `json:"title" validate:"required"`
	DescriptionHTML string      `json:"description" validate:"required"`
	AvatarFile      interface{} `json:"avatar_file"`
}

type UpdateAboutRequest struct {
	Id              int    `json:"id" validate:"required"`
	Title           string `json:"title" validate:"required"`
	DescriptionHTML string `json:"description" validate:"required"`
}

type UpdateAboutPayload struct {
	Title           string `json:"title"`
	DescriptionHTML string `json:"description"`
	AvatarURL       string `json:"avatar_url"`
	AvatarFileName  string `json:"avatar_file_name"`
}

type DeleteAboutRequest struct {
	ID int `json:"id" binding:"required"`
}
