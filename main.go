package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/Logan9312/GOG-Tournament-Bot/commands"
	"github.com/Logan9312/GOG-Tournament-Bot/requests"
	"github.com/Logan9312/GOG-Tournament-Bot/routers"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

var commandList = []*discordgo.ApplicationCommand{
	&commands.StartCommand,
}

func main() {

	godotenv.Load(".env")

	token := os.Getenv("DISCORD_TOKEN")
	s, err := discordgo.New("Bot " + token)
	if err != nil {
		fmt.Println("Error creating Discord session: ", err)
		return
	}

	requests.API_KEY = os.Getenv("CHALLONGE_KEY")
	requests.USER = os.Getenv("CHALLONGE_USER")

	s.AddHandler(InteractionHandler)

	err = s.Open()
	if err != nil {
		fmt.Println("Failed to open a websocket connection with discord. Likely due to an invalid token. ", err)
		return
	}

	//Builds global commands
	_, err = s.ApplicationCommandBulkOverwrite(s.State.User.ID, "", commandList)
	if err != nil {
		fmt.Println("bulk Overwrite Prod Command Error:", err)
		return
	}

	fmt.Println("Bot is running!")

	routers.HealthCheck()
}

func InteractionHandler(s *discordgo.Session, i *discordgo.InteractionCreate) {
	switch i.Type {
	case discordgo.InteractionApplicationCommand:
		switch i.ApplicationCommandData().Name {
		case "start":
			commands.StartTournament(s, i)
		}

	case discordgo.InteractionMessageComponent:
		switch strings.Split(i.MessageComponentData().CustomID, ":")[0] {
		case "join":
			commands.JoinTournament(s, i)
		}
	}
}
