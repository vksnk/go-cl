package ocl
/*
#include <stdlib.h>
#include <CL/cl.h>

static char**makeCharArray(int size) {
        return calloc(sizeof(char*), size); 
} 

static void setArrayString(char **a, char *s, int n) { 
        a[n] = s; 
} 

static void freeCharArray(char **a, int size) { 
        int i; 
        for (i = 0; i < size; i++) 
                free(a[i]); 
        free(a); 
}

static cl_context_properties platform2property(cl_platform_id pid) {
	return (cl_context_properties)pid;
}
*/
import "C"
import "unsafe"
import "fmt"
import "runtime"

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

func (cq *CommandQueue) Finish() {
	fmt.Printf("Supress.\n")
	C.clFinish(cq.item);
}

type Context struct {
	item C.cl_context
}

func CreateContext(platform *Platform, devType DeviceType) *Context {
	var ctx Context
	var error C.cl_int
	properties := []C.cl_context_properties{CL_CONTEXT_PLATFORM, C.platform2property(platform.item), C.cl_context_properties(0)}
	ctx.item = C.clCreateContextFromType(&properties[0], C.cl_device_type(devType), nil, nil, &error)
	fmt.Printf("Context error:%d\n", error)
	return &ctx
}

type Program struct {
	item C.cl_program
}

func CreateProgram(context *Context, sources []string) *Program {
	runtime.LockOSThread()
	var program Program
	var error C.cl_int

	csources := C.makeCharArray(C.int(len(sources)))
	clenghts := make([]C.size_t, len(sources))

	defer C.freeCharArray(csources, C.int(len(sources)))

	for i := 0; i < len(sources); i++ {
		C.setArrayString(csources, C.CString(sources[i]), C.int(i))
		clenghts[i] = C.size_t(len(sources[i]))
		fmt.Printf("Program log:%d %s\n",clenghts[i], sources[i])
	}

	program.item = C.clCreateProgramWithSource(
				context.item,
				C.cl_uint(len(sources)),
				csources,
				&clenghts[0],
				&error)

	fmt.Printf("Program error:%d\n", error)

	return &program
}

type Kernel struct {
	item C.cl_kernel
}


func CreateKernel(program *Program, name string) *Kernel {
	var kernel Kernel
	var error C.cl_int
	kernel.item = C.clCreateKernel(program.item, C.CString(name), &error)
	fmt.Printf("Kernel error:%d\n", error)
	return &kernel
}

func (*Kernel) Foo() {

}

type Buffer struct {
	item C.cl_mem
}
