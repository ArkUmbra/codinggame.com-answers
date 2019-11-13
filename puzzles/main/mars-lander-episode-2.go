package main

import "fmt"

/*
 New Plan...

 Briefly: Fly up, across, then down

 Check if there is a terrain piece above current y, which is between current x and landing x
	if so
		fly up until above the Y of that terrain piece
	else if not above landing zone
		fly sideways towards landing x, maintain height, slow down as target gets closer
	else
		descend slowly
*/


type Point struct {
	x int
	y int
}

func main() {
	landingL, landingR := findLandingZone()
	target := Point{landingR.x - landingL.x, landingR.y}
	//fmt.Printf("%+v %+v\n", landingL, landingR)


	//for {
	//	present, height := highTerrainBetweenHereAndLandingZone()
	//	if present {
	//		flyUpToSafeHeight(height)
	//	} else if ()
	//}

	for {
		// hSpeed: the horizontal speed (in m/s), can be negative.
		// vSpeed: the vertical speed (in m/s), can be negative.
		// fuel: the quantity of remaining fuel in liters.
		// rotate: the rotation angle in degrees (-90 to 90).
		// power: the thrust power (0 to 4).
		var X, Y, hSpeed, vSpeed, fuel, rotate, power int
		fmt.Scan(&X, &Y, &hSpeed, &vSpeed, &fuel, &rotate, &power)
		//mt.Printf("%d %d %d %d %d %d %d ", X, Y, hSpeed, vSpeed, fuel, rotate, power)

		var newRotate, newPower int
		// fmt.Println(os.Stderr, "Debug messages...")
		dist := target.x - X


		// try to control speeds which are too high...
		// if hSpeed > 20 {
		//   newRotate, newPower = 90
		// }

		if withinPoints(landingL, landingR, X) {
			// flatten out descent angle
			newRotate, newPower = angSpdForControlDescent(hSpeed, vSpeed)
			// newRotate = balanceHorSp(hSpeed)
			// newPower = calcDescentPower(vSpeed)
		} else {

			// if dist > -50 && dist < 50 {
			//      newRotate = 0
			//      newPower = calcDescentPower(vSpeed)


			// do angle relative to distance (somehow)
			// min angle , max angle (-)30
			cappedDist := cap(dist, 5000)
			// maxDist 1000 / maxAngle 20 = 500

			newRotate = cappedDist / 50
			newRotate = cap(newRotate, 30)
			newPower = calcDescentPower(vSpeed)
		}

		// rotate power. rotate is the desired rotation angle. power is the desired thrust power.
		//fmt.Println("-20 3")
		fmt.Printf("%d %d\n", newRotate, newPower)
	}
}


func highTerrainBetweenHereAndLandingZone() (bool, int) {
	return false, -1
}

func flyUpToSafeHeight(targetHeight int) {

}



func findLandingZone() (Point, Point) {
	// surfaceN: the number of points used to draw the surface of Mars.
	var surfaceN int
	fmt.Scan(&surfaceN)

	// var prevX int
	// var prevY int
	// Find flat pointer
	searching := true

	var landingL Point
	var landingR Point

	for i := 0; i < surfaceN; i++ {
		// landX: X coordinate of a surface point. (0 to 6999)
		// landY: Y coordinate of a surface point. By linking all the points together in a sequential fashion, you form the surface of Mars.
		var landX, landY int
		fmt.Scan(&landX, &landY)

		if !searching {
			// We have already found a landing spot, but we need to ensure we read all the remaining
			// terrain points before trying to fly
			continue;
		}

		if i > 0 && landY - landingL.y == 0 {
			landingR = Point{landX, landY}
			searching = false
		} else {
			landingL = Point{landX, landY}
		}
	}

	return landingL, landingR
}

func cap(dist int, max int) int {
	if dist >= 0 {
		if dist <= max {
			return dist
		} else {
			return max
		}
	} else {
		if dist >= max * -1 {
			return dist
		} else {
			return max * -1
		}
	}
}

func withinPoints(l Point, r Point, x int) bool {
	return l.x < x && r.x > x
}

func angSpdForControlDescent(h int, vSpeed int) (int, int) {
	if h < 10 && h > -10 {
		return 0, calcDescentPower(vSpeed)

	} else if h >= 10 && h < 20 {
		if vSpeed > 35 || vSpeed < -35 {
			return 10, 4
		} else {
			return 30, 4
		}
	} else if h >= 20 {
		if vSpeed > 35 || vSpeed < -35 {
			return 25, 4
		} else {
			return 60, 4
		}

	} else if h < -10 && h > -20 {
		if vSpeed > 35 || vSpeed < -35 {
			return -10, 4
		} else {
			return -30, 4
		}

		// } else if h <= -20 {
	} else {
		if vSpeed > 35 || vSpeed < -35 {
			return -25, 4
		} else {
			return -60, 4
		}
	}
}

func calcDescentPower(vSpeed int) int {
	//fmt.Printf("Speed %d \n", dSpeed)
	// check how fast it is falling
	if vSpeed < -20 {
		//fmt.Println("Try to set max")
		return 4
	} else if vSpeed < -6 {
		return 2
	} else {
		return 0
	}

}