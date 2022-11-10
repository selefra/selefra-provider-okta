package okta_client

import (
	"context"
	"fmt"
	"github.com/okta/okta-sdk-golang/v2/okta"
	"github.com/pkg/errors"
)

type Client struct {
	Okta   *okta.Client
	Domain string
}

func (c *Client) ID() string {
	return c.Domain
}

func NewClients(configs OktaProviderConfigs) ([]*Client, error) {
	var clients []*Client
	for _, provider := range configs.Providers {
		client, err := newClient(provider)
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}
	return clients, nil
}

func newClient(config OktaProviderConfig) (*Client, error) {
	if config.Domain == "" {
		return nil, errors.New("The configuration OKTA domain name is missing")
	}

	if len(config.Token) == 0 {
		return nil, errors.New("The configuration OKTA token name is missing")
	}

	_, c, err := okta.NewClient(
		context.Background(),
		okta.WithOrgUrl(config.Domain),
		okta.WithToken(config.Token),
		okta.WithCache(true),
	)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	return &Client{
		Okta:   c,
		Domain: config.Domain,
	}, nil
}
