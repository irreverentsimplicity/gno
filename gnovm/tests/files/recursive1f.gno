package main

func main() {
	type S struct {
		T S
	}

	var a, b S
	println(a == b)
}

// Error:
// main/files/recursive1f.gno:3:1: invalid recursive type: S -> S
