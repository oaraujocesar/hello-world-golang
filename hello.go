package main

func main() {
	println(Hello("César", "Spanish"))
}

const spanish = "Spanish"
const french = "French"
const portuguese = "Portuguese"

const helloPrefixEnglish = "Hello, "
const helloPrefixPortuguese = "Olá, "
const helloPrefixSpanish = "Hola, "
const helloPrefixFrench = "Bonjour, "

const helloSufix = "!"

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return greetingsPrefix(language) + name + helloSufix
}

func greetingsPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = helloPrefixSpanish
	case french:
		prefix = helloPrefixFrench
	case portuguese:
		prefix = helloPrefixPortuguese
	default:
		prefix = helloPrefixEnglish
	}

	return
}
