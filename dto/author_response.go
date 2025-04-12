package dto

type AuthorResponse struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	AvatarURL      string  `json:"avatar_url"`
	AvatarFileName *string `json:"avatar_file"`
	CreatedAt      string  `json:"created_at"`
}
type AuthorSingleResponse struct {
	ID             uint    `json:"id"`
	Name           string  `json:"name"`
	AvatarURL      string  `json:"avatar_url"`
	AvatarFileName *string `json:"avatar_file"`
	CreatedAt      string  `json:"created_at"`
}

type AuthorUpdateSingleResponse struct {
	Name      string `json:"name"`
	AvatarURL string `json:"avatar_url"`
}

type AuthorUpdateResponse struct {
	Name           string  `json:"name"`
	AvatarURL      string  `json:"avatar_url"`
	AvatarFileName *string `json:"avatar_file"`
}

type AuthorDeleteSingleResponse struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	DeletedAt *string `json:"deleted_at"`
}

type AuthorUploadResponse struct {
	AvatarFileName string `json:"avatar_file_name"`
	AvatarURL      string `json:"avatar_url"`
}
