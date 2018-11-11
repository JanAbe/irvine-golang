package main

import "fmt"

func main() {
	var a, v, s, t float64
	fmt.Println("Enter a value for acceleration: ")
	fmt.Scan(&a)

	fmt.Println("Enter a value for initial velocity: ")
	fmt.Scan(&v)

	fmt.Println("Enter a value for initial displacement: ")
	fmt.Scan(&s)

	fmt.Println("Enter a value for time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, v, s)
	fmt.Printf("Displacement travelled after time %.2f: %.2f\n", t, fn(t))
}

// GenDisplaceFn expects acceleration, init. velocity and init. displacement. It returns a func which computes displacement as a function of time.
func GenDisplaceFn(a, v, s float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*(a*(t*t)) + (v * t) + s
	}
}
