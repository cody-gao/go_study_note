package main

func main() {
	x := 100
	f := test(x)
	f()
}

func test(x int) func() {
	return func() {
		println(x)
	}
}
