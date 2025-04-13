package dto

type CreateProjectContentImageRequest struct {
	ProjectId *string `json:"project_id" `
	IsUsed    string  `json:"is_used" validate:"required,oneof=Y N"`
	ImageFile *string `json:"image_file" validate:"required"`
}

type UpdateProjectContentImageRequest struct {
	Id        int    `json:"id" validate:"required"`
	ProjectId string `json:"project_id"  validate:"required"`
	IsUsed    string `json:"is_used" validate:"required,oneof=Y N"`
}

type UpdateProjectContentImagePayload struct {
	ProjectId     string `json:"project_id"`
	IsUsed        string `json:"is_used"`
	ImageURL      string `json:"image_url"`
	ImageFileName string `json:"image_file_name"`
}

type DeleteProjectContentImageRequest struct {
	ID int `json:"id" binding:"required"`
}

type ProjectContentImageResponse struct {
	ID            uint    `json:"id"`
	ProjectId     *int    `json:"project_id"`
	ProjectName   *string `json:"project_name"`
	IsUsed        string  `json:"is_used"`
	ImageURL      string  `json:"image_url"`
	ImageFileName string  `json:"image_file_name"`
	CreatedAt     string  `json:"created_at"`
}

type ProjectContentImageUpdateResponse struct {
	ProjectId     string `json:"project_id"`
	IsUsed        string `json:"is_used"`
	ImageURL      string `json:"image_url"`
	ImageFileName string `json:"image_file_name"`
}

type ProjectContentImageDeleteSingleResponse struct {
	ID        int     `json:"id"`
	ProjectId string  `json:"project_id"`
	IsUsed    string  `json:"is_used"`
	DeletedAt *string `json:"deleted_at"`
}

type ProjectContentImageUploadResponse struct {
	ImageURL      string `json:"image_url"`
	ImageFileName string `json:"image_file_name"`
}
