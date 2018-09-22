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

func ToString(any interface{}) (ttS string, err error) {
    if v, ok := any.(GoStringer); ok {
        return v.GoString(), nil
    }
    switch v := any.(type) {
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


func main() {
  Boss := &Poppler{
   Person: &Person{"Jefe", 30},
   Role: "CEO",
  }
  Boss.IncreaceAge();

  fmt.Printf("hello, world %s \n", Boss.GetName());
  fmt.Printf("With Age: %d \n", Boss.Age);
  fmt.Printf("Clojure StringToInt: %d \n", StringToInt("10")())
  name, err := ToString(Boss);
  if err != nil {
    fmt.Print(err)
  }
  fmt.Printf("ToString: %s \n", name);
}
