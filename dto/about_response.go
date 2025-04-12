package dto

type AboutResponse struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	DescriptionHTML *string `json:"description"`
	AvatarURL       string  `json:"avatar_url"`
	AvatarFileName  *string `json:"avatar_file"`
	CreatedAt       string  `json:"created_at"`
}
type AboutSingleResponse struct {
	ID              uint    `json:"id"`
	Title           string  `json:"title"`
	DescriptionHTML *string `json:"description"`
	AvatarURL       string  `json:"avatar_url"`
	AvatarFileName  *string `json:"avatar_file"`
	CreatedAt       string  `json:"created_at"`
}

type AboutUpdateSingleResponse struct {
	Title           string  `json:"title"`
	DescriptionHTML *string `json:"description"`
	AvatarURL       string  `json:"avatar_url"`
}

type AboutUpdateResponse struct {
	Title           string  `json:"title"`
	DescriptionHTML *string `json:"description"`
	AvatarURL       string  `json:"avatar_url"`
	AvatarFileName  *string `json:"avatar_file"`
}

type AboutDeleteSingleResponse struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	DeletedAt *string `json:"deleted_at"`
}

type AboutUploadResponse struct {
	AvatarFileName string `json:"avatar_file_name"`
	AvatarURL      string `json:"avatar_url"`
}
