package main

import "go.uber.org/fx"

func main() {
	app := fx.New(fx.Provide(), fx.Invoke())
	app.Run()
}
