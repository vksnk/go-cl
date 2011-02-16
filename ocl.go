package ocl
/*
#include <stdlib.h>
#include <CL/cl.h>

*/
import "C"


func PlatformsNumber() uint {
	var numPlatforms C.cl_uint
	C.clGetPlatformIDs(0, nil, &numPlatforms)
	var res uint = uint(numPlatforms)
	return res
}

type Platform C.cl_platform_id
type Device C.cl_device_id
type Context C.cl_context
type CommandQueue C.cl_command_queue
type Buffer C.cl_mem
type Kernel C.cl_kernel
type Event C.cl_event

func Platforms(num uint) []Platform {
	res := make([]Platform, num)
	C.clGetPlatformIDs(C.cl_uint(num), (*C.cl_platform_id)(&res[0]), nil)
	return res
}
/*
func CreateQueue(ctx *Context, dev *Device) (*CommandQueue) {
	return nil
}

func CreateContext() (*Context){
	return nil
}

func (cq CommandQueue) Finish() {
	C.clFinish(cq);
}
*/
