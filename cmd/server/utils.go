package main

import (
	"errors"
	"io"
)

func (app *Application) render(page string, w io.Writer, data any) error {
	t, ok := app.cache[page]
	if !ok {
		return errors.New("Page not in cache")
	}

	err := t.Execute(w, data)
	if err != nil {
		return errors.New("Cannot execute template")
	}

	return nil
}
