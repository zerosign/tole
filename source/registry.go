package source

import (
	md "github.com/zerosign/tole/metadata"
)

type SourceRegistry struct {
	// source maps
	inner Sources
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
func (r SourceRegistry) Source(name string) (Source, bool) {
	// var factory SourceFactory
	// var source Source
	var flag bool = false

	// resolve source from alias
	// if doesn't exist return false

	// check alias in source maps
	// if factory

	return nil, flag
}

func (r SourceRegistry) IsEmpty() bool {
	return len(r.inner) == 0
}

func (r SourceRegistry) Size() int {
	return len(r.inner)
}
