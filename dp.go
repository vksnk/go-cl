package dp

/*
#include <stdlib.h>
#include <CL/cl.h>

static cl_context_properties* getPropArray(cl_platform_id pid) {
	cl_context_properties* props = malloc(sizeof(cl_context_properties) * 3);
	props[0] = CL_CONTEXT_PLATFORM;
	props[1] = (cl_context_properties)pid;
	props[2] = 0;
	return props;
}
*/
import "C"
import "fmt"
import "runtime"

func Run() {
	runtime.LockOSThread()
	fmt.Printf("Start of direct hellocl port\n")

	var err C.cl_int;
	var platform C.cl_platform_id
	var device_id C.cl_device_id
	var device_num C.cl_uint
	var context C.cl_context
	var commands C.cl_command_queue

	err = C.clGetPlatformIDs(1, &platform, nil)

	err = C.clGetDeviceIDs(platform, C.CL_DEVICE_TYPE_GPU, 1, &device_id, &device_num);

	fmt.Printf("%d\n", device_num)

	if err != C.CL_SUCCESS {
		fmt.Printf("Error: Failed to create a device group !\n")
		return
	}

	//context = C.clCreateContext(C.getPropArray(platform), 1, &device_id, nil, nil, &err);

	context = C.clCreateContextFromType(C.getPropArray(platform), C.CL_DEVICE_TYPE_GPU, nil, nil, &err);

	if err != C.CL_SUCCESS {
		fmt.Printf("Error: Failed to create a compute context %v!\n", err);
		return
	}

	commands = C.clCreateCommandQueue(context, device_id, 0, &err);

	if err != C.CL_SUCCESS {
		fmt.Printf("Error: Failed to create a command commands!\n");
		return
	}

	if commands == nil {
	}
}
