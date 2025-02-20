package vo

import (
	"math/rand"
	"strconv"
)

type ID struct {
	Value string
}

func NewID() *ID {
	var id string
	for i := 0; i <= 6; i++ {
		id += strconv.Itoa(rand.Intn(10))
	}

	return &ID{
		Value: id,
	}
}