#include "nss_fnmaker.h"
#include <nss.h>
#include <stddef.h>
#include <grp.h>

extern enum nss_status __FNAME(setgrent)(int stayopen);
extern enum nss_status __FNAME(endgrent)(void);
extern enum nss_status __FNAME(getgrent_r)(struct group *result, char *buffer, size_t buflen, int *errnop);
extern enum nss_status __FNAME(getgrgid_r)(gid_t gid, struct group *grp, char *buffer, size_t buflen, int *errnop);
extern enum nss_status __FNAME(getgrnam_r)(const char *name, struct group *grp, char *buffer, size_t buflen, int *errnop);
