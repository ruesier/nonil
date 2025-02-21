package nonil

import "errors"

type Result[T any, E error] interface {
	Handle(isValid func(t T), isErr func(e E))
	IfValid(func(t T))
	IfErr(func(e E))

	IsValid() bool
	// Valid result, can panic
	Valid() T
	IsError() bool
	// Error result, can panic
	Error() E
}

type valid[T any, E error] struct {
	wrap T
}

var ErrValidNotError = errors.New("attempted to get error value from valid Result")

func Valid[T any, E error](wrap T) Result[T, E] {
	return valid[T, E]{
		wrap: wrap,
	}
}

func (v valid[T, E]) Handle(isValid func(t T), isErr func(e E)) { isValid(v.wrap) }
func (v valid[T, E]) IfValid(f func(t T))                       { f(v.wrap) }
func (v valid[T, E]) IfErr(_ func(e E))                         {}
func (v valid[T, E]) IsValid() bool                             { return true }
func (v valid[T, E]) Valid() T                                  { return v.wrap }
func (v valid[T, E]) IsError() bool                             { return false }
func (v valid[T, E]) Error() E                                  { panic(ErrValidNotError) }

type invalid[T any, E error] struct {
	wrap E
}

var ErrErrorNotValid = errors.New("attempted to get value from InValid Result")

func InValid[T any, E error](wrap E) Result[T, E] {
	return invalid[T, E]{
		wrap: wrap,
	}
}

func (iv invalid[T, E]) Handle(isValid func(t T), isErr func(e E)) { isErr(iv.wrap) }
func (iv invalid[T, E]) IfValid(_ func(t T))                       {}
func (iv invalid[T, E]) IfErr(f func(e E))                         {f(iv.wrap)}
func (iv invalid[T, E]) IsValid() bool                             { return false }
func (iv invalid[T, E]) Valid() T                                  { panic(ErrErrorNotValid) }
func (iv invalid[T, E]) IsError() bool                             { return true }
func (iv invalid[T, E]) Error() E                                  { return iv.wrap }

func ToResult[T any, E error](t T, e E) (r Result[T, E]) {
	defer func() {
		if rec := recover(); rec != nil {
			// assuming that the panic would be from e.Error() 
			// because e was nil
			// therefore, a nil error means this 
			// should return a valid result
			r = Valid[T, E](t)
		}
	}()
	if e.Error() != "" {
		return InValid[T, E](e)
	}
	return Valid[T, E](t)
}