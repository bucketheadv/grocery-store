package core

type Pair[T any, R any] struct {
	Left  T `json:"left"`
	Right R `json:"right"`
}
