package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"protobuf-lesson/pb"

	"google.golang.org/protobuf/proto"
)

func main() {
	employee := &pb.Employee{
		Id:          1,
		Name:        "Suzuki",
		Email:       "test@example.com",
		Occupation:  pb.Occupation_ENGINEER,
		PhoneNumber: []string{"080-1234-5678", "070-1234-5678"},
		Profile: &pb.Employee_Text{
			Text: "My name is Suzuki",
		},
		Birthday: &pb.Date{
			Year:  2000,
			Month: 1,
			Date:  1,
		},
	}

	binDate, err := proto.Marshal(employee)
	if err != nil {
		log.Fatalln("Can't serialize", err)
	}

	if err := ioutil.WriteFile("test.bin", binDate, 0666); err != nil {
		log.Fatalln("Can't serialize", err)
	}

	in, err := ioutil.ReadFile("test.bin")
	if err != nil {
		log.Fatalln("Can't serialize", err)
	}

	readEmployee := &pb.Employee{}

	err = proto.Unmarshal(in, readEmployee)
	if err != nil {
		log.Fatalln("Can't serialize", err)
	}

	fmt.Println(readEmployee)
}
