package publisher

import "github.com/bwmarrin/discordgo"

type DiscordPublisher struct {
	*discordgo.Session

	channelID string
}

func NewDiscordPublisher(token, channelID string) (*DiscordPublisher, error) {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, err
	}
	return &DiscordPublisher{discord, channelID}, nil
}

func (dp *DiscordPublisher) Publish(message string) error {
	_, err := dp.ChannelMessageSend(dp.channelID, message)
	return err
}
