package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names with messages
	messages := make(map[string]string)
	// loop through names
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// in the map associate message with name
		messages[name] = message+"\n"
	}
	return messages, nil
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
	// seed or it will be set to 1
	rand.Seed(time.Now().UnixNano())
    // A slice of message formats.
    formats := []string{
        "1 Hi, %v. Welcome!",
        "2 Great to see you, %v!",
        "3 Hail, %v! Well met!",
    }

    // Return a randomly selected message format by specifying
    // a random index for the slice of formats.
    return formats[rand.Intn(len(formats))]
}
