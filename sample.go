package main

import (
	"fmt"
	"ocl"
)


func main() {
	fmt.Printf("go OpenCL example\n")
	fmt.Printf("Platforms in system %d\n", ocl.GetNumberOfPlatforms())
}
