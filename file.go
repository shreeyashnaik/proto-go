package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"google.golang.org/protobuf/proto"
)

/*
	1. First, serialize message to bytes.
	2. Second, write serialized bytes to a file.
*/
func WriteToFile(fpath string, pb proto.Message) {
	out, err := proto.Marshal(pb)
	if err != nil {
		log.Println("cannot serialize to bytes,", err)
		return
	}

	if err := ioutil.WriteFile(fpath, out, 0644); err != nil {
		log.Println("cannot write to file,", err)
		return
	}

	fmt.Println("Data has been written to file.")
}

/*
	1. First, read bytes from file.
	2. Second, unmarshal bytes to message struct.
*/
func ReadFromFile(fpath string, pb proto.Message) {
	in, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Println("cannot read bytes from file", err)
		return
	}

	if err := proto.Unmarshal(in, pb); err != nil {
		log.Println("cannot unmarshal bytes to message", err)
		return
	}

	fmt.Println(pb)

}
