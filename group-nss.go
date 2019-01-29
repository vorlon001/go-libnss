package nss
//#include <grp.h>
//#include <errno.h>
//#include <stdlib.h>
/*
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
*/
import "C"

import(
	"bytes"
	"syscall"
	"unsafe"
)


var entries_group = make([]Group, 0)
var entry_index_group int

//export go_setgrent
func go_setgrent(stayopen C.int) Status {
	var status Status
	status, entries_group = implemented.GroupAll()
	entry_index_group = 0
	return status;
}

//export go_endgrent
func go_endgrent() Status {
	entries_group = make([]Group, 0)
	entry_index_group = 0
	return StatusSuccess;
}

//export go_getgrent_r
func go_getgrent_r(grp *C.struct_group, buf *C.char, buflen C.size_t, errnop *C.int) Status {
	if entry_index_group == len(entries_group) {
		return StatusNotfound
	}
	setCGroup(&entries_group[entry_index_group], grp , buf, buflen, errnop)
	entry_index_group++
	return StatusSuccess;
}

//export go_getgrnam_r
func go_getgrnam_r(name string, grp *C.struct_group, buf *C.char, buflen C.size_t, errnop *C.int) Status {
	status, group := implemented.GroupByName(name)
	if status != StatusSuccess {
		return status
	}
	setCGroup(&group, grp , buf, buflen, errnop)
	return StatusSuccess;
}

//export go_getgrgid_r
func go_getgrgid_r(gid GID, grp *C.struct_group, buf *C.char, buflen C.size_t, errnop *C.int) Status {
	status, group := implemented.GroupByGid(gid)
	if status != StatusSuccess {
		return status
	}
	setCGroup(&group, grp , buf, buflen, errnop)
	return StatusSuccess;
}

// Sets the C values for libnss
func setCGroup(p *Group, grp *C.struct_group, buf *C.char, buflen C.size_t, errnop *C.int) Status {
	// TODO: Need to add length for Members....
	if len(p.Groupname)+len(p.Password)+5 > int(buflen) {
		*errnop = C.int(syscall.EAGAIN)
		return StatusTryagain
	}

	gobuf := C.GoBytes(unsafe.Pointer(buf), C.int(buflen))
	b := bytes.NewBuffer(gobuf)
	b.Reset()

	grp.gr_name = (*C.char)(unsafe.Pointer(&gobuf[b.Len()]))
	b.WriteString(p.Groupname)
	b.WriteByte(0)

	grp.gr_passwd = (*C.char)(unsafe.Pointer(&gobuf[b.Len()]))
	b.WriteString("x")
	b.WriteByte(0)

	grp.gr_gid = C.uint(p.GID)
	
	// ################ MAKING **C.char in GO! 
	// Making a list of the members...
	// NOTE: There has to be a better way to do this.
	// I'm also making an assumption the process running this dies, freeing up the memory.

	grp.gr_mem = C.makeCharArray(C.int(len(p.Members)))
	//defer C.freeCharArray(grp.gr_mem, C.int(len(p.Members)))
	for i, u := range p.Members {
		C.setArrayString(grp.gr_mem, C.CString(u), C.int(i))
	}
	// ################ DONE MAKING **C.char in GO! 

	return StatusSuccess
}
