package main

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
