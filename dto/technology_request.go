package dto

type TechnologyQueryParams struct {
	Sort        string
	Order       string
	FilterName  string
	FilterDesc  string
	IsMajor     string // "Y" or "N"
	IsDelete    string // "Y" or "N"
	CreatedFrom string
	CreatedTo   string
	Page        int
	Limit       int
}

type CreateTechnologyRequest struct {
	Name            string      `json:"name" validate:"required"`
	DescriptionHTML string      `json:"description" validate:"required"`
	LogoFile        interface{} `json:"logo_file"`
	IsMajor         string      `json:"is_major" validate:"required,oneof=Y N"` // use Y or N
}

type UpdateTechnologyRequest struct {
	Id              int    `json:"id" validate:"required"`
	Name            string `json:"name" validate:"required"`
	DescriptionHTML string `json:"description" validate:"required"`
	IsMajor         string `json:"is_major" validate:"oneof=Y N"` // use Y or N
}

type DeleteTechnologyRequest struct {
	ID int `json:"id" binding:"required"`
}
