#include "nss_fnmaker.h"
#include <nss.h>
#include <stddef.h>
#include <shadow.h>

enum nss_status __FNAME(setspent)(void);
enum nss_status __FNAME(endspent)(void);
enum nss_status __FNAME(getspent_r)(struct spwd *result, char *buffer, size_t buflen, int *errnop);
enum nss_status __FNAME(getspnam_r)(const char *name, struct spwd *result, char *buffer, size_t buflen, int *errnop);
