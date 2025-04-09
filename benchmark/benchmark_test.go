package main

import (
	"database/sql"
	"testing"

	"github.com/joho/godotenv"
	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/models"
	"gorm.io/gorm"
)

var (
	sqlDB  *sql.DB
	gormDB *gorm.DB
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		_ = godotenv.Load("../.env") // fallback if not found in local dir
	}

	config.InitConfigForTest()
	config.ConnectDB()

	gormDB = config.DB
	sqlDB, err = gormDB.DB()
	if err != nil {
		panic("Failed to get *sql.DB")
	}
}

func BenchmarkGORM_RawScan(b *testing.B) {
	main()
	for i := 0; i < b.N; i++ {
		var techs []models.Technology
		err := gormDB.Raw(`SELECT id, name, logo_url, description_html, is_major FROM technologies`).Scan(&techs).Error
		if err != nil {
			b.Error(err)
		}
	}
}

func BenchmarkSQL_QueryScan(b *testing.B) {
	main()
	for i := 0; i < b.N; i++ {
		rows, err := sqlDB.Query(`SELECT id, name, logo_url, description_html, is_major FROM technologies`)
		if err != nil {
			b.Error(err)
		}
		defer rows.Close()

		var techs []models.Technology
		for rows.Next() {
			var t models.Technology
			if err := rows.Scan(&t.ID, &t.Name, &t.LogoURL, &t.DescriptionHTML, &t.IsMajor); err != nil {
				b.Error(err)
			}
			techs = append(techs, t)
		}
	}
}
