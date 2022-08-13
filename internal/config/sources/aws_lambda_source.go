package sources

import (
	"context"

	"github.com/frizz925/covid19-update-bot/internal/config"
	"github.com/frizz925/covid19-update-bot/internal/lambda"
)

type awsLambdaSource struct {
	envSource
	event *lambda.Event
}

func AWSLambdaSource(event *lambda.Event) config.Source {
	return &awsLambdaSource{
		event: event,
	}
}

func (s *awsLambdaSource) Load(ctx context.Context) (*config.Config, error) {
	cfg, err := s.envSource.Load(ctx)
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
	return cfg, nil
}
