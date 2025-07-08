package main

import "fmt"

const (
	french  = "French"
	spanish = "Spanish"

	englishHello = "Hello"
	spanishHello = "Hola"
	frenchHello  = "Bonjour"
	defaultValue = ":)"
)

func main() {
	fmt.Println(Hello("John Doe"))
}

func Hello(name string) string {
	if name == "" {
		name = defaultValue
	}

	return fmt.Sprintf("%s %s", englishHello, name)
}

func HelloLanguage(name, language string) string {
	if name == "" {
		name = defaultValue
	}

	return fmt.Sprintf("%s %s", greetingPrefix(language), name)
}

/*
* Named return value (prefix string)
* this will create a variable create prefix in your function
* we can just use return rather than return prefix
 */
func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHello
	case french:
		prefix = frenchHello
	default:
		prefix = englishHello
	}

	return
}
