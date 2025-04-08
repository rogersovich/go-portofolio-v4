package services

import (
	"fmt"

	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/models"
)

func GetAllTechnologies() ([]models.Technology, error) {
	db, err := config.DB.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get raw DB: %w", err)
	}

	rows, err := db.Query(`SELECT id, name, logo_url, description_html, is_major FROM technologies`)
	if err != nil {
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var technologies []models.Technology

	for rows.Next() {
		var tech models.Technology
		if err := rows.Scan(&tech.ID, &tech.Name, &tech.LogoURL, &tech.DescriptionHTML, &tech.IsMajor); err != nil {
			return nil, err
		}
		technologies = append(technologies, tech)
	}

	return technologies, nil
}
