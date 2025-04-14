package dto

type ProjectRawResponse struct {
	ID            uint
	Title         string
	Status        string
	Summary       string
	ImageURL      string
	RepositoryURL string
	PublishedAt   *string
	TechID        int
	TechName      string
}

type ProjectTechnologyDTO struct {
	ID   int    `json:"tech_id"`
	Name string `json:"tech_name"`
}

type ProjectDTO struct {
	ID            int     `json:"id"`
	Title         string  `json:"title"`
	Status        string  `json:"status"`
	Summary       string  `json:"summary"`
	ImageURL      string  `json:"image_url"`
	RepositoryURL string  `json:"repository_url"`
	PublishedAt   *string `json:"published_at"`
}

type ProjectGetAllDTO struct {
	ID            int                    `json:"id"`
	Title         string                 `json:"title"`
	Status        string                 `json:"status"`
	Summary       string                 `json:"summary"`
	ImageURL      string                 `json:"image_url"`
	RepositoryURL string                 `json:"repository_url"`
	PublishedAt   *string                `json:"published_at"`
	Technologies  []ProjectTechnologyDTO `json:"technologies"`
}

type ProjectOnlyRawResponse struct {
	ID            uint    `json:"project_id"`
	Title         string  `json:"title"`
	Status        string  `json:"status"`
	Summary       string  `json:"summary"`
	ImageURL      string  `json:"image_url"`
	RepositoryURL string  `json:"repository_url"`
	PublishedAt   *string `json:"published_at"`
}

type ProjectTechOnlyRawResponse struct {
	ProjectID uint
	TechID    uint
	TechName  string
}

type ProjectSingleResponse struct {
	ID            uint              `json:"id"`
	Title         string            `json:"title"`
	Description   string            `json:"description"`
	ImageURL      string            `json:"image_url"`
	ImageFileName string            `json:"image_file_name"`
	RepositoryURL string            `json:"repository_url"`
	Summary       string            `json:"summary"`
	Status        string            `json:"status"`
	StatisticId   uint              `json:"statistic_id"`
	PublishedAt   *string           `json:"published_at"`
	CreatedAt     string            `json:"created_at"`
	Statistic     StatisticResponse `json:"statistic"`
}

type ProjectCreateResponse struct {
	ID            uint   `json:"id"`
	Title         string `json:"title"`
	Description   string `json:"description"`
	ImageURL      string `json:"image_url"`
	ImageFileName string `json:"image_file_name"`
	RepositoryURL string `json:"repository_url"`
	Summary       string `json:"summary"`
	Status        string `json:"status"`
	StatisticId   int    `json:"statistic_id"`
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
