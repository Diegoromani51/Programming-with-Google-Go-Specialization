package main

import ("fmt")

func GenDisplaceFn(acc, initVel, initDis float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*acc*(t*t) + initVel*t + initDis
	}
}

func main() {
	var a, Vo, So float64
	var t float64
	fmt.Println("Please Enter the Acceleration(a):")
	fmt.Scanln(&a)

	fmt.Println("Please Enter the Initial Velocity(Vo):")
	fmt.Scanln(&Vo)

	fmt.Println("Please Enter the Initial Displacement(So):")
	fmt.Scanln(&So)

	fmt.Println("Please Enter the Amount of Time(t):")
	fmt.Scanln(&t)
	fmt.Println("--------------------------------------------------")

	fn := GenDisplaceFn(a, Vo, So)
	fmt.Println("The Displacement After", t, "Seconds:", fn(t))
	fmt.Println("--------------------------------------------------")


}
