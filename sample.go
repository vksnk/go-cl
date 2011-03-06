package main

import (
	"fmt"
	"ocl"
)


func main() {
	kernelSrc := "__kernel void hello(void) {}"

	fmt.Printf("go OpenCL example\n")
	platformsNum := ocl.PlatformsNumber()
	fmt.Printf("Platforms in system %d\n", platformsNum)
	platforms := ocl.Platforms(platformsNum)
	for i := 0; i < len(platforms); i++ {
		fmt.Printf("%d :: %s\n", i, platforms[i].Info(ocl.CL_PLATFORM_PROFILE))
		fmt.Printf("  :: %s\n", platforms[i].Info(ocl.CL_PLATFORM_VERSION))
		fmt.Printf("  :: %s\n", platforms[i].Info(ocl.CL_PLATFORM_NAME))
		fmt.Printf("  :: %s\n", platforms[i].Info(ocl.CL_PLATFORM_VENDOR))
		fmt.Printf("  :: %s\n", platforms[i].Info(ocl.CL_PLATFORM_EXTENSIONS))
	}

	ctx := ocl.CreateContext(ocl.CL_DEVICE_TYPE_GPU)
	program := ocl.CreateProgram(ctx, []string{kernelSrc})
	helloKernel := ocl.CreateKernel(program, "hello")
	helloKernel.Foo()
}
