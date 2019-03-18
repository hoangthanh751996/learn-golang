package main

import (
	"fmt"
	"net/http"
)

type car struct {
	gas_pedal      uint16
	brake_pedal    uint16
	steering_wheel int16
	top_speed_kmh  float64
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Whow, hello Go")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Expert web design by Thanh Hoang")
}

func main() {
	a_car := car{
		gas_pedal:      22314,
		brake_pedal:    0,
		steering_wheel: 12516,
		top_speed_kmh:  225.0}
	fmt.Println("car is ", a_car.gas_pedal)
}
