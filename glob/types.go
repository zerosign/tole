package glob

type ExpanderFn (func(parent globPath) ([]globPath, []error))

type ResolveFn (func(path string, statFn StatFn) (globPath, error))

type StatFn (func(path string) (os.FileMode, error))

type WalkFn (func(path string, expanderFn ExpanderFn, err error) []error)

type TraverseFn (func(path globPath, walkFn WalkFn) error)

type PathLike interface {
	Expander() ExpanderFn
	Resolver() ResolveFn
	Stat() StatFn
	Traverse() TraverseFn
}

func Resolve(plike PathLike, path string) (GlobPath, error) {
	info, err := plike.Stat()(path)

	if err != nil {
		return EmptyGlobPath(), err
	}

	return GlobPath{path, info}, nil
}

func Expand(plike PathLike, path GlobPath) ([]GlobPath, []error) {
	subpaths, errors := plike.Expander()(path)
	return subpaths, errors
}

func Stat(plike PathLike, path string) (os.FileMode, error) {
	info, err := plike.Stat()(path)
	return info, err
}

func Traverse(plike PathLike, path GlobPath, walkFn WalkFn) []error {
	return nil
}
