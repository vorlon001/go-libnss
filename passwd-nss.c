#include "nss_fnmaker.h"
#include "nss.h"
#include "_cgo_export.h"
#include <string.h>
#include <pwd.h>

enum nss_status __FNAME(setpwent) () {
	return go_setpwent();
}

enum nss_status __FNAME(endpwent) () {
	return go_endpwent();
}

enum nss_status __FNAME(getpwent_r) (struct passwd *p, char *buf, size_t len, int *errnop) {
	return go_getpwent(p, buf, len, errnop);
}

enum nss_status __FNAME(getpwnam_r) (const char *name, struct passwd *p, char *buf, size_t len, int *errnop) {
	GoString goname = {name, strlen(name) };
	return go_getpwnam(goname, p, buf, len, errnop);
}

enum nss_status __FNAME(getpwuid_r) (uid_t uid, struct passwd *p, char *buf, size_t len, int *errnop) {
	return go_getpwuid(uid, p, buf, len, errnop);
}
