package sn

import (
	"github.com/bwmarrin/snowflake"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

var (
	un = uuid.New()
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func r(n int) chan int {
	out := make(chan int)

	go func(x int) {
		for {
			out <- rand.Intn(x)
		}
	}(n)

	return out
}

var letterRunes = []rune("ABCDEFGHJKLMNPQRSTUVWXYZ23456789")

func RandomString(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func Random(min int, max int) int {
	o := r(max - min)
	return min + <-o
}

func SnowflakeId() (*snowflake.ID, error) {
	var node *snowflake.Node
	node, err := snowflake.NewNode(1)
	if err != nil {
		return nil, err
	}

	id := node.Generate()

	return &id, nil
}

func UuidString() string {
	return un.String()
}
