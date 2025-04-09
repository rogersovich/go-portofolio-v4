package services

import (
	"fmt"

	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/models"
	"github.com/rogersovich/go-portofolio-v4/utils"
	"github.com/sirupsen/logrus"
)

func GetAllTechnologies() ([]models.Technology, error) {
	db, _ := config.DB.DB()

	query := `SELECT id, name, logo_url, description_html, is_major FROM technologies`

	rows, err := db.Query(query)
	if err != nil {
		utils.LogError(err.Error(), query)
		return nil, fmt.Errorf("query error: %w", err)
	}
	defer rows.Close()

	var technologies []models.Technology

	for rows.Next() {
		var tech models.Technology
		if err := rows.Scan(&tech.ID, &tech.Name, &tech.LogoURL, &tech.DescriptionHTML, &tech.IsMajor); err != nil {
			utils.LogWarning(err.Error(), query)
			return nil, err
		}
		technologies = append(technologies, tech)
	}

	utils.Log.WithFields(logrus.Fields{
		"query": query,
		"count": len(technologies),
	}).Debug("✅ Fetched technologies from DB")

	return technologies, nil
}
