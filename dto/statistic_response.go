package dto

type StatisticResponse struct {
	ID        uint   `json:"id"`
	Likes     int    `json:"likes"`
	Views     int    `json:"views"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
}
type StatisticSingleResponse struct {
	ID        uint   `json:"id"`
	Likes     int    `json:"likes"`
	Views     int    `json:"views"`
	Type      string `json:"type"`
	CreatedAt string `json:"created_at"`
}

type StatisticUpdateResponse struct {
	Likes int    `json:"likes"`
	Views int    `json:"views"`
	Type  string `json:"type"`
}

type StatisticDeleteSingleResponse struct {
	ID        int     `json:"id"`
	Likes     int     `json:"likes"`
	Views     int     `json:"views"`
	Type      string  `json:"type"`
	DeletedAt *string `json:"deleted_at"`
}
