package vulkango

import "unsafe"

type Structure interface {
	getType() VkStructureType
	toC() (*C.void, func())
}

type VkInstanceCreateInfo struct {
	Next                   Structure
	Flags                  VkInstanceCreateFlags
	PApplicationInfo       *VkApplicationInfo
	PEnabledLayerNames     []string
	PEnabledExtensionNames []string
}

func (s *VkInstanceCreateInfo) getType() VkStructureType {
	return VkStructureTypeInstanceCreateInfo
}

func (s *VkInstanceCreateInfo) toC() (*C.void, func()) {

	// --------- Initalization ---------

	// allocate c struct
	p := (*C.VkInstanceCreateInfo)(C.malloc(
		C.size_t(
			unsafe.Sizeof(C.VkInstanceCreateInfo{}),
		),
	))

	// cancel array
	cancels := make([]func(), 0)

	// --------- Copy ---------

	// structure type
	p.sType = (C.VkStructureType)(s.GetType())

	// Any pointer cpy
	if s.Next != nil {
		next, cancel := s.Next.toC()
		cancels = append(cancels, cancel)
		p.pNext = next;
	}

	// flag / primative types
	p.Flags = (C.VkInstanceCreateFlags)(s.Flags)

	// ptr copy
	if s.PApplicationInfo != nil {
		appInfo, cancel := s.PApplicationInfo.toC()
		cancels = append(cancels, cancel)
		p.PApplicationInfo = (*C.VkApplicationInfo)(appInfo)
	}

	// Array / str allocation
	p.enabledLayerCount = len(s.PEnabledLayerNames)	
	p.ppEnabledLayerNames = make([]*C.char, 0)
	for _, e := range s.PEnabledLayerNames {
		cStr := C.CString(e)
		cancels = append(cancels, func() {C.free(unsafe.Pointer(cStr))})
		p.ppEnabledLayerNames = append(p.ppEnabledLayerNames, cStr)
	}

	// ...

	// --------- Finalization ---------

	return p, func() {
		// free c struct
		C.free(unsafe.Pointer(p))

		// free children
		for _, c := range cancels {
			c()
		}
	}
}
