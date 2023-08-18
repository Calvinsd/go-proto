package main

import (
	"encoding/json"
	"fmt"
	"log"

	// "os"

	"github.com/Calvinsd/go-proto/model"
	"google.golang.org/protobuf/proto"
)

type PersonJson struct {
	Name string `json:"Name"`
	Age  int32  `json:"Age"`
}

func main() {
	jsonPerson := PersonJson{
		"Newton",
		150,
	}

	protoPerson := model.Person{
		Name: "Newton",
		Age:  150,
	}

	serializedJson, err := json.Marshal(jsonPerson)
	serializedProto, err := proto.Marshal(&protoPerson)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Serialized json length", len(serializedJson))
	fmt.Println("Serialized proto length", len(serializedProto))

	for _, val := range serializedProto {
		fmt.Print(val, " ")
	}

	// protoFile, err := os.OpenFile("./serializedProtoData", os.O_CREATE|os.O_RDWR, 0666)
	// jsonFile, err := os.OpenFile("./serializedJsonData.json", os.O_CREATE|os.O_RDWR, 0666)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// jsonFile.Write(serializedJson)
	// protoFile.Write(serializedProto)
}
