/*
This macro is used to generate function names in C for new libnss libraries.
*/
#ifndef __LIB_NSS_NAME
#define __LIB_NSS_NAME gonss
#endif

#define __XPASTER(x,y,z) x ## y ## _ ## z
#define __PASTER(x,y,z) __XPASTER(x,y,z)
#define __FNAME(x) __PASTER(_nss_, __LIB_NSS_NAME, x)

