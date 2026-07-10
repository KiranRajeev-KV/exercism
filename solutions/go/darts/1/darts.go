package darts

import "math"

func Score(x, y float64) int {
	distance := math.Sqrt((x-0)*(x-0) + (y-0)*(y-0))
    if distance>10 {
        return 0
    } else if distance<=10 && distance > 5 {
        return 1
    } else if distance <=5 && distance > 1 {
        return 5
    } else if distance <=1 && distance>=0 {
        return 10
    }
    return -1
}