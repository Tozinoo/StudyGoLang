package main

import "fmt"

type AngleType float64

const (
	None AngleType = iota
	North
	East
	South
	West
)

func GetAngle(angle float64) AngleType {
	switch {
	case angle >= 300, angle >= 0 && angle < 45:
		return North
	case angle >= 45 && angle < 100:
		return East
	case angle >= 100 && angle < 200:
		return South
	case angle >= 200 && angle < 300:
		return West
	default:
		return None
	}
}

func AngleToString(angle AngleType) string {
	switch angle {
	case North:
		return "North"
	case East:
		return "East"
	case South:
		return "South"
	case West:
		return "West"
	default:
		return "None"
	}
}

func main() {
	fmt.Println("방향", AngleToString(GetAngle(0)))
}
