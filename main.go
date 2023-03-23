package main

import (
	applikationssg "github.com/cricton/applikations-sg"
	commmiddleware "github.com/cricton/comm-middleware"
	hmisg "github.com/cricton/hmi-sg"
)

func main() {

	//start HMI-SG
	//start Applikations-SG

	//TODO read this from config file?
	sg1 := &applikationssg.Airbacksg{
		ControlUnit: &applikationssg.ControlUnit{},
	}

	sg2 := &applikationssg.Assistantsg{
		ControlUnit: &applikationssg.ControlUnit{},
	}

	sg3 := &applikationssg.Navigationsg{
		ControlUnit: &applikationssg.ControlUnit{},
	}

	sg4 := &applikationssg.Infosg{
		ControlUnit: &applikationssg.ControlUnit{},
	}

	middleware := &commmiddleware.Middleware{}
	hmi := hmisg.HMI{}

	//create channel for sg-middleware interaction
	sg1.ControlUnit.RegisterClient(middleware)
	sg2.ControlUnit.RegisterClient(middleware)
	sg3.ControlUnit.RegisterClient(middleware)
	sg4.ControlUnit.RegisterClient(middleware)

	hmi.RegisterHMI(middleware)

	go middleware.Mainloop()

	go sg1.Mainloop()
	// go sg2.Mainloop()
	// go sg3.Mainloop()
	// go sg4.Mainloop()

	hmi.HMI_main_loop()

}
