package main

import (
	"log"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func toJson(pb proto.Message) string {
	out, err := protojson.Marshal(pb)
	if err != nil {
		log.Println("cannot marshal to JSON", err)
		return ""
	}

	return string(out)
}

func fromJson(in string, pb proto.Message) {
	if err := protojson.Unmarshal([]byte(in), pb); err != nil {
		log.Println("cannot unmarshal from json to proto", err)
		return
	}
}
