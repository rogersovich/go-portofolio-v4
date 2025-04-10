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
	Title           string      `json:"title" binding:"required"`
	DescriptionHTML string      `json:"description" binding:"required"`
	AvatarFile      interface{} `json:"avatar_file" binding:"required"` // Assuming File is represented as an interface{}
}

type UpdateAboutRequest struct {
	Title           string `json:"title" binding:"required"`
	DescriptionHTML string `json:"description" binding:"required"`
	AvatarFile      string `json:"avatar_file" binding:"required"`
}

type DeleteAboutRequest struct {
	ID int `json:"id" binding:"required"`
}
