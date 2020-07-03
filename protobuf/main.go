package main

import (
	"fmt"
	"github.com/gotechiecode/goprotobuff/protobuf/model"
	"github.com/golang/protobuf/proto"
)

func main()  {
	fmt.Println("Hello GoTechie...!")

	cust := &model.Customer{
		Name:"John",
		Age:29,
		MastEmployee:true,
	}

	//Marshalling data into protobuff
	data, err := proto.Marshal(cust)
	if err != nil {
		fmt.Println("Marshalling error: ", err)
	}

	newCust := &model.Customer{}

	err = proto.Unmarshal(data, newCust)
	if err !=nil {
		fmt.Println("UnMarshalling error: ", err)
	}

	fmt.Println("UnMarshalled Data: ", newCust)

}