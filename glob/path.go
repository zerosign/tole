package glob

import (
	"os"
)

var (
	emptyGlobPath = GlobPath{}
)

type GlobPath struct {
	path     string
	fileMode os.FileMode
}

func (g GlobPath) Path() string {
	return g.path
}

func (g GlobPath) FileMode() os.FileMode {
	return g.fileMode

}

func EmptyGlobPath() GlobPath {
	return emptyGlobPath
}
