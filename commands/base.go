package commands

import "github.com/bwmarrin/discordgo"

var StartCommand = discordgo.ApplicationCommand{

	Name:                     "test",
	Description:              "This is a test command",
	Options:                  []*discordgo.ApplicationCommandOption{},
}
