package app

import "go.uber.org/dig"

var (
	c = dig.New()
)

func Provide(constructor interface{}, opts ...dig.ProvideOption) error {
	return c.Provide(constructor, opts...)
}

func Invoke(function interface{}, opts ...dig.InvokeOption) error {
	return c.Invoke(function, opts...)
}
