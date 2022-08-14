package aws

import (
	"context"

	"github.com/frizz925/covid19-update-bot/internal/config"
	"github.com/frizz925/covid19-update-bot/internal/config/sources"
	"github.com/frizz925/covid19-update-bot/internal/lambda"
)

type lambdaSource struct {
	config.Source
	event *lambda.Event
}

func LambdaSource(event *lambda.Event) config.Source {
	return &lambdaSource{
		Source: sources.EnvSource(),
		event:  event,
	}
}

func (s *lambdaSource) Load(ctx context.Context) (*config.Config, error) {
	cfg, err := s.Source.Load(ctx)
	if err != nil {
		return nil, err
	}
	if s.event.DataSources != nil {
		cfg.DataSources = s.event.DataSources
	}
	if s.event.Channels != nil {
		cids := make([]string, len(s.event.Channels))
		for idx, ch := range s.event.Channels {
			cids[idx] = ch.ID
		}
		cfg.Discord.ChannelIDs = cids
	}
	if s.event.Storage.Type != "" {
		cfg.Storage.Type = s.event.Storage.Type
	}
	if s.event.Storage.S3Region != "" {
		cfg.Storage.S3Region = s.event.Storage.S3Region
	}
	if s.event.Storage.S3Bucket != "" {
		cfg.Storage.S3Bucket = s.event.Storage.S3Bucket
	}
	return cfg, nil
}
