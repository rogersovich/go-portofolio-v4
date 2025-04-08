package seeds

import (
	"fmt"

	"github.com/bxcodec/faker/v4"
	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/models"
)

func SeedTechnologies() {
	for i := 0; i < 10; i++ {
		var tech models.Technology
		if err := faker.FakeData(&tech); err != nil {
			fmt.Printf("❌ Failed to generate technology: %v\n", err)
			continue
		}

		// Optionally control some fields manually
		tech.IsMajor = (i%2 == 0)

		if err := config.DB.Create(&tech).Error; err != nil {
			fmt.Printf("❌ Failed to seed technology: %v\n", err)
		} else {
			fmt.Printf("✅ Seeded technology: %s\n", tech.Name)
		}
	}
}
