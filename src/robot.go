package main

import (
	"fmt"
)

type Electronic interface {
	drainPower()
	killRobot()
}

type Robot struct {
	Name  string
	Power int
}

/* Drains the robots power by 1 */
func (r *Robot) drainPower() {
	if r.Power != 0 {
		r.Power--
	} else {
		fmt.Println("This poor robot! Can't you see you've already depleted its battery!?")
		return
	}

	if r.Power == 0 {
		fmt.Printf("WARNING!\n%v's power level has depleted!\n", r.Name)
	} else {
		fmt.Printf("%v's power level is %v.\n", r.Name, r.Power)
	}
}

/* Kills the robot by depleting it's power to 0. */
func (r *Robot) killRobot() {
	currPower := r.Power
	for i := 0; i < currPower; i++ {
		r.drainPower()
	}
}
