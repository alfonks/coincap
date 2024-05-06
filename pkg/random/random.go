package random

import (
	"math/rand"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandomString() string {
	var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	b := make([]rune, 36)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}
