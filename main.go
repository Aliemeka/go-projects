package main

import (
	"fmt"
)

type petrolEngine struct {
	kpl    uint16
	litres uint16
}

func (e petrolEngine) distanceLeft() uint16 {
	return e.kpl * e.litres
}

type electricEngine struct {
	kpkwh uint16
	kwh   uint16
}

func (e electricEngine) distanceLeft() uint16 {
	return e.kpkwh * e.kwh
}

type engine interface {
	distanceLeft() uint16
}

func canIReachDestination(e engine, distance uint16) bool {
	return e.distanceLeft() >= distance
}

func main() {
	var myEngine petrolEngine
	myEngine.kpl = 100
	myEngine.litres = 10
	fmt.Println("My engine capacity is", myEngine.kpl, "kpl and has", myEngine.litres, "litres of fuel")
	fmt.Println("I can travel", myEngine.distanceLeft(), "km with the fuel I have")

	var myElectricEngine electricEngine = electricEngine{kpkwh: 100, kwh: 10}
	fmt.Println("I can travel", myElectricEngine.distanceLeft(), "km with the battery I have")

	distance := int16(1000)
	var canReachWithPetrol bool = canIReachDestination(myEngine, uint16(distance))
	fmt.Println("I", map[bool]string{true: "can", false: "cannot"}[canReachWithPetrol], "reach my destination with the petrol engine")
}
