package publisher

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"
	"github.com/frizz925/covid19-update-bot/internal/config"
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
		if err == nil {
			continue
		}
		// HACK: Just log the errors in case we failed to send the message
		log.Printf("Failed to send message to channel ID %s, error: %s", cid, err)
	}
	return nil
}

func (dp *DiscordPublisher) PublishEmbed(embed *Embed) error {
	me := discordgo.MessageEmbed{
		Title:       embed.Title,
		Description: embed.Description,
		URL:         embed.URL,
	}
	if embed.Author.Name != "" {
		me.Author = &discordgo.MessageEmbedAuthor{
			Name: embed.Author.Name,
			URL:  embed.Author.URL,
		}
	}
	if embed.ImageURL != "" {
		me.Image = &discordgo.MessageEmbedImage{
			URL: embed.ImageURL,
		}
	}
	if embed.Footer != "" {
		me.Footer = &discordgo.MessageEmbedFooter{
			Text: embed.Footer,
		}
	}
	if !embed.Timestamp.IsZero() {
		me.Timestamp = embed.Timestamp.Format(time.RFC3339)
	}
	if len(embed.Fields) > 0 {
		me.Fields = make([]*discordgo.MessageEmbedField, len(embed.Fields))
		for idx, field := range embed.Fields {
			me.Fields[idx] = &discordgo.MessageEmbedField{
				Name:  field.Name,
				Value: field.Value,
			}
		}
	}
	for _, cid := range dp.channelIDs {
		_, err := dp.ChannelMessageSendEmbed(cid, &me)
		if err == nil {
			continue
		}
		// HACK: Just log the errors in case we failed to send the message
		log.Printf("Failed to send message to channel ID %s, error: %s", cid, err)
	}
	return nil
}
