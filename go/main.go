package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	data, err := ioutil.ReadFile("../things/pac4200.jsonld")

	if err != nil {
		fmt.Println(err)
		return
	}
	td, err := UnmarshalTdParser(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	for name, prop := range td.Properties {
		fmt.Println(name)
		fmt.Println("Forms:")
		for _, form := range prop.Forms {
			fmt.Println("  " + *form.ModbusEntity)
		}
	}
	//	serialized, _ := td.Marshal()
	//	fmt.Println(string(serialized))
}
