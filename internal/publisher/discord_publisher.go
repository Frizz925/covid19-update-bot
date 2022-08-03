package publisher

import (
	"github.com/bwmarrin/discordgo"
	"github.com/frizz925/covid19japan-chatbot/internal/config"
)

type DiscordPublisher struct {
	*discordgo.Session

	channelIDs []string
}

func NewDiscordPublisher(cfg *config.Discord) (*DiscordPublisher, error) {
	discord, err := discordgo.New("Bot " + cfg.BotToken)
	if err != nil {
		return nil, err
	}
	return &DiscordPublisher{discord, cfg.ChannelIDs}, nil
}

func (dp *DiscordPublisher) Publish(message string) error {
	for _, cid := range dp.channelIDs {
		_, err := dp.ChannelMessageSend(cid, message)
		if err != nil {
			return err
		}
	}
	return nil
}
