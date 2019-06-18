package engine

import (
	md "github.com/zerosign/tole/metadata"
)

// Registry : registery (for uri alias) only registry no unregistration
type Registry struct {
	inner map[string]md.LimitedURI
}

func (r *Registry) Register(alias string, base md.LimitedURI) error {
	if _, ok := r.inner[alias]; ok {
		// if uri is same do nothing
		// but if uri is different return an error
	} else {
		r.inner[alias] = base
	}
}

// ExpandURI : expanding URI recursively
//
//
//
func (r *Registry) ExpandURI(uri md.LimitedURI) (md.LimitedURI, error) {

	return nil, nil
}
