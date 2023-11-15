package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

var StartCommand = discordgo.ApplicationCommand{
	Name:        "start",
	Description: "Starts a tournament",
	Options:     []*discordgo.ApplicationCommandOption{},
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
				Timestamp:   "",
				Color:       0x00bfff,
				Footer:      &discordgo.MessageEmbedFooter{},
				Image:       &discordgo.MessageEmbedImage{},
				Thumbnail:   &discordgo.MessageEmbedThumbnail{},
				Video:       &discordgo.MessageEmbedVideo{},
				Provider:    &discordgo.MessageEmbedProvider{},
				Author:      &discordgo.MessageEmbedAuthor{},
				Fields:      []*discordgo.MessageEmbedField{},
			},
		},
		TTS:             false,
		Components:      []discordgo.MessageComponent{},
		Files:           []*discordgo.File{},
		AllowedMentions: &discordgo.MessageAllowedMentions{},
		Reference:       &discordgo.MessageReference{},
		StickerIDs:      []string{},
		Flags:           0,
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
			Components:      []discordgo.MessageComponent{},
			Embeds:          []*discordgo.MessageEmbed{},
			AllowedMentions: &discordgo.MessageAllowedMentions{},
			Files:           []*discordgo.File{},
			Flags:           0,
			Choices:         []*discordgo.ApplicationCommandOptionChoice{},
			CustomID:        "",
			Title:           "",
		},
	})
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
}
