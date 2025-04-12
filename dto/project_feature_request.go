package dto

type ProjectFeatureQueryParams struct {
	Sort        string
	Order       string
	Description string
	IsDelete    string // "Y" or "N"
	CreatedFrom string
	CreatedTo   string
	Page        int
	Limit       int
}

type CreateProjectFeatureRequest struct {
	Description string      `json:"description" validate:"required"`
	ImageFile   interface{} `json:"image_file"`
}

type UpdateProjectFeatureRequest struct {
	Id          int    `json:"id" validate:"required"`
	Description string `json:"description" validate:"required"`
}

type UpdateProjectFeaturePayload struct {
	Title         string `json:"title"`
	Description   string `json:"description"`
	ImageURL      string `json:"image_url"`
	ImageFileName string `json:"image_file_name"`
}

type DeleteProjectFeatureRequest struct {
	ID int `json:"id" binding:"required"`
}
