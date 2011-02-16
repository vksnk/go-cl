package ocl
/*
#include <stdlib.h>
#include <CL/cl.h>

*/
import "C"
import "unsafe"

//constants for clGetPlatformInfo function
const (
	CL_PLATFORM_PROFILE = C.CL_PLATFORM_PROFILE
	CL_PLATFORM_VERSION = C.CL_PLATFORM_VERSION
	CL_PLATFORM_NAME = C.CL_PLATFORM_NAME
	CL_PLATFORM_VENDOR = C.CL_PLATFORM_VENDOR
	CL_PLATFORM_EXTENSIONS = C.CL_PLATFORM_EXTENSIONS
)

type PlatformInfo C.cl_platform_info

func PlatformsNumber() uint {
	var numPlatforms C.cl_uint
	C.clGetPlatformIDs(0, nil, &numPlatforms)
	var res uint = uint(numPlatforms)
	return res
}

type aPlatform C.cl_platform_id
type Platform struct {
	item aPlatform
}

type aDevice C.cl_device_id
type Device struct {
	item aDevice
}

type aContext C.cl_context
type Context struct {
	item aContext
}

type aCommandQueue C.cl_command_queue
type CommandQueue struct {
	item aCommandQueue
}

type Buffer C.cl_mem
type Kernel C.cl_kernel
type Event C.cl_event

//type ContextProperty C.cl_context_property

func Platforms(num uint) []Platform {
	if num == 0 { num = PlatformsNumber() }
	platforms := make([]aPlatform, num)
	C.clGetPlatformIDs(C.cl_uint(num), (*C.cl_platform_id)(&platforms[0]), nil)
	res := make([]Platform, num)
	for i := 0; i < len(platforms); i++ {
		res[i].item = platforms[i]
	}

	return res
}


func (pl *Platform) Info(pinfo PlatformInfo) string {
	const bufSize = 4096
	var bufReal C.size_t = 0

	var cStr unsafe.Pointer = unsafe.Pointer(C.malloc(bufSize))
	defer C.free(cStr)

	C.clGetPlatformInfo(pl.item, C.cl_platform_info(pinfo), bufSize, cStr, &bufReal)
	res := C.GoString((*C.char)(cStr))

	return res
}

/*
func CreateQueue(ctx *Context, dev *Device) (*CommandQueue) {
	return nil
}

func CreateContext() (*Context){
	return nil
}
*/


func (cq *CommandQueue) Finish() {
	C.clFinish(cq.item);
}


