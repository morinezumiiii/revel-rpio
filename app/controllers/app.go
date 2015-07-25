package controllers

import (
	"github.com/revel/revel"
	"github.com/stianeikeland/go-rpio"
	"fmt"
)

type App struct {
	*revel.Controller
}

func (c App) Index(pinValue bool) revel.Result {
	fmt.Println(pinValue)
	//pinValue = true
	return c.Render(pinValue)
}

func (c App) Gpio17(pinValue bool) revel.Result {

	fmt.Println("GPIO17 pin value is: ", pinValue)
	if err := rpio.Open(); err != nil {
		fmt.Println(err)
	} else {
		defer rpio.Close() // 関数がリターンする直前に実行されるもの
		pin := rpio.Pin(17) // Physical pin# 11
		pin.Output() // Output mode
		//pin.Mode(rpio.Output)

		if pinValue {
			pin.High() // Set pin High
			//pin.Write(rpio.High)
		} else {
			pin.Low() // Set pin Low
			//pin.Write(rpio.Low)
		}	
		//pin.Toggle() // Toggle pin (Low -> High -> Low)
	}
	return c.Redirect(App.Index)
}
