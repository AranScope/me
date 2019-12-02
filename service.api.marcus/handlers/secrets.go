package handlers

import (
	"github.com/hashicorp/vault/api"
	"os"
)

var vaultToken = os.Getenv("VAULT_TOKEN")
var vaultAddr = os.Getenv("VAULT_ADDR")
var credentialPath = os.Getenv("CREDENTIAL_SECRET_PATH")
var usernameKey = os.Getenv("USERNAME_KEY")
var passwordKey = os.Getenv("PASSWORD_KEY")

type Credentials struct {
	Username string
	Password string
}

func RetrieveMarcusCredentials() (*Credentials, error) {
	config := &api.Config{
		Address: vaultAddr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	client.SetToken(vaultToken)
	c := client.Logical()
	secret, err := c.Read(credentialPath)
	if err != nil {
		return nil, err
	}

	return &Credentials{
		Username: secret.Data[usernameKey].(string),
		Password: secret.Data[passwordKey].(string),
	}, nil
}
