Important: Any shared libraries made with CGO will inherit any design choices of the Go runtime as well as drawbacks involved with interoperability. This seems to hold true for Rust, Swift, and to my surprise C++. As such, it makes more sense to either make an ldap service to interface with libnss-ldapd/sssd, or write libnss modules in C. Fixing and using go-libnss comes with drawbacks that will likely set your project up for failure.

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
