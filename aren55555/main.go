package main

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/google/jsonapi"
)

type Foo []int

type Model struct {
	ID int   `jsonapi:"primary,model"`
	S  []int `jsonapi:"attr,s"`
}

func main() {
	p, err := jsonapi.MarshalOne(&Model{
		S: Foo{1, 2, 3},
	},
	)
	if err != nil {
		panic(err)
	}

	b, err := json.MarshalIndent(p, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println("SERIALIZED:")
	fmt.Println(string(b))

	m := &Model{}
	if err := jsonapi.UnmarshalPayload(bytes.NewReader(b), m); err != nil {
		panic(err)
	}
	fmt.Println(m)
}
