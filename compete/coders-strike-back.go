package main

import "fmt"

/**
 * Solution for https://www.codingame.com/ide/puzzle/coders-strike-back
 **/

type Checkpoint struct {
	x int
	y int
	i int // sequence in the lap i.e. 4th out of 5 checkpoints
}

func main() {
	tick := 1
	for {
		// nextCheckpointX: x position of the next check point
		// nextCheckpointY: y position of the next check point
		// nextCheckpointDist: distance to the next checkpoint
		// nextCheckpointAngle: angle between your pod orientation and the direction of the next checkpoint
		var x, y, nextCPX, nextCPY, ncDist, ncAngle int
		fmt.Scan(&x, &y, &nextCPX, &nextCPY, &ncDist, &ncAngle)

		var opponentX, opponentY int
		fmt.Scan(&opponentX, &opponentY)


		/*
		new strategy...

		if isFirstLap() {
			recordCheckPoint(nextCPX, nextCPY)
			doFirstLapStrategy()

		} else {
			if justPassedCheckpoint() {
				turnToNextCheckpointAndStartBoosting()

			} else if closeToNextCheckpoint() {
				turnSlightlyForNextCheckpoint()
			} else {
				boostFullPower()
			}
		}
		*/







		// fmt.Fprintln(os.Stderr, "Debug messages...")
		thrust := 100

		flatAngle := flatAngle(ncAngle)

		if flatAngle > 90 {
			thrust = 0
		} else if flatAngle > 65 {
			thrust = 25
		} else if flatAngle > 40 {
			thrust = 50

		} else if ncDist < 1000 {
			thrust = slowDown(thrust, 20)
		} else if ncDist < 2000 {
			thrust = slowDown(thrust, 35)
		}

		// if ncAngle > 10 {
		//   thrust = thrust - ncAngle / 2
		// } else if ncAngle < -10 {
		//   thrust = thrust + ncAngle / 2
		// }

		// You have to output the target position
		// followed by the power (0 <= thrust <= 100)
		// i.e.: "x y thrust"
		if tick == 0 {
			fmt.Printf("%d %d BOOST\n", nextCPX, nextCPY)
		} else {
			fmt.Printf("%d %d %d\n", nextCPX, nextCPY, thrust)
		}
		tick++
	}
}

func flatAngle(angle int) int {
	if angle >= 0 {
		return angle
	} else {
		return angle * -1
	}
}

func slowDown(thr int, amountToSlow int) int {
	newThrust := thr - amountToSlow
	//fmt.Println(newThrust)
	//return newThrust
	if newThrust < 0 {
		return 0
	}

	return newThrust
}