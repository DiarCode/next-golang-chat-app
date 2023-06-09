package database

import (
	"fmt"

	"github.com/DiarCode/next-golang-chat-app/chat/src/config"
	chatpb "github.com/DiarCode/next-golang-chat-app/chat/src/gen/chat"
	"github.com/DiarCode/next-golang-chat-app/chat/src/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		config.AppConfig.DB_HOST,
		config.AppConfig.DB_USER,
		config.AppConfig.DB_PASSWORD,
		config.AppConfig.DB_NAME,
		config.AppConfig.DB_PORT,
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		utils.LoggerFatalf("Failed to connect to the database! \n", err)
	}

	utils.LoggerInfo("Connected to database!")
	utils.LoggerInfo("Running migrations")

	err = DB.AutoMigrate(
		&chatpb.Room{},
		&chatpb.ChatMessage{},
	)

	if err != nil {
		utils.LoggerInfof("Failed to connect to migrate! \n", err)
	}

	utils.LoggerInfo("Migrations done!")
}
