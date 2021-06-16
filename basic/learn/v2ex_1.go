//https://www.v2ex.com/t/614689#reply0
//学习上面链接中介绍的小方法，用于基于其他库的二次开发

package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {

	log.Logger.Debug().
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")

	u := Ulog{Logger: log.Logger}
	u.Debug().
		ID("hello").
		Str("Scale", "833 cents").
		Float64("Interval", 833.09).
		Msg("Fibonacci is everywhere")
}

type Ulog struct {
	zerolog.Logger
}

func (u *Ulog) Debug() *event {
	return &event{u.Logger.Debug()}
}

func (u *Ulog) Info() *event {
	return &event{u.Logger.Info()}
}

func (u *Ulog) Warn() *event {
	return &event{u.Logger.Warn()}
}

func (u *Ulog) Error() *event {
	return &event{u.Logger.Error()}
}

type event struct {
	*zerolog.Event
}

func (e *event) ID(id string) *event {
	return &event{e.Str("ID", id)}
}


