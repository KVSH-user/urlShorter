package random

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func NewRandomString(aliasLength int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	randAlias := make([]rune, aliasLength)
	for i := range randAlias {
		randAlias[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(randAlias)
}
