package dto

type TechnologyResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	DescriptionHTML string `json:"description"`
	Logo            string `json:"logo"`
	Major           string `json:"is_major"`
	CreatedAt       string `json:"created_at"`
}
type TechnologySingleResponse struct {
	ID              uint   `json:"id"`
	Name            string `json:"name"`
	DescriptionHTML string `json:"description"`
	LogoURL         string `json:"logo"`
	IsMajor         string `json:"is_major"`
	CreatedAt       string `json:"created_at"`
}

type TechnologyUpdateSingleResponse struct {
	Name            string `json:"name"`
	DescriptionHTML string `json:"description"`
	LogoURL         string `json:"logo"`
	IsMajor         string `json:"is_major"`
}

type TechnologyDeleteSingleResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	DeletedAt *string `json:"deleted_at"`
}
