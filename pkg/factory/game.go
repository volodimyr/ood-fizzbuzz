package factory

import (
	"errors"
	"github.com/volodimyr/ood-fizzbuzz/pkg/game"
	"io"
)

var (
	errNoSuchFizzBuzzGame = errors.New("there is no such fizz buzz game")
)

type GameCreator interface {
	Create(_type string, limit int, print io.Writer) (game.FizzBuzzer, error)
}

func NewCreator() GameCreator {
	return &Creator{}
}

type Creator struct{}

//express dependencies between packages in terms of interfaces not concrete types
func (c Creator) Create(_type string, limit int, print io.Writer) (game.FizzBuzzer, error) {
	switch _type {
	case game.TypeFizz:
		return game.NewFizz(print, limit), nil

	case game.TypeBuzz:
		return game.NewBuzz(print, limit), nil

	case game.TypeFizzBuzz:
		return game.NewFizzBuzz(print, limit), nil
	}

	return nil, errNoSuchFizzBuzzGame
}
