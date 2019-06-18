package source

import (
	"github.com/zerosign/tole/base"
	md "github.com/zerosign/tole/metadata"

	"sync"
)

type Credentials map[string]base.StringList

type OptionArgs map[string]interface{}

// SourceFactory : Factory for source
//
//
type SourceFactory (func(uri md.LimitedURI, credential Credentials, options OptionArgs) (Source, error))

var (
	// source factory maps
	once    sync.Once
	factory map[string]SourceFactory
)

func RegisterSource(name string, factory SourceFactory) {
	// TODO: how to register source factory
}

func init() {
	once.Do(func() {
		// initialize factory
		factory = make(map[string]SourceFactory)
	})
}
