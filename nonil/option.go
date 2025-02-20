package nonil

import "errors"

type Option[T any] interface {
	// Handle, along with IfSome and IfNone, give you the ability to interact with an Option
	// through functions only called when there is something or nothing, as appropriate.
	Handle(ifsome func(t T), ifnone func())
	IfSome(func(t T))
	IfNone(func())

	// The methods below allow operations with the Option type without callbacks. 
	// This allows for greater performance in exchange for the possibility of panicing with `Some()`. 
	// The difference is largely due to the overhead of callback functions. 
	// This cost is not likely to scale with the complexity of the callback

	IsSome() bool
	// Some returns the internal value, but panics if Option is Nothing.
	// Programmers should first check if value is available with IsSome or IsNone.
	Some() T
	IsNone() bool
}

type some[T any] struct {
	wrap T
}

func Some[T any](wrap T) Option[T] {
	return some[T]{
		wrap: wrap,
	}
}

func (s some[T]) Handle(ifsome func(t T), ifnone func()) { ifsome(s.wrap) }
func (s some[T]) IfSome(f func(t T))                     { f(s.wrap) }
func (s some[T]) IfNone(_ func())                        {}
func (s some[T]) IsSome() bool                           { return true }
func (s some[T]) Some() T                                { return s.wrap }
func (s some[T]) IsNone() bool                           { return false }

type none[T any] struct{}

var ErrNoneNotSome = errors.New("attempted to get something from nothing")

func None[T any]() Option[T]                             { return none[T]{} }
func (n none[T]) Handle(ifsome func(t T), ifnone func()) { ifnone() }
func (n none[T]) IfSome(_ func(t T))                     {}
func (n none[T]) IfNone(f func())                        { f() }
func (n none[T]) IsSome() bool                           { return false }
func (n none[T]) Some() T                                { panic(ErrNoneNotSome) }
func (n none[T]) IsNone() bool                           { return true }
