# Go-Libnss Example
We have a mock database that is built in `testdata.go` and `implementation.go` is a test structure with methods that utilize that test database.

The goal is to provide a working example with populated data so that you can get off something the ground with `Go-Libnss`.

# How to Compile
The following command will build `libnss_gotest.so.2`, which you then put in `/lib64/`:

```
CGO_CFLAGS="-g -O2 -D __LIB_NSS_NAME=gotest" go build --buildmode=c-shared -o libnss_gotest.so.2 implementation.go testdata.go
```

The `CGO_CFLAGS` portion is important. A value has to be set for `__LIB_NSS_NAME` and the resulting object file should be set to `libnss_{__LIB_NSS_NAME}.so.2`. There are some built in C macros to generate function names that libnss relies on.

After you have placed the shared library in `/lib64/`, you then edit `/etc/nsswitch.conf` like so (replacing `{__LIB_NSS_NAME}` with whatever your library name is):

```
passwd:     files {__LIB_NSS_NAME} sss
shadow:     files {__LIB_NSS_NAME} sss
group:      files {__LIB_NSS_NAME} sss
```

Once that is done, you should see results from `getent passwd`, `getent group`, and `getent shadow` commands.
