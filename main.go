package main

import (
	"fmt"
	"log"
	"math/big"

	"github.com/nnn-gif/curvesimulator/simulator"
)

func main() {

	sim, err := simulator.New()
	if err != nil {
		log.Fatalf("Error initializing CurveSimulation: %v", err)
	}
	defer sim.Close()

	finalop, err := sim.Simulate("0xEeeeeEeeeEeEeeEeEeEeeEEEeeeeEeeeeeeeEEeE", "0xae7ab96520DE3A18E5e111B5EaAb095312D7fE84", big.NewInt(1000000000000000000))
	if err != nil {
		fmt.Println("err", err)
		return

	}

	fmt.Println("finalop", finalop)

}
