package server

import (
	"github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	client *api.Client
}

func NewConsulClient(addr string) (*ConsulClient, error) {
	config := api.DefaultConfig()
	config.Address = addr
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}
	return &ConsulClient{client: client}, nil
}

func (c *ConsulClient) RegisterService(name, id, address string, port int) error {
	serviceDef := &api.AgentServiceRegistration{
		ID:      id,
		Name:    name,
		Address: address,
		Port:    port,
	}
	return c.client.Agent().ServiceRegister(serviceDef)
}
