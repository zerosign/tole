package secret

import (
	v "github.com/hashicorp/vault/api"
	_ "github.com/zerosign/tole/source"
	"log"
)

//
// Client will only holds lifecycles for this time
// and let the underlying secret lifecycle handle
// each lease renewal or credential renewal.
//
type Source struct {
	lifecycles map[string]SecretLifecycle
	client     *v.Client
}

func (s *Source) SLookup(path string) (string, error) {
	s.client.Logical().Read(path)
}

func (s *Source) Close() {
	log.Printf("closing secret vault source")
	s.Close()
}
