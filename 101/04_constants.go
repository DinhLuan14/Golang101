package main

import "fmt"
import "math"

const s string = "constants"
const i int = 1000


func main(){
	fmt.Println("values const s: %s, i: %d",s,i) // %s ->string %d decima int
	fmt.Printf("values const s: %s, i: %d\n",s,i)
	const n  = 50

	const d = 1e4 / n
	fmt.Println(d)

	fmt.Println(2^2) // XOR
	fmt.Println(math.Pow(2,2))
	fmt.Println(math.Pi)
	fmt.Println(math.Exp(2))

}