package main

import (
	"fmt"
	"math/rand"
)

type enemyAttacker interface {
	fireWeapons()
	move()
	assignDriver(driverName string)
}

func run(a enemyAttacker, driverName string) {
	a.assignDriver(driverName)
	a.fireWeapons()
	a.move()
}

type enemyTank struct {
	damage    int
	moveSpace int
}

func NewEnemyTank() *enemyTank {
	fmt.Println("A new enemy tank has been created")
	return &enemyTank{
		damage:    rand.Intn(20) + 1,
		moveSpace: rand.Intn(10) + 1,
	}
}

func (t enemyTank) fireWeapons() {
	fmt.Println("Cause", t.damage, "damages")
}

func (t enemyTank) move() {
	fmt.Println("Move", t.moveSpace, "spaces")
}

func (t enemyTank) assignDriver(driverName string) {
	fmt.Println("Assign to driver", driverName)
}

type enemyRobot struct {
	smashDamage  int
	walkDistance int
}

// A robot causes less damage and moves slower than a tank
func NewEnemyRobot() *enemyRobot {
	fmt.Println("A new enemy robot has been created")
	return &enemyRobot{
		smashDamage:  rand.Intn(10) + 1,
		walkDistance: rand.Intn(5) + 1,
	}
}

func (r enemyRobot) smashWithHands() {
	fmt.Println("Smash cause", r.smashDamage, "damages")
}

func (r enemyRobot) walk() {
	fmt.Println("Walk in", r.walkDistance, "spaces")
}

func (r enemyRobot) marchWithDriver(driverName string) {
	fmt.Println("Stop tramping, marching with", driverName)
}

type robotAdapter struct {
	robot *enemyRobot
}

func NewRobotAdapter(r *enemyRobot) *robotAdapter {
	fmt.Println("A new enemy robot adapter has been created")
	return &robotAdapter{
		robot: r,
	}
}

func (a robotAdapter) fireWeapons() {
	a.robot.smashWithHands()
}

func (a robotAdapter) move() {
	a.robot.walk()
}

func (a robotAdapter) assignDriver(driverName string) {
	a.robot.marchWithDriver(driverName)
}

func main() {
	// enemy tank
	tank := NewEnemyTank()
	run(tank, "Nathan")

	// enemy robot
	robot := NewEnemyRobot()
	robot.marchWithDriver("Graham")
	robot.smashWithHands()
	robot.walk()
	// the following commented out lines will have compile errors because of incompatibility
	// run(robot, "Graham") // DO NOT do this
	// However, the following lines will pass and work just fine
	// enemy adapter
	adapter := NewRobotAdapter(robot)
	run(adapter, "Donald")
}
