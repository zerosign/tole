package source

import (
	"github.com/zerosign/tole/base"
	"github.com/zerosign/tole/source"
)

// SourceFactory : Factory for source
//
//
type SourceFactory (func(uri md.LimitedURI, options map[string]base.StringList) (source.Source, error))

type SourceRegistry struct {
	inner map[string]SourceFactory
}

var (
	emptySourceRegistry = SourceRegistry{}
)

func Registry(sources map[string]md.LimitedURI) (SourceRegistry, []error) {
	return emptySourceRegistry, nil
}

// Source : Get source from source registry
//
//
func (f *SourceFactory) Source(name string) (source Source, flag bool) {
	source, flag = f.inner[name]
	return source, flag
}
