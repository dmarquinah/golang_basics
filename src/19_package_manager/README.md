# Package Managing in Go
We have a way to handle package installation as a module in our local machine.

Previous versions of Go used to install packages with:

```console
go get golang.org/x/tour
```

But neither the current version uses `go get` nor have the package `golang.org/x/tour` available.

Currently packages can be installed using the command:

```console
go install package
```

As an example, we have the updated version of the Golang Tour on:

```console
go install golang.org/x/website/tour@latest
```

And the installed package can be executed from the $GOPATH with:
```console
tour
```