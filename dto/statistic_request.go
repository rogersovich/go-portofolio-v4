package dto

type StatisticQueryParams struct {
	Sort        string
	Order       string
	Type        string // "Project" and "Blog"
	IsDelete    string // "Y" or "N"
	CreatedFrom string
	CreatedTo   string
	Page        int
	Limit       int
}

type CreateStatisticRequest struct {
	Likes *int   `json:"likes" binding:"required"`
	Views *int   `json:"views" binding:"required"`
	Type  string `json:"type" binding:"required,oneof=Project Blog"`
}

type UpdateStatisticRequest struct {
	Id    int    `json:"id" binding:"required,numeric"`
	Likes *int   `json:"likes" binding:"required"`
	Views *int   `json:"views" binding:"required"`
	Type  string `json:"type" binding:"required,oneof=Project Blog"`
}

type DeleteStatisticRequest struct {
	ID int `json:"id" binding:"required"`
}
