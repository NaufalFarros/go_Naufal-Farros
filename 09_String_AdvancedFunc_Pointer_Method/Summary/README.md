## Rangkuman Materi String,Advanced Function, Pointer, Package, & Error Handling
* String 
** Len String
``` 
sentence := "Hello"
	lenSentence := len(sentence)
	fmt.Println("lenSentence:", lenSentence)
```
** Compare
```
str1 := "abc"
	str2 := "abd"

	fmt.Println("str1 == str2:", str1 == str2)
```
** Contains return bool
```
res := strings.Contains(str, substr)

	fmt.Println(res)
```
** SubString
```
value := "cat;dog"

	substring := value[4:len(value)]
	fmt.Println(substring)
```
** Replace 
```
s := "this[things] i would like to remove"
	t := strings.Replace(s, "[", " ", -1)
	fmt.Printf("%s\n", t)
```
** Insert 
```
p := "green"
	index := 2
	q := p[:index] + "HI" + p[index:]

	fmt.Println(p, q)
```
* Function 

** Variadic Function

```
package main

import (
  "fmt"
)

func sum(numbers ...int) int {

  var total int = 0
  for _, number := range numbers {
    total += number
  }
  return total
}

func main() {
  avg := sum(2, 4, 3, 5)
  fmt.Println(avg)
}

```

** Anonymous function

```
package main

import "fmt"

func main() {
  // Anonymous function
  func() {
    fmt.Println("Welcome! to GeeksforGeeks")
  }()

  // Assigning an anonymous function to a variable
  value := func() {
    fmt.Println("Welcome! to GeeksforGeeks")
  }
  value()

  // Passing arguments in anonymous function
  func(sentence string) {
    fmt.Println(sentence)
  }("GeeksforGeeks")
}
```
** Closure 
``` 
package main

import "fmt"

func newCounter() func() int {
  count := 0
  return func() int {
    count += 1
    return count
  }
}

func main() {
  counter := newCounter()
  fmt.Println(counter())
  fmt.Println(counter())
}

```

** Defer Function

```
package main

import "fmt"

func main() {
  defer func() {
    fmt.Println("Later")
  }()

  fmt.Println("First")

}
```

* Pointer 
```
package main

import "fmt"

func main() {
  var name string = "John"
  var nameAddress *string = &name
  fmt.Println("name (value)   :", name) // John
  fmt.Println("name (address) :", &name) // 0xc20800a220
  fmt.Println("nameAddress (value)   :", *nameAddress) // John
  fmt.Println("nameAddress (address) :", nameAddress) // 0xc20800a220

  name = "Doe"
  
  fmt.Println("")  
  fmt.Println("name (value)   :", name) // Doe
  fmt.Println("name (address) :", &name) // 0xc20800a220
  fmt.Println("nameAddress (value)   :", *nameAddress) // Doe
  fmt.Println("nameAddress (address) :", nameAddress) // 0xc20800a220
}
```

* Struct adalah sebuah objek 

```
package main

import "fmt"

type Person struct {
  FirstName string
  LastName  string
  Age       int
}

func main() {
   // long declaration
  var Person0 Person
  Person0.FirstName = "Muchson"
  Person0.LastName = "Ibi"
  Person0.Age = 27
  fmt.Println(Person0.FirstName, Person0.LastName, Person0.Age)

  // long declaration with assigned value
  var Person1 = Person{"Rizky", "Kurniawan", 26}
  fmt.Println(Person1)

  // long declaration with assigned value each name fields
  var Person2 = Person{
    FirstName: "Iswanul",
    LastName:  "Umam",
    Age:       25,
  }
  fmt.Println(Person2)

  // sort declaration
  Person3 := Person{"Pranadya", "Bagus", 23}
  fmt.Println(Person3)

  // short declaration with new keyword
  Person4 := new(Person)
  Person4.FirstName = "Muhammad"
  Person4.LastName = "Ismail"
  Person4.Age = 30
  fmt.Println(*Person4)

  }
```
* Method 

```
package main

import "fmt"

type Employee struct {
    FirstName, LastName string
}

func (e Employee) fullName() string {
    return e.FirstName + " " + e.LastName
}

func main() {
    e := Employee{
        FirstName: "Ross",
        LastName:  "Geller",
    }
    fmt.Println(e.fullName())
}
```

* Interface 

```
package main

import "fmt"

type calculate interface {
  large() int
}

type square struct {
  side int
}

func (s square) large() int {
  return s.side * s.side
}

func main() {
  var dimResult calculate
  dimResult = square{10}
  fmt.Println("large square :", dimResult.large())
}
```
* Error Handling 

** Error Object
```
package main

import "fmt"

import (
  "errors"
)

func myFunc(i int) (int, error) {
  if i <= 0 {
    return -1, errors.New("should be greater than zero")
  }
  return i, nil
}

func main() {
  result, err := myFunc(-1)
  fmt.Println(result, err)
}

```
** Error Panic Recover
```

package main

import "fmt"

func myMethod() {
  defer func() {
    if err := recover(); err != nil {
      fmt.Println("Error Message:", err)
    }
  }()

  anOddCondition := true
  if anOddCondition {
    panic("I am panicking")
  }
}

func main() {
  myMethod()
}
```






