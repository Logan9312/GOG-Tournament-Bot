package main

import (
	"fmt"
	"os"

	"github.com/Logan9312/GOG-Tournament-Bot/routers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file:", err)
	}

	token := os.Getenv("DISCORD_TOKEN")
	discordgo.New("Bot " + token)

	fmt.Println("Bot is running!")

	routers.HealthCheck()
}
