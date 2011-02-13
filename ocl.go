package ocl

//#include <CL/cl.h>
import "C"

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


