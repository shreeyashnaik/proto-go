package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/go-faker/faker/v4"
	pb "github.com/shreeyashnaik/proto-go/proto"
	"google.golang.org/protobuf/proto"
)

func doSimple() *pb.Simple {
	return &pb.Simple{
		Id:          2116,
		IsSimple:    true,
		Name:        "Shreeyash",
		SampleLists: []int32{1, 2, 3, 4},
	}
}

func doComplex() *pb.Complex {
	return &pb.Complex{
		SingleDummy: &pb.Dummy{Id: 31, Name: "Shreeyash"},
		MultipleDummies: []*pb.Dummy{
			{Id: 1, Name: "Rital"},
			{Id: 2, Name: "Minal"},
		},
	}
}

func doEnumeration() *pb.Enumeration {
	return &pb.Enumeration{
		EyeColor: pb.EyeColor_EYE_COLOR_BLACK,
	}
}

func doOneof(payload interface{}) {
	switch x := payload.(type) {
	case *pb.Result_Id:
		fmt.Println(payload.(*pb.Result_Id).Id)
	case *pb.Result_Msg:
		fmt.Println(payload.(*pb.Result_Msg).Msg)
	default:
		fmt.Errorf("payload has unexpected type: %v\n", x)
	}

}

func doMaps() *pb.MapIds {
	return &pb.MapIds{
		Ids: map[string]*pb.IdWrapper{
			"id1": &pb.IdWrapper{Id: 20},
			"id2": &pb.IdWrapper{Id: 30},
		},
	}
}

func doFile(p proto.Message) {
	fpath := "simplebin"

	WriteToFile(fpath, p)

	msg := &pb.Simple{}
	ReadFromFile(fpath, p)

	fmt.Println(msg)
}

func main() {
	// fmt.Println(doSimple())
	// fmt.Println(doComplex())
	// fmt.Println(doEnumeration())

	// doOneof(&pb.Result_Id{Id: 1})
	// doOneof(&pb.Result_Msg{Msg: "Shreeyash"})

	// fmt.Println(doMaps())

	// doFile(doSimple())

	doAddressBook()

}

func ReadFile(fpath string, m proto.Message) {
	b, err := ioutil.ReadFile(fpath)
	if err != nil {
		log.Println("Unable to open file")
		return
	}

	if err := proto.Unmarshal(b, m); err != nil {
		log.Println("Unable to unmarshal file")
		return
	}

	fmt.Println(m)
}

func WriteFile(fpath string, m proto.Message) {
	b, err := proto.Marshal(m)
	if err != nil {
		log.Println("Cannot serialize to bytes")
		return
	}

	if err := ioutil.WriteFile(fpath, b, 0644); err != nil {
		log.Println("Unable to write to file")
		return
	}
}

func doAddressBook() {
	abs := &pb.AddressBook{}

	for i := 0; i < 4; i++ {
		abs.People = append(abs.People, &pb.Person{
			Name:  faker.Name(),
			Id:    rand.Int31(),
			Email: faker.Email(),
			Phones: []*pb.Person_PhoneNumber{
				{
					Number: faker.Phonenumber(),
					Type:   1,
				},
			},
		})
	}

	WriteFile("addressbook.bin", abs)

	absNew := &pb.AddressBook{}
	ReadFile("addressbook.bin", absNew)
}
