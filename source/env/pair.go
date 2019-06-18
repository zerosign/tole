package env

type Pair []string
type Pairs []Pair

func NewPair(key, value string) Pair {
	return []string{key, value}
}

func EmptyPair() Pair {
	return emptyPair
}

func EmptyPairs() Pairs {
	return emptyPairs
}

func IsEmptyPair(pair Pair) bool {
	return len(pair) == 0
}

func IsEmptyPairs(pairs Pairs) bool {
	return len(pairs) == 0
}

var (
	emptyPair  = []string{}
	emptyPairs = []Pair{}
)
