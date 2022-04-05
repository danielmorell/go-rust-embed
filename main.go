package main

import (
	_ "embed"
	"fmt"
	"github.com/amenzhinsky/go-memexec"
)

//go:generate rustc main.rs -o bin/main_rust
//go:embed bin/main_rust
var rustBinary []byte

func main() {
	fmt.Println("Hello from Go")
	exe, err := memexec.New(rustBinary)
	if err != nil {
		panic(err)
	}
	defer exe.Close()

	cmd := exe.Command("Hello to Rust from Go")
	result, err := cmd.Output()
	fmt.Println(string(result))
}
