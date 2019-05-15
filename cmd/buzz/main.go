//each package is itself a small program
//a single unit of change with a single responsibility
package main

import (
	"github.com/volodimyr/ood-fizzbuzz/pkg/factory"
	"github.com/volodimyr/ood-fizzbuzz/pkg/game"
	"os"
)

func main() {
	f := factory.NewCreator()

	var g game.FizzBuzzer
	var err error
	if g, err = f.Create(game.TypeBuzz, 100, os.Stdout); err != nil {
		panic(err)
	}

	g.On()
}
