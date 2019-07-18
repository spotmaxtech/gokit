package gokit

import (
	"github.com/hashicorp/consul/api"
)

type Consul struct {
	address string
	health  *api.Health
	kv      *api.KV
}

func NewConsul(address string) *Consul {
	config := api.DefaultConfig()
	config.Address = address
	client, err := api.NewClient(config)
	if err != nil {
		panic("connect consul failed, err:" + err.Error())
	}

	consul := &Consul{
		address: address,
		health:  client.Health(),
		kv:      client.KV(),
	}
	return consul
}

func (c *Consul) GetKey(key string) ([]byte, error) {
	pair, _, err := c.kv.Get(key, nil)
	if err != nil {
		return nil, err
	}
	if pair == nil {
		return nil, ErrConsulKeyNotExist
	}
	return pair.Value, nil
}

func (c *Consul) GetService(service string) ([]*api.ServiceEntry, error) {
	entry, _, err := c.health.Service(service, "", true, nil)
	return entry, err
}

func (c *Consul) PutKey(key string, value []byte) error {
	pair := &api.KVPair{Key: key, Value: value}
	_, err := c.kv.Put(pair, nil)
	return err
}
