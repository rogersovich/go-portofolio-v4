package dto

type ProjectFeatureResponse struct {
	ID            uint    `json:"id"`
	Description   *string `json:"description"`
	ImageURL      string  `json:"image_url"`
	ImageFileName *string `json:"image_file"`
	CreatedAt     string  `json:"created_at"`
}
type ProjectFeatureSingleResponse struct {
	ID            uint    `json:"id"`
	Description   *string `json:"description"`
	ImageURL      string  `json:"image_url"`
	ImageFileName *string `json:"image_file"`
	CreatedAt     string  `json:"created_at"`
}

type ProjectFeatureUpdateSingleResponse struct {
	Description *string `json:"description"`
	ImageURL    string  `json:"image_url"`
}

type ProjectFeatureUpdateResponse struct {
	Description   *string `json:"description"`
	ImageURL      string  `json:"image_url"`
	ImageFileName *string `json:"image_file"`
}

type ProjectFeatureDeleteSingleResponse struct {
	ID        int     `json:"id"`
	DeletedAt *string `json:"deleted_at"`
}

type ProjectFeatureUploadResponse struct {
	ImageFileName string `json:"image_file_name"`
	ImageURL      string `json:"image_url"`
}
