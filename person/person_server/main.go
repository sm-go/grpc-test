package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/smith-golang/grpc-test/person/personpb"
	"google.golang.org/protobuf/proto"
)

func main() {
	person := &personpb.Person{
		Firstname: "John",
		Lastname:  "Doe",
	}

	serializedPerson, err := proto.Marshal(person)
	if err != nil {
		log.Fatal("marshalling errr", err)
	}

	ioutil.WriteFile("person.data", serializedPerson, 0644)

	person2 := &personpb.Person{}

	err = proto.Unmarshal(serializedPerson, person2)
	if err != nil {
		log.Fatal("unmarshalling error:", err)
	}
	fmt.Println(person2.GetFirstname())
	fmt.Println(person2.GetLastname())
}
