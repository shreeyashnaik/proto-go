package main

import (
	"fmt"

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
	fmt.Println(doSimple())
	fmt.Println(doComplex())
	fmt.Println(doEnumeration())

	doOneof(&pb.Result_Id{Id: 1})
	doOneof(&pb.Result_Msg{Msg: "Shreeyash"})

	fmt.Println(doMaps())

	doFile(doSimple())

}
