package handlers

import (
	"github.com/hashicorp/vault/api"
	"os"
)

var vaultToken = os.Getenv("VAULT_TOKEN")
var vaultAddr = os.Getenv("VAULT_ADDR")
var secretPath = os.Getenv("SECRET_PATH")
var accessTokenKey = os.Getenv("ACCESS_TOKEN_KEY")

type Credentials struct {
	AccessToken string
}

func RetrieveMonzoCredentials() (*Credentials, error) {
	config := &api.Config{
		Address: vaultAddr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	client.SetToken(vaultToken)
	c := client.Logical()
	secret, err := c.Read(secretPath)
	if err != nil {
		return nil, err
	}

	return &Credentials{
		AccessToken: secret.Data[accessTokenKey].(string),
	}, nil
}
