package main

func main() {
	println(Hello("César"))
}

const helloPrefix = "Hello, "
const helloSufix = "!"

func Hello(name string) string {
	return helloPrefix + name + helloSufix
}
