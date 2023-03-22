package main

import (
	"fmt"
	"math/rand"
	"time"

	applikationssg "github.com/cricton/applikations-sg"
)

func coroutineTest() {

	for i := 0; i < 100; i++ {
		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

		fmt.Println(i)
	}
}
func main() {

	//start HMI-SG
	//start Applikations-SG

	//TODO read this from config file
	sg1 := &applikationssg.Airbacksg{
		ControlUnit: &applikationssg.ControlUnit{Name: "Airback System"},
	}

	//hmi := hmisg.HMI{}

	go coroutineTest()
	go coroutineTest()
	coroutineTest()
	fmt.Println(sg1.ControlUnit.GetName())

}
