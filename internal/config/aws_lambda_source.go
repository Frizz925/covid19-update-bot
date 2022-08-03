package config

import (
	"context"

	"github.com/frizz925/covid19japan-chatbot/internal/data"
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
	if s.event.ChannelIDs != nil {
		cfg.Discord.ChannelIDs = s.event.ChannelIDs
	}
	return cfg, nil
}
