package dto

type CreateProjectRequest struct {
	Title         string   `json:"title" validate:"required"`
	Description   string   `json:"description" validate:"required"`
	ImageFile     string   `json:"image_file"`
	RepositoryURL string   `json:"repository_url"`
	Summary       string   `json:"summary" validate:"required"`
	Status        string   `json:"status" validate:"oneof=Published Unpublished Deleted"`
	IsPublihed    string   `json:"is_published" validate:"oneof=Y N"`
	TechnologyIds []string `json:"technology_ids"`
	ContentImages []string `json:"content_images"`
}

type UpdateProjectRequest struct {
	Id            int    `json:"id" validate:"required"`
	Title         string `json:"title" validate:"required"`
	Description   string `json:"description" validate:"required"`
	RepositoryURL string `json:"repository_url"`
	Summary       string `json:"summary"`
	Status        string `json:"status" validate:"oneof=Published Unpublished Deleted"`
	IsPublihed    string `json:"is_published" validate:"oneof=Y N"`
}

type UpdateProjectPayload struct {
	Title         string `json:"title"`
	Description   string `json:"description" validate:"required"`
	RepositoryURL string `json:"repository_url"`
	Summary       string `json:"summary"`
	Status        string `json:"status"`
	IsPublihed    string `json:"is_published" `
	ImageURL      string `json:"image_url"`
	ImageFileName string `json:"image_file_name"`
}

type DeleteProjectRequest struct {
	ID int `json:"id" binding:"required"`
}
