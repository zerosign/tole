package source

import (
	"regexp"

	"github.com/zerosign/tole/base"
)

type SourceType int

const (
	Env = iota
	Metadata
	Store
)

type Sources map[string]Source

func (ss Sources) Close() {
	for _, source := range ss {
		source.Close()
	}
}

type SourceIdentifier interface {
	Pattern() *regexp.Regexp
	Path(path string) (string, error)
}

type Source interface {
	SLookup(path string) (string, error)
	ALookup(path string) (base.AbstractArray, error)
	HLookup(path string) (base.AbstractMap, error)
	Close()
}

// Static source
type StaticSource map[string]interface{}

type DynamicSource interface {
	Source
	Replace(source *Source)
}
