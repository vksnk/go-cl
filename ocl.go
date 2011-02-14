package ocl
/*
#include <CL/cl.h>

cl_platform_id* cl_platform_null = 0;
*/
import "C"

func GetPlatformsNumber() uint {
	var numPlatforms C.cl_uint
	C.clGetPlatformIDs(0, C.cl_platform_null, &numPlatforms)
	var res uint = uint(numPlatforms)
	return res
}

type Context struct {
	AContext C.cl_context
}

type CommandQueue struct {
	AQueue C.cl_command_queue
}

func (cq *CommandQueue) Finish() {
	C.clFinish(cq.AQueue);
}

type Buffer struct {
	ABuffer C.cl_mem
}

type Kernel struct {
	AKernel C.cl_kernel
}


