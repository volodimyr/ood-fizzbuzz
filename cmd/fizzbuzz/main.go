//each package is itself a small program
//a single unit of change with a single responsibility
package main

import (
	"github.com/volodimyr/ood-fizzbuzz/pkg/factory"
	"github.com/volodimyr/ood-fizzbuzz/pkg/game"
	"os"
)

//import is a source level of coupling

var path = "log.txt"

func main() {
	fact := factory.NewCreator()

	f, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	var g game.FizzBuzzer
	if g, err = fact.Create(game.TypeFizzBuzz, 100, f); err != nil {
		panic(err)
	}

	g.On()
}
