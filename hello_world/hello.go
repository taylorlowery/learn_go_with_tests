package main

import "fmt"

const (
	spanish                 = "Spanish"
	french                  = "French"
	american                = "American"
	english                 = "English"
	cantonese               = "Cantonese"
	frenchGreetingPrefix    = "Bonjour"
	spanishGreetingPrefix   = "Hola"
	yeehawGreetingPrefix    = "Howdy"
	englishGreetingPrefix   = "Hello"
	cantoneseGreetingPrefix = "Nei Ho"
)

func getGreeting(language string) (greeting string) {
	switch language {
	case english:
		greeting = englishGreetingPrefix
	case french:
		greeting = frenchGreetingPrefix
	case spanish:
		greeting = spanishGreetingPrefix
	case cantonese:
		greeting = cantoneseGreetingPrefix
	default:
		greeting = yeehawGreetingPrefix
	}
	return
}

func Greet(name, language string) string {
	if name == "" {
		name = "World"
	}
	greeting := getGreeting(language)
	return fmt.Sprintf("%s, %s!", greeting, name)
}

func main() {
	fmt.Println(Greet("Sam Yuu", cantonese))
}
