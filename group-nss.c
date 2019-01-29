#include "nss_fnmaker.h"
#include "nss.h"
#include "_cgo_export.h"
#include <string.h>
#include <grp.h>

enum nss_status __FNAME(setgrent) (int stayopen) {
	return go_setgrent(stayopen);
}

enum nss_status __FNAME(endgrent) (void) {
	return go_endgrent();
}

enum nss_status __FNAME(getgrent_r) (struct group *result, char *buffer, size_t buflen, int *errnop) {
	return go_getgrent_r(result, buffer, buflen, errnop);
}

enum nss_status __FNAME(getgrgid_r) (gid_t gid, struct group *grp, char *buffer, size_t buflen, int *errnop) {
	return go_getgrgid_r(gid, grp, buffer, buflen, errnop);
}

enum nss_status __FNAME(getgrnam_r) (const char *name, struct group *grp, char *buffer, size_t buflen, int *errnop) {
	GoString goname = {name, strlen(name) };
	return go_getgrnam_r(goname, grp, buffer, buflen, errnop);
}
