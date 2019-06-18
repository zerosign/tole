package base

var (
	emptyAbstractMap   = make(AbstractMap)
	emptyAbstractArray = make(AbstractArray, 0)
)

type AbstractMap map[string]interface{}

func (instance AbstractMap) IsArray() bool {
	return false
}

func (instance AbstractMap) IsMap() bool {
	return true
}

func (instance AbstractMap) IsPrimitive() bool {
	return false
}

func (instance AbstractMap) ParseQuery(path string) (Query, error) {
	return nil, nil
}

type AbstractArray []interface{}

func (instance AbstractArray) ParseQuery(path string) (Query, error) {
	return nil, nil
}

func (instance AbstractArray) IsArray() bool {
	return true
}

func (instance AbstractArray) IsMap() bool {
	return false
}

func (instance AbstractArray) IsPrimitive() bool {
	return false
}

func EmptyAbstractArray() AbstractArray {
	return emptyAbstractArray
}

func EmptyAbstractMap() AbstractMap {
	return emptyAbstractMap
}
