package dto

type ProjectResponse struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	ImageURL      string `json:"image_url"`
	ImageFileName string `json:"image_file_name"`
	RepositoryURL string `json:"repository_url"`
	Summary       string `json:"summary"`
	Status        string `json:"status"`
	CreatedAt     string `json:"created_at"`
}

type ProjectUpdateSingleResponse struct {
	Title         string `json:"title"`
	ImageURL      string `json:"image_url"`
	RepositoryURL string `json:"repository_url"`
	Summary       string `json:"summary"`
	Status        string `json:"status"`
}

type ProjectDeleteSingleResponse struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	DeletedAt *string `json:"deleted_at"`
}

type ProjectUploadResponse struct {
	ImageFileName string `json:"image_file_name"`
	ImageURL      string `json:"image_url"`
}
