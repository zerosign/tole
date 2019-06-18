package source

import (
	md "github.com/zerosign/tole/metadata"
	"regexp"
	"strings"
)

// Subpath newtype
type Subpath []string

// ServicePath newtype that represents service identifier in path
//
type ServicePath Subpath

// Namespace newtype that represents namespace identifier in the path
//
type Namespace string

// BackendType newtype that represents backend type in the path
// ex: kv or db
type BackendType int

const (
	// Kv : BackendType key value represents any key value storage
	Kv = iota
	// Db : BackendType database represents any database other than key value storage
	Db
)

type Backend struct {
	kind    BackendType
	backend string
}

func BackendOf(str string) Backend {
	if str == "kv" {
		return Backend{Kv, str}
	} else {
		return Backend{Db, str}
	}
}

// LookupPath : Represents the actual path of given source declaration
//
//
type LookupPath struct {
	version   md.Version
	namespace Namespace
	backend   Backend
	service   ServicePath
	env       string
	path      string
}

var (
	// Pattern : the pattern of the path
	//
	// ex:
	// vault+https://cluster01.company.internal/v1.1/company/team10/kv/peroject_01/stg/environment
	// <source>://<host>/<version>/<namespace>/<type>/<service>/<env>/<path>
	// default://
	//
	// <version>/<namespace>/<type>/<service>/<env>/<path>
	//
	// alias+staging://environment
	// alias+<alias-name>://<path>
	//
	Pattern = regexp.MustCompile("^v([\\w\\.]+)\\/(.*)\\/(kv|db)+\\/(.*)\\/(stg|prod|canary|test|local)\\/(.*)")

	emptyLookupPath = LookupPath{}
)

func ParseLookupPath(path string) (LookupPath, error) {
	var version md.Version
	var err error
	var queryPath string = ""

	matches := Pattern.FindAllStringSubmatch(path, -1)

	if len(matches) == 0 {
		// TODO: add new error struct
		// <version>/<namespace>/<type>/<service>/<env>/<path>
		return emptyLookupPath, UnsupportedLookupPath(path)
	}

	size := len(matches[0])

	if size != 5 && size != 6 {
		return emptyLookupPath, LookupPathIncomplete(path)
	}

	if version, err = md.ParseVersion(matches[0][1]); err != nil {
		return emptyLookupPath, err
	}

	if len(matches[0]) == 6 {
		queryPath = matches[0][6]
	}

	return LookupPath{
		version,
		Namespace(matches[0][2]),
		BackendOf(matches[0][3]),
		strings.Split(matches[0][4], "/"),
		matches[0][5],
		queryPath,
	}, nil
}
