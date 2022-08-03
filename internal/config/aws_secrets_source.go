package config

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type awsSecretsSource struct {
	*secretsmanager.Client
	SecretId string
}

func AWSSecretsSource(ctx context.Context, secretId string) (Source, error) {
	cfg, err := loadAWSConfig(ctx)
	if err != nil {
		return nil, err
	}
	return &awsSecretsSource{
		Client:   secretsmanager.NewFromConfig(cfg),
		SecretId: secretId,
	}, nil
}

func (s *awsSecretsSource) Load(ctx context.Context) (*Config, error) {
	res, err := s.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(s.SecretId),
	})
	if err != nil {
		return nil, err
	}
	var cfg Config
	if err := json.Unmarshal(res.SecretBinary, &cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}
