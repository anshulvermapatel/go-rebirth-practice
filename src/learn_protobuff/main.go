package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/golang/protobuf/proto"
)

func main() {
	anshul := &Person{
		Name: "Anshul Verma",
		Age:  24,
		Socialfollowers: &SocialFollowers{
			Youtube: 400,
			Twitter: 500,
		},
	}

	data, err := proto.Marshal(anshul)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data))
	var b bytes.Buffer
	b.Write(data)
	b.WriteTo(os.Stdout)
	//fmt.Println(string(data))

	newAnshul := &Person{}
	err = proto.Unmarshal(data, newAnshul)
	if err != nil {
		panic(err)
	}
	fmt.Println(newAnshul.String())

}
