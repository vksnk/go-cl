package ocl
/*
#include <stdlib.h>
#include <CL/cl.h>

*/
import "C"
import "unsafe"
import "fmt"

const (
	CL_CONTEXT_PLATFORM = C.CL_CONTEXT_PLATFORM
)

//constants for clGetPlatformInfo function
const (
	CL_PLATFORM_PROFILE = C.CL_PLATFORM_PROFILE
	CL_PLATFORM_VERSION = C.CL_PLATFORM_VERSION
	CL_PLATFORM_NAME = C.CL_PLATFORM_NAME
	CL_PLATFORM_VENDOR = C.CL_PLATFORM_VENDOR
	CL_PLATFORM_EXTENSIONS = C.CL_PLATFORM_EXTENSIONS
)

type PlatformInfo C.cl_platform_info


const (
	CL_DEVICE_TYPE_DEFAULT = C.CL_DEVICE_TYPE_DEFAULT
	CL_DEVICE_TYPE_CPU = C.CL_DEVICE_TYPE_CPU
	CL_DEVICE_TYPE_GPU = C.CL_DEVICE_TYPE_GPU
	CL_DEVICE_TYPE_ACCELERATOR = C.CL_DEVICE_TYPE_ACCELERATOR
	CL_DEVICE_TYPE_ALL = C.CL_DEVICE_TYPE_ALL
)

type DeviceType C.cl_device_type

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
	var realNum C.cl_uint = 0
	C.clGetPlatformIDs(C.cl_uint(num), (*C.cl_platform_id)(&platforms[0]), &realNum)
	res := make([]Platform, realNum)
	var i C.cl_uint
	for i = 0; i < realNum; i++ {
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

func (pl* Platform) Devices(devType DeviceType, num uint) []Device {
	devices := make([]aDevice, num)
	var realNum C.cl_uint = 0

	error := C.clGetDeviceIDs(pl.item, C.cl_device_type(devType), C.cl_uint(num), (*C.cl_device_id)(&devices[0]), &realNum)
	res := make([]Device, realNum)
	fmt.Printf("xxx %d\n", int(error))
	var i C.cl_uint
	for i = 0; i < realNum; i++ {
		res[i].item = devices[i]
	}
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
	fmt.Printf("Supress.\n")
	C.clFinish(cq.item);
}

type aContext C.cl_context
type Context struct {
	item aContext
}

func CreateContext(devType DeviceType) *Context {
	var ctx Context
	var error C.cl_int
	ctx.item = aContext(C.clCreateContextFromType(nil, C.cl_device_type(devType), nil, nil, &error))
	fmt.Printf("Context error:%d\n", error)
	return &ctx
}

func (*Context) Foo() {
}
