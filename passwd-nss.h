#include "nss_fnmaker.h"
#include <nss.h>
#include <stddef.h>
#include <pwd.h>

extern enum nss_status __FNAME(setpwent)();
extern enum nss_status __FNAME(endpwent)();
extern enum nss_status __FNAME(getpwent_r)(struct passwd *p, char *buf, size_t len, int *errnop);
extern enum nss_status __FNAME(getpwnam_r)(const char *name, struct passwd *, char *buf, size_t len, int *errnop);
extern enum nss_status __FNAME(getpwuid_r)(uid_t uid, struct passwd *, char *buf, size_t len, int *errnop);
