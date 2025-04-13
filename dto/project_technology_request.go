package dto

type CreateProjectTechnologyRequest struct {
	ProjectID    int `json:"project_id" binding:"required"`
	TechnologyID int `json:"technology_id" binding:"required"`
}

type UpdateProjectTechnologyRequest struct {
	Id           int `json:"id" binding:"required,numeric"`
	ProjectID    int `json:"project_id" binding:"required"`
	TechnologyID int `json:"technology_id" binding:"required"`
}

type DeleteProjectTechnologyRequest struct {
	ID int `json:"id" binding:"required"`
}
