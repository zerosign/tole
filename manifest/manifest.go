package manifest

import (
	"errors"
	"fmt"

	md "github.com/zerosign/tole/metadata"
	yaml "gopkg.in/yaml.v3"
	"os"
)

const (
	Kind = "ConfigManifest"
)

// Compatibility : check whether current manifest are compatible for this build
//
// for now any manifest with 0.1.x will be compatible with this binary
//
// return true if current version is compatible
//        false if not
//
func Compatibility(version md.Version) bool {
	return version.Major() == 0 && version.Minor() == 1
}

func CompatibleError(version md.Version) error {
	return errors.New(fmt.Sprintf("manifest version %s are not compatible to this binary", version.Format()))
}

// rawManifest : intermediate manifest for unmarshalling from yaml only purpose
//
type rawManifest struct {
	Kind    string              `yaml:"kind"`
	Version string              `yaml:"version"`
	Sources map[string]string   `yaml:"sources"`
	Aliases map[string]string   `yaml:"aliases"`
	Mounts  []map[string]string `yaml:"mounts"`
}

type Manifest struct {
	version md.Version
	sources map[string]md.LimitedURI
	aliases map[string]md.LimitedURI
	mounts  Mounts
}

func (m Manifest) Version() md.Version {
	return m.version
}

func (m Manifest) LocalSources() map[string]md.LimitedURI {
	var sources map[string]md.LimitedURI = make(map[string]md.LimitedURI)

	for name, source := range m.sources {
		if source.IsLocal() {
			sources[name] = source
		}
	}

	return sources
}

func (m Manifest) Sources() map[string]md.LimitedURI {
	return m.sources
}

func (m Manifest) Aliases() map[string]md.LimitedURI {
	return m.aliases
}

func (m Manifest) Mounts() Mounts {
	return m.mounts
}

var (
	emptyRawManifest = rawManifest{}
	emptyManifest    = Manifest{}
)

func intoMapOfURI(rawMap map[string]string) (map[string]md.LimitedURI, []error) {
	var uriMap map[string]md.LimitedURI = make(map[string]md.LimitedURI)
	var errors []error = make([]error, 0)

	for key, value := range rawMap {
		uri, err := md.ParseURI(value)

		if err != nil {
			errors = append(errors, err)
		} else if oldValue, ok := uriMap[key]; ok {
			errors = append(errors, DuplicatedField(key, oldValue, value))
		} else {
			uriMap[key] = uri
		}
	}

	return uriMap, errors
}

func intoManifest(rawManifest rawManifest) (Manifest, []error) {
	var errors []error = make([]error, 0)
	var newErrors []error
	var sources, aliases map[string]md.LimitedURI
	var mounts Mounts

	if rawManifest.Kind != Kind {
		return emptyManifest, []error{UnsupportedManifest(rawManifest.Kind)}
	}

	version, err := md.ParseVersion(rawManifest.Version)

	if err != nil {
		return emptyManifest, []error{err}
	}

	sources, newErrors = intoMapOfURI(rawManifest.Sources)

	if len(newErrors) != 0 {
		errors = append(errors, newErrors...)
	}

	aliases, newErrors = intoMapOfURI(rawManifest.Aliases)

	if len(newErrors) != 0 {
		errors = append(errors, newErrors...)
	}

	mounts, newErrors = intoMounts(rawManifest.Mounts)

	if len(newErrors) != 0 {
		errors = append(errors, newErrors...)
	}

	return Manifest{
		version,
		sources,
		aliases,
		mounts,
	}, errors
}

func ParseManifest(path string) (Manifest, []error) {
	var rawManifest rawManifest = rawManifest{}
	var err error
	var desc *os.File

	desc, err = os.Open(path)

	if err != nil {
		return emptyManifest, []error{err}
	}

	decoder := yaml.NewDecoder(desc)
	decoder.KnownFields(true)

	if err = decoder.Decode(&rawManifest); err != nil {
		return emptyManifest, []error{err}
	}

	if manifest, errors := intoManifest(rawManifest); len(errors) != 0 {
		return emptyManifest, errors
	} else {

		if !Compatibility(manifest.Version()) {
			return emptyManifest, []error{CompatibleError(manifest.Version())}
		}

		return manifest, []error{}
	}
}
