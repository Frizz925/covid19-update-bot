package sources

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/frizz925/covid19-update-bot/internal/config"
)

type awsSecretsSource struct {
	*secretsmanager.Client
	SecretId string
}

func AWSSecretsSource(ctx context.Context, secretId string) (config.Source, error) {
	cfg, err := loadAWSConfig(ctx)
	if err != nil {
		return nil, err
	}
	return &awsSecretsSource{
		Client:   secretsmanager.NewFromConfig(cfg),
		SecretId: secretId,
	}, nil
}

func (s *awsSecretsSource) Load(ctx context.Context) (*config.Config, error) {
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
