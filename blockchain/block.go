package blockchain

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{[]byte{}, []byte(data), prevHash, 0}
	pow := NewProof(block)

	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis-12089.c291.ap-southeast-2-1.ec2.cloud.redislabs.com:12089",
		Password: "fEK7MQ0uuq5mUT2utYvHk1gnMqPlUCGP",
	})

	bs, err := json.Marshal(block)
	Handle(err)

	err = rdb.Set(ctx, string(hash), string(bs), 0).Err()
	Handle(err)

	return block
}

func Genesis() *Block {
	return CreateBlock("Genesis", []byte{})
}

func Handle(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
}
