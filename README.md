⚠️ Avoid relying on Go for shared libraries. The [runtime](https://pkg.go.dev/runtime) uses signals it assumed to be unused for concurrency management. Assumptions like this are why the verion of Go can not be updated for `go-libnss` and make Go less predicable when building shared libraries.

The current recommendation is to build services to be compatible with existing protocols like LDAP.

# Go-Libnss
This is an abstracted library in Go that wraps around libnss, allowing you to write new nss modules purely in Go.

## What's Working?
Right now we only have binding for `passwd`, `group`, and `shadow`. This may change in the future to also implement other NSS features. These are just the most obvious first targets.

## How do I use it?
First you should `go get` the package:
```
go get github.com/protosam/go-libnss
```
Take a gander at the `example` directory. It includes information on how to compile the example and you can use the implementation as a boiler plate for your own project.

## License
MIT License

# Projects Using Go-Libnss
[Go-Libnss-Etcd](https://github.com/protosam/go-libnss-etcd) - An NSS package that stores passwd, group, and shadow data in etcd. It also comes with CLI tools for managing users and groups stored in etcd.
