package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var StartCommand = discordgo.ApplicationCommand{
	Name:        "start",
	Description: "Starts a tournament",
}

// TODO Track tournament number
var tournamentNumber = 0

func StartTournament(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Content: "<@949044843949203576>",
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       fmt.Sprintf("Tournament %d!", tournamentNumber),
				Description: "Tournament Information:",
				Color:       0x00bfff,
			},
		},
	})
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			TTS:             false,
			Content:         "Success",
		},
	})
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
}
