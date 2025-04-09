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
	Name            string `json:"name" binding:"required"`
	DescriptionHTML string `json:"description" binding:"required"`
	LogoURL         string `json:"logo" binding:"required"`
	IsMajor         string `json:"is_major" binding:"required,oneof=Y N"` // use Y or N
}
