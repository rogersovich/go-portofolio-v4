package seeds

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v4"
	"github.com/rogersovich/go-portofolio-v4/config"
	"github.com/rogersovich/go-portofolio-v4/models"
)

func SeedTechnologies() {
	for i := 0; i < 10000; i++ {
		var tech models.Technology
		if err := faker.FakeData(&tech); err != nil {
			fmt.Printf("❌ Failed to generate technology: %v\n", err)
			continue
		}

		listTechs := []string{"Vue Js", "React", "Golang", "Angular", "Github", "Docker", "Elastic Search", "AWS", "Node Js", "GraphQL"}

		// Create a local rand generator with its own seed
		r := rand.New(rand.NewSource(time.Now().UnixNano()))

		// Pick a random tech
		randomTech := listTechs[r.Intn(len(listTechs))]

		tech.Name = randomTech

		// Optionally control some fields manually
		tech.IsMajor = (i%2 == 0)
		tech.LogoURL = "https://picsum.photos/200"

		if err := config.DB.Create(&tech).Error; err != nil {
			fmt.Printf("❌ Failed to seed technology: %v\n", err)
		} else {
			fmt.Printf("✅ Seeded technology: %s\n", tech.Name)
		}
	}
}
