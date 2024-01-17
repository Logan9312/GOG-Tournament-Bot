package commands

import (
	"fmt"

	"github.com/Logan9312/GOG-Tournament-Bot/requests"
	"github.com/bwmarrin/discordgo"
)

var StartCommand = discordgo.ApplicationCommand{
	Name:        "start",
	Description: "Starts a tournament",
}

// Start off day with a reminder ping that Tournament is today. Charge headsets. Check Discord/Challonge names match. Remind them to switch back to Live if they were on Beta
// Brackets are made based off previous wins. So best players ranked from 1 - ?
// Flurry of last minute adds/drops is always irritating - easy button for them to leave/join?
// Brackets locked
// Threads created with Room number matching Match number, In Game Room number, and @ the players involved in match
// Biggest issue with threads, Discord and Challonge dont match.
// Before Tournament begins, tell them to quit out of game and relaunch so Comp button changes to Tournaments.
// Winners get reported, winner moves on and brackets are updated
// Round 3 or 4 we begin streaming, post that we go Live on Twitch in General and Tournament General
// Calculate top 8. post to Leaderboard
// It doesn’t have to connect to challonge as long as brackets can get made
// And there can be a ranking of players based on wins'

// TODO Track tournament number
var tournamentNumber = 27

var (
	tournamentRoleID = "1197231913560195082"
	deepSkyBlue      = 0x00bfff
)

func StartTournament(s *discordgo.Session, i *discordgo.InteractionCreate) {

	tournament, err := requests.CreateTournament(fmt.Sprintf("Tournament %d", tournamentNumber))
	if err != nil {
		fmt.Println("Error creating tournament:", err)
		return
	}

	_, err = s.ChannelMessageSendComplex(i.ChannelID, &discordgo.MessageSend{
		Content: fmt.Sprintf("<@&%s>", tournamentRoleID),
		Embeds: []*discordgo.MessageEmbed{
			{
				Title:       fmt.Sprintf("Tournament %d!", tournamentNumber),
				Description: "Tournament Information:",
				Color:       deepSkyBlue,
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Join",
						Style:    discordgo.SuccessButton,
						CustomID: "join",
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

	url := "https://challonge.com/" + tournament.URL

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Embeds: []*discordgo.MessageEmbed{
				{
					Title:       fmt.Sprintf("Tournament %d!", tournamentNumber),
					Description: "Tournament has been successfully created!",
					Color:       deepSkyBlue,
					Fields: []*discordgo.MessageEmbedField{
						{
							Name:   "URL",
							Value:  url,
							Inline: true,
						},
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}
}

func JoinTournament(s *discordgo.Session, i *discordgo.InteractionCreate) {

}
