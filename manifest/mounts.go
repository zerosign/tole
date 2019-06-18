package manifest

import (
	md "github.com/zerosign/tole/metadata"
)

const (
	FIELD_MOUNTS       = "mounts"
	FIELD_MOUNT_SOURCE = "source"
	FIELD_MOUNT_TARGET = "target"
)

type Mount struct {
	source md.LimitedURI
	target md.LimitedURI
}

func (m Mount) Target() md.LimitedURI {
	return m.target
}

func (m Mount) Source() md.LimitedURI {
	return m.source
}

type Mounts []Mount

var (
	emptyMount = Mount{}
)

func intoMount(rawMount map[string]string) (Mount, error) {
	var value string
	var uri md.LimitedURI
	var values []md.LimitedURI = make([]md.LimitedURI, 2)
	var err error
	var ok bool

	if value, ok = rawMount[FIELD_MOUNT_SOURCE]; !ok {
		return emptyMount, FieldNotExists(FIELD_MOUNTS, FIELD_MOUNT_SOURCE)
	}

	uri, err = md.ParseURI(value)

	if err != nil {
		return emptyMount, err
	}

	values[0] = uri

	if value, ok = rawMount[FIELD_MOUNT_TARGET]; !ok {
		return emptyMount, FieldNotExists(FIELD_MOUNTS, FIELD_MOUNT_TARGET)
	}

	uri, err = md.ParseURI(value)

	if err != nil {
		return emptyMount, err
	}

	values[1] = uri

	return Mount{values[0], values[1]}, nil
}

func intoMounts(rawMounts []map[string]string) (Mounts, []error) {
	var mounts Mounts = make(Mounts, 0)
	var errors []error = make([]error, 0)

	for _, rawMount := range rawMounts {
		mount, err := intoMount(rawMount)

		if err == nil {
			mounts = append(mounts, mount)
		} else {
			errors = append(errors, err)
		}
	}

	return mounts, errors
}
