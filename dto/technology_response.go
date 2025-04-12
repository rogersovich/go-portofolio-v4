package dto

type TechnologyResponse struct {
	ID              uint    `json:"id"`
	Name            string  `json:"name"`
	DescriptionHTML string  `json:"description"`
	LogoURL         string  `json:"logo_url"`
	LogoFileName    *string `json:"logo_file_name"`
	Major           string  `json:"is_major"`
	CreatedAt       string  `json:"created_at"`
}
type TechnologySingleResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	DescriptionHTML string `json:"description"`
	LogoURL         string `json:"logo_url"`
	LogoFileName    string `json:"logo_file_name"`
	IsMajor         string `json:"is_major"`
	CreatedAt       string `json:"created_at"`
}

type TechnologyUpdateSingleResponse struct {
	Name            string `json:"name"`
	DescriptionHTML string `json:"description"`
	LogoURL         string `json:"logo_url"`
	LogoFileName    string `json:"logo_file_name"`
	IsMajor         string `json:"is_major"`
}

type TechnologyDeleteSingleResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	DeletedAt *string `json:"deleted_at"`
}
