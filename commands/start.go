package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"gorm.io/gorm"
)

type BotContext struct {
	database *gorm.DB
}

var StartCommand = discordgo.ApplicationCommand{
	Name:        "start",
	Description: "Starts a tournament",
}

// Start off day with a reminder ping that Tournament is today. Charge headsets. Check Discord/Challonge names match. Remind them to switch back to Live if they were on Beta
// Brackets are made based off previous wins. So best players ranked from 1 - ?
// flurry of last minute adds/drops is always irritating - easy button for them to leave/join?
// Brackets locked
// Threads created with Room number matching Match number, In Game Room number, and @ the players involved in match
// Biggest issue with threads, Discord and Challonge dont match.
// Before Tournament begins, tell them to quit out of game and relaunch so Comp button changes to Tournaments.
// winners get reported, winner moves on and brackets are updated
// Round X we begin streaming, post that we go Live on Twitch in General and Tournament General
// Calculate top 8. post to Leaderboard

// TODO Track tournament number
var tournamentNumber = 25

var (
	tournamentRoleID = "1197231913560195082"
	deepSkyBlue      = 0x00bfff
)

func StartTournament(s *discordgo.Session, i *discordgo.InteractionCreate) {
	_, err := s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Content: fmt.Sprintf("<@&%s>", tournamentRoleID),
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       fmt.Sprintf("Tournament %d!", tournamentNumber),
				Description: "Tournament Information:",
				Color:       deepSkyBlue,
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
			Content: "Success",
		},
	})
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
}
