package commands

import (
	"fmt"
	"strings"

	"github.com/Logan9312/GOG-Tournament-Bot/requests"
	"github.com/bwmarrin/discordgo"
)

var StartCommand = discordgo.ApplicationCommand{
	Name:        "start",
	Description: "Starts a tournament",
	Options: []*discordgo.ApplicationCommandOption{
		// TODO discuss if this should be number or a name
		{
			Type:        discordgo.ApplicationCommandOptionString,
			Name:        "name",
			Description: "Name of the tournament",
			Required:    true,
		},
	},
}

// Start off day with a reminder ping that Tournament is today. Charge headsets. Check Discord/Challonge names match. Remind them to switch back to Live if they were on Beta 2 hours before tournament start
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
	signUpChannel    = "1197622237168156742"
	deepSkyBlue      = 0x00bfff
)

func ParseSlashCommand(i *discordgo.InteractionCreate) map[string]interface{} {
	var options = make(map[string]interface{})
	for _, option := range i.ApplicationCommandData().Options {
		options[option.Name] = option.Value
	}

	return options
}

func SendReminder(s *discordgo.Session, i *discordgo.InteractionCreate) {
	embed := &discordgo.MessageEmbed{
		Title:       fmt.Sprintf("Tournament #%d Reminder", tournamentNumber),
		Description: "Today is the tournament day! Please make sure to prepare accordingly.",
		Color:       deepSkyBlue,
		Fields: []*discordgo.MessageEmbedField{
			{
				Name:   "Headsets",
				Value:  "Please charge your headsets.",
				Inline: false,
			},
			{
				Name:   "Discord/Challonge Names",
				Value:  "Check that your Discord and Challonge names match.",
				Inline: false,
			},
			{
				Name:   "Switch from Beta",
				Value:  "If you were on Beta, switch back to Live 2 hours before the tournament starts.",
				Inline: false,
			},
			{
				Name:   "Brackets",
				Value:  "Brackets are made based on previous wins. Best players are ranked from 1 onwards.",
				Inline: false,
			},
			{
				Name:   "Join/Leave Last Minute",
				Value:  "For last-minute adds/drops, use the easy join/leave button.",
				Inline: false,
			},
			{
				Name:   "In-Game Instructions",
				Value:  "Before the tournament begins, quit out of the game and relaunch so the 'Comp' button changes to 'Tournaments'.",
				Inline: false,
			},
			{
				Name:   "Reporting Wins",
				Value:  "Winners must report their wins. The winner moves on and brackets are updated accordingly.",
				Inline: false,
			},
			{
				Name:   "Streaming",
				Value:  "From Round 3 or 4, we begin streaming. We'll go live on Twitch - stay tuned in General and Tournament General channels.",
				Inline: false,
			},
		},
		Footer: &discordgo.MessageEmbedFooter{
			Text: "Good luck to all participants!",
		},
	}

	channelID := "" // Set the channel ID where you want to send the reminder
	_, err := s.ChannelMessageSendEmbed(channelID, embed)
	if err != nil {
		fmt.Println("Error sending reminder embed:", err)
	}
}

func StartTournament(s *discordgo.Session, i *discordgo.InteractionCreate) {

	options := ParseSlashCommand(i)
	tournament, err := requests.CreateTournament(options["name"].(string))
	if err != nil {
		fmt.Println("Error creating tournament:", err)
		return
	}

	_, err = s.ChannelMessageSendComplex(signUpChannel, &discordgo.MessageSend{
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
						CustomID: "join:" + tournament.URL,
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

	tournament, err := requests.GetTournament(strings.Split(i.MessageComponentData().CustomID, ":")[1])
	if err != nil {
		fmt.Println("Error getting tournament:", err)
		return
	}

	participant, err := requests.AddParticipant(tournament, i.Member.User.Username, i.Member.User.ID)
	if err != nil {
		fmt.Println("Error adding participant:", err)
		return
	}

	err = s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseUpdateMessage,
		Data: &discordgo.InteractionResponseData{
			Content: participant.Name + " You have successfully joined the tournament!",
		},
	})
	if err != nil {
		fmt.Println("Error sending message:", err)
		return
	}

}
