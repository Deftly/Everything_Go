package greetings

import (
  "errors"
  "fmt"
  "math/rand"
  "time"
)

// Hello returns a greeting for the named person
func Hello(name string) (string, error) {
  if name == "" {
    return "", errors.New("empty name")
  }
  message := fmt.Sprintf(randomFormat(), name)
  return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
  messages := make(map[string]string)

  for _, name := range names {
    message, err := Hello(name)
    if err != nil {
      return nil, err
    }
    messages[name] = message
  }
  return messages, nil
}

// init sets initial values for variables used in the function.
func init() {
  rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages. The returned
// message is selected at random.
func randomFormat() string {
  // A slice of message formats.
  formats := []string {
    "Hi, %v. Welcome!",
    "Great to see you, %v!",
    "Hail, %v! Well met!",
  }

  // Return a randomly selected message format.
  return formats[rand.Intn(len(formats))]
}
