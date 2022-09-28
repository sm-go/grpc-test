package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/smith-golang/grpc-test/hello/hellopb"
	"google.golang.org/protobuf/proto"
)

func main() {
	hello := &hellopb.Hello{
		Firstname: "Smith",
		Lastname:  "Golang",
		World: &hellopb.World{
			World: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
		},
		Mgl: &hellopb.Mingalarpar{
			Speech: "Hello Everybody",
			Lang:   "English",
		},
	}

	serializedHello, err := proto.Marshal(hello)
	if err != nil {
		log.Fatal("Marshalling Error", err)
	}
	ioutil.WriteFile("hello/hello.data", serializedHello, 0644)

	helloshow := &hellopb.Hello{}
	err = proto.Unmarshal(serializedHello, helloshow)
	if err != nil {
		log.Fatal("Unmarshalling error", err)
	}
	fmt.Println("Firstname - ", helloshow.GetFirstname())
	fmt.Println("Lastname - ", helloshow.GetLastname())
	fmt.Println("World - ", helloshow.World.GetWorld())
	fmt.Println("Speech - ", helloshow.Mgl.GetSpeech())
	fmt.Println("Lang - ", helloshow.Mgl.GetLang())
	fmt.Println("Mgl Obj - ", helloshow.GetMgl())
	fmt.Println("Wrold Obj - ", helloshow.GetWorld())
}
