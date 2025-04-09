package config

import (
	"fmt"
	"time"

	"github.com/rogersovich/go-portofolio-v4/utils"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbConf := Config.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbConf.User,
		dbConf.Password,
		dbConf.Host,
		dbConf.Port,
		dbConf.Name,
	)

	var err error
	maxAttempts := 5

	for attempts := 1; attempts <= maxAttempts; attempts++ {
		var db *gorm.DB
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			message := fmt.Sprintf("❌ Attempt %d: Failed to connect to DB: %v", attempts, err)
			utils.LogError(message, "")
			time.Sleep(2 * time.Second)
			continue
		}

		sqlDB, err := db.DB()
		if err != nil {
			message := fmt.Sprintf("❌ Attempt %d: Failed to get sql.DB: %v", attempts, err)
			utils.LogError(message, "")
			time.Sleep(2 * time.Second)
			continue
		}

		// Ping check
		err = sqlDB.Ping()
		if err != nil {
			message := fmt.Sprintf("❌ Attempt %d: Ping failed: %v", attempts, err)
			utils.LogError(message, "")
			time.Sleep(2 * time.Second)
			continue
		}

		// ✅ Set connection pool settings (PRODUCTION FRIENDLY)
		sqlDB.SetMaxOpenConns(100)                 // max open connections
		sqlDB.SetMaxIdleConns(10)                  // idle connections in pool
		sqlDB.SetConnMaxLifetime(30 * time.Minute) // lifetime before closing

		DB = db
		utils.Log.Info("✅ Database connected with pooling!")
		return
	}

	// log.Fatalf("❌ Could not connect to database after %d attempts", maxAttempts)
	utils.Log.Fatalf("❌ Could not connect to database after %d attempts", maxAttempts)
}
