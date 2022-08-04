package config

import (
	"context"

	"github.com/frizz925/covid19-update-bot/internal/data"
)

type awsLambdaSource struct {
	envSource
	event *data.LambdaEvent
}

func AWSLambdaSource(event *data.LambdaEvent) Source {
	return &awsLambdaSource{
		event: event,
	}
}

func (s *awsLambdaSource) Load(ctx context.Context) (*Config, error) {
	cfg, err := s.envSource.Load(ctx)
	if err != nil {
		return nil, err
	}
	if s.event.CountryID != "" {
		cfg.CountryID = s.event.CountryID
	}
	if s.event.Channels != nil {
		cids := make([]string, len(s.event.Channels))
		for idx, ch := range s.event.Channels {
			cids[idx] = ch.ID
		}
		cfg.Discord.ChannelIDs = cids
	}
	return cfg, nil
}
