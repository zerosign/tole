package base

type Value interface {
	IsArray() bool
	IsMap() bool
	IsPrimitive() bool
}

type SafeConvert interface {
	ToArray() (AbstractArray, error)
	ToMap() (AbstractMap, error)
}

type Query func(Value) (Value, error)

type ParseQuery interface {
	ParseQuery(path string) (Query, error)
}

type UnsafeConvert interface {
	UnsafeArray() AbstractArray
	UnsafeMap() AbstractMap
}

func ToArray(v Value) (AbstractArray, error) {
	if v.IsArray() {
		return v.(AbstractArray), nil
	} else {
		return EmptyAbstractArray(), ArrayConvError(v)
	}
}

func UnsafeArray(v Value) AbstractArray {
	return v.(AbstractArray)
}

func UnsafeMap(v Value) AbstractMap {
	return v.(AbstractMap)
}

func ToMap(v Value) (AbstractMap, error) {
	if v.IsMap() {
		return v.(AbstractMap), nil
	} else {
		return EmptyAbstractMap(), MapConvError(v)
	}
}

type Traversable interface {
	Query(path string) (Value, error)
}

func QueryValue(v Value, path string) (Value, error) {
	if v.IsPrimitive() {
		return nil, PrimitiveQueryError(v)
	} else if v.IsArray() {
		array := UnsafeArray(v)
		query, err := array.ParseQuery(path)
		if err != nil {
			return nil, err
		}

		return query(v)
	} else if v.IsMap() {
		dict := UnsafeMap(v)
		query, err := dict.ParseQuery(path)
		if err != nil {
			return nil, err
		}

		return query(v)
	} else {
		return nil, UnknownTypeError(v)
	}
}
