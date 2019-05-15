//structure functions / types into packages
//that exhibit natural cohesion
package game

import (
	"io"
	"log"
	"strconv"
)

const (
	TypeFizz     = "fizz"
	TypeBuzz     = "buzz"
	TypeFizzBuzz = "fizzbuzz"
)

type FizzBuzzer interface {
	GetLimit() int
	check(int) string
	On()
}

type Fizz struct {
	limit int
}

func (f Fizz) check(n int) string {
	if n%3 != 0 {
		return strconv.Itoa(n)
	}

	return f.String()
}

func (f Fizz) GetLimit() int {
	return f.limit
}

func (f Fizz) String() string {
	return TypeFizz
}

func (f Fizz) On() {
	for i := 1; i <= f.limit; i++ {
		log.Println(f.check(i))
	}
}

func NewFizz(print io.Writer, limit int) FizzBuzzer {
	log.SetOutput(print)

	return &Fizz{
		limit: limit,
	}
}

type Buzz struct {
	limit int
}

func (b Buzz) check(n int) string {
	if n%5 != 0 {
		return strconv.Itoa(n)
	}

	return b.String()
}

func (b Buzz) GetLimit() int {
	return b.limit
}

func (b Buzz) String() string {
	return TypeBuzz
}

func (b Buzz) On() {
	for i := 1; i <= b.limit; i++ {
		log.Println(b.check(i))
	}
}

func NewBuzz(print io.Writer, limit int) FizzBuzzer {
	log.SetOutput(print)

	return &Buzz{
		limit: limit,
	}
}

//FizzBuzz
type FizzBuzz struct {
	//embedding is a powerful tool
	//which allows Golang's types
	//to be open for extension
	Fizz
	Buzz

	//encapsulation: private / public (exported, unexported)
	limit int
}

func (fb FizzBuzz) GetLimit() int {
	return fb.limit
}

func (fb FizzBuzz) String() string {
	return TypeFizzBuzz
}

func (fb FizzBuzz) check(n int) string {
	var result string

	if fRes := fb.Fizz.check(n); fRes == fb.Fizz.String() {
		result += fRes
	}
	if bRes := fb.Buzz.check(n); bRes == fb.Buzz.String() {
		result += bRes
	}

	return result
}

func (fb FizzBuzz) On() {
	for i := 1; i <= fb.limit; i++ {

		switch result := fb.check(i); result {

		case fb.String(), fb.Fizz.String(), fb.Buzz.String():
			log.Println(result)

		default:
			log.Println(strconv.Itoa(i))
		}
	}
}

//clients shouldn't be forced to depend on methods they don't use
//which means functions, methods should only depend on the behaviour that they need
func NewFizzBuzz(print io.Writer, limit int) FizzBuzzer {
	log.SetOutput(print)

	return &FizzBuzz{
		limit: limit,
	}
}
