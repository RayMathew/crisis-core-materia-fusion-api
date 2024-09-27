package main

import (
	"fmt"
	"net/http"
)

func (app *application) backgroundTask(r *http.Request, fn func() error) {
	app.wg.Add(1)

	go func() {
		defer app.wg.Done()

		defer func() {
			err := recover()
			if err != nil {
				app.reportServerError(r, fmt.Errorf("%s", err))
			}
		}()

		err := fn()
		if err != nil {
			app.reportServerError(r, err)
		}
	}()
}

func (app *application) getCachedData(key string) (interface{}, bool) {
	app.mu.Lock()
	defer app.mu.Unlock()

	data, found := app.cache[key]
	return data, found
}

func (app *application) setCache(key string, value interface{}) {
	app.mu.Lock()
	defer app.mu.Unlock()

	app.cache[key] = value
}
