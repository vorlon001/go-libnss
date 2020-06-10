package nss

//#include <grp.h>
//#include <errno.h>
//#include <stdlib.h>
import "C"

import (
	"bytes"
	"syscall"
	"unsafe"

	. "github.com/protosam/go-libnss/structs"
)

var entries_group = make([]Group, 0)
var entry_index_group int

//export go_setgrent
func go_setgrent(stayopen C.int) Status {
	var status Status
	status, entries_group = implemented.GroupAll()
	entry_index_group = 0
	return status
}

//export go_endgrent
func go_endgrent() Status {
	entries_group = make([]Group, 0)
	entry_index_group = 0
	return StatusSuccess
}

//export go_getgrent_r
func go_getgrent_r(grp *C.struct_group, buf *C.char, buflen C.size_t, errnop *C.int) Status {
	if entry_index_group == len(entries_group) {
		return StatusNotfound
	}
	setCGroup(&entries_group[entry_index_group], grp, buf, buflen, errnop)
	entry_index_group++
	return StatusSuccess
}

//export go_getgrnam_r
func go_getgrnam_r(name string, grp *C.struct_group, buf *C.char, buflen C.size_t, errnop *C.int) Status {
	status, group := implemented.GroupByName(name)
	if status != StatusSuccess {
		return status
	}
	setCGroup(&group, grp, buf, buflen, errnop)
	return StatusSuccess
}

//export go_getgrgid_r
func go_getgrgid_r(gid uint, grp *C.struct_group, buf *C.char, buflen C.size_t, errnop *C.int) Status {
	status, group := implemented.GroupByGid(gid)
	if status != StatusSuccess {
		return status
	}
	setCGroup(&group, grp, buf, buflen, errnop)
	return StatusSuccess
}

// Sets the C values for libnss
func setCGroup(p *Group, grp *C.struct_group, buf *C.char, buflen C.size_t, errnop *C.int) Status {
	size := int(buflen)

	if len(p.Groupname)+len(p.Password)+5 > size {
		*errnop = C.int(syscall.EAGAIN)
		return StatusTryagain
	}

	gobuf := C.GoBytes(unsafe.Pointer(buf), C.int(buflen))
	b := bytes.NewBuffer(gobuf)
	b.Reset()

	grp.gr_name = (*C.char)(unsafe.Pointer(&gobuf[b.Len()]))
	b.WriteString(p.Groupname)
	b.WriteByte(0)
	size -= len(p.Groupname) + 1

	grp.gr_passwd = (*C.char)(unsafe.Pointer(&gobuf[b.Len()]))
	b.WriteString(p.Password)
	b.WriteByte(0)
	size -= len(p.Password) + 1

	grp.gr_gid = C.uint(p.GID)

	// Create members list
	//
	// Calculate the size of the array.
	// sizeof(char*) * (len(src) + 1)
	sizeOfCharS := unsafe.Sizeof(uintptr(0))
	length := int(sizeOfCharS) * (len(p.Members) + 1)
	if length > size {
		*errnop = C.int(syscall.EAGAIN)
		return StatusTryagain
	}

	// Set the address of the array.
	bufp := (**C.char)(unsafe.Pointer(&gobuf[b.Len()]))
	b.Write(make([]byte, length))
	size -= length
	grp.gr_mem = bufp

	for _, s := range p.Members {
		// Check buflen
		length = len(s) + 1
		if length > size {
			*errnop = C.int(syscall.EAGAIN)
			return StatusTryagain
		}

		// Set the address of each element
		*bufp = (*C.char)(unsafe.Pointer(&gobuf[b.Len()]))

		// Write element
		b.WriteString(s)
		b.WriteByte(0)
		size -= length

		// Go to next element
		bufp = (**C.char)(unsafe.Pointer(uintptr(unsafe.Pointer(bufp)) + sizeOfCharS))
	}

	// End the array with a nil pointer
	*bufp = nil

	return StatusSuccess
}
