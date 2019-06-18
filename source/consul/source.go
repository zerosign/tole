package consul

import (
	"github.com/hashicorp/consul/api"
)

type Client struct {
	client *api.KV
}
