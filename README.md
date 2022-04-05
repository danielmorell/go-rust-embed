# Embed a Rust Binary In a Go Program

This is a really simple proof of concept that embeds a Rust binary in a Go program. The
Rust binary can be executed from the Go program, and is also shipped as part of the Go
binary.

You can run this little example as follows...

```
$ go generate . && go run .
```

And build it like so...

```
$ go generate . && go build .
```

The `go generate .` part will discover the `//go:generate` comment in 
[main.go:9](main.go#9) and execute the rest of the line as a shell command.

On [main.go:10](main.go#10) we use `//go:embed` to include the compiled rust binary into
the Go program when it is compiled.