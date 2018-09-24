package main

import (
  "fmt"
  "strconv"
  "log"
  "errors"
)

type Person struct {
  Name string
  Age int
}

func (p *Person) IncreaceAge() {
  p.Age += 1;
}

func (p *Poppler) GetName() string {
  return p.Name;
}

func (p *Poppler) GoString() string {
  return p.Name;
}

type Poppler struct {
  *Person
  Role string
  Capacity int
}

func StringToInt(cadena string) func () int {
  return func () int {
    i, err := strconv.Atoi(cadena)
    if err != nil {
      log.Fatal(err)
    }
    return i
  }
}

func (p *Poppler) IncreaceCapacity() {
 p.Capacity += 1;
}

type GoStringer interface {
    GoString() string
}

// Check if ANY has a method GoString first cating to GoStringer
// if is posible.then check for method GoString (Something like fmt)
func ToString(any interface{}) (ttS string, err error) {
    if v, ok := any.(GoStringer); ok {
        return v.GoString(), nil
    }
    switch v := any.(type) { // Check for the type of "any"
    case int:
      return strconv.Itoa(v), nil
    case float64:
      return strconv.FormatFloat(v, 'f', -1, 64), nil
    case string:
      return v, nil
    case bool:
      return strconv.FormatBool(v), nil
    }
    return "", errors.New("TypeError \n")
}

// Create a new Poppler
func NewPoppler(name string, role string, age int) *Poppler {
  return &Poppler{Person: &Person{name, age}, Role: role}
}

func main() {
  Boss := NewPoppler("Jefe", "CEO", 30)
  Boss.IncreaceAge();

  fmt.Printf("hello, world %s \n", Boss.GetName())
  fmt.Printf("With Age: %d \n", Boss.Age)
  fmt.Printf("Clojure StringToInt: %d \n", StringToInt("10")())

  name, err := ToString(Boss)
  if err != nil {
    log.Panic(err)
  }

  fmt.Printf("ToString: %s \n", name)
}
