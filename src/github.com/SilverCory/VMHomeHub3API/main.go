package main

import (
	"flag"
	"fmt"
	"github.com/SilverCory/VMHomeHub3API/vmapi"
)

func main() {
	passwordPtr := flag.String("password", "ChangeMe", "The password to log in to the router.")
	flag.Parse()

	instance, err := vmapi.New(*passwordPtr)
	if err != nil {
		panic(err)
	}

	defer instance.Close()

	fmt.Printf("Instance: %#v\n", *instance)

}
