package aws

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/frizz925/covid19-update-bot/internal/config"
)

type secretsSource struct {
	*secretsmanager.Client
	SecretId string
}

func SecretsSource(cfg aws.Config, secretId string) config.Source {
	return &secretsSource{
		Client:   secretsmanager.NewFromConfig(cfg),
		SecretId: secretId,
	}
}

func (s *secretsSource) Load(ctx context.Context) (*config.Config, error) {
	res, err := s.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(s.SecretId),
	})
	if err != nil {
		return nil, err
	}
	var cfg config.Config
	if err := json.Unmarshal(res.SecretBinary, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
