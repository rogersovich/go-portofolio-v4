package dto

type ProjectTechnologyResponse struct {
	ID           uint   `json:"id"`
	ProjectID    int    `json:"project_id"`
	TechnologyID int    `json:"technology_id"`
	CreatedAt    string `json:"created_at"`
}

type ProjectTechnologySingleResponse struct {
	ID           uint   `json:"id"`
	ProjectID    int    `json:"project_id"`
	TechnologyID int    `json:"technology_id"`
	CreatedAt    string `json:"created_at"`
}

type ProjectTechnologyUpdateResponse struct {
	ProjectID    int `json:"project_id"`
	TechnologyID int `json:"technology_id"`
}

type ProjectTechnologyDeleteSingleResponse struct {
	ID           int     `json:"id"`
	ProjectID    int     `json:"project_id"`
	TechnologyID int     `json:"technology_id"`
	DeletedAt    *string `json:"deleted_at"`
}
