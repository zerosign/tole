package metadata

import (
	"github.com/zerosign/tole/base"
	"net/url"
)

var (
	LOCAL_SCHEMES = base.MakeStrSet(
		[]string{"local", "dotenv"},
	)
)

type LimitedURI struct {
	scheme Scheme
	inner  url.URL
}

func (u LimitedURI) Scheme() Scheme {
	return u.scheme
}

func (u LimitedURI) URL() url.URL {
	return u.inner
}

func (u LimitedURI) IsLocal() bool {
	LOCAL_SCHEMES.Contains(u.scheme.Base())
}

func (u LimitedURI) Lists() []LimitedURI {

}

// ParseURI : parse raw uri string into limited form of uri.
//
//
func ParseURI(rawuri string) (LimitedURI, error) {
	uri, err := url.Parse(rawuri)

	if err != nil {
		return LimitedURI{}, err
	}

	scheme, err := ParseScheme(uri.Scheme)

	if err != nil {
		return LimitedURI{}, err
	}

	return LimitedURI{scheme, uri}, nil
}
