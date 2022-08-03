package config

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

type awsSecretsSource struct {
	*secretsmanager.Client
	SecretId string
}

func AWSSecretsSource(ctx context.Context, secretId string) (Source, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	return &awsSecretsSource{
		Client:   secretsmanager.NewFromConfig(cfg),
		SecretId: secretId,
	}, nil
}

func (sm *awsSecretsSource) Load(ctx context.Context) (*Config, error) {
	res, err := sm.GetSecretValue(ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(sm.SecretId),
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
