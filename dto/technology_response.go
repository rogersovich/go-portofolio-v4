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
