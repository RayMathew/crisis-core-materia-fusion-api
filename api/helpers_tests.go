package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSetCache(t *testing.T) {
	app := newTestApplication()

	key := "testKey"
	value := "testValue"
	app.setCache(key, value)

	// Check if the value has been set in the cache
	cachedValue, exists := app.getCachedData(key)

	assert.True(t, exists, "Value should exist in the cache")
	assert.Equal(t, value, cachedValue, "Cached value should match the set value")
}

func TestGetCachedData(t *testing.T) {
	app := newTestApplication()

	key := "testKey"
	value := "testValue"
	app.setCache(key, value)

	cachedData, found := app.getCachedData(key)

	// Assert that the data exists
	assert.True(t, found, "Value should be found in the cache")
	assert.Equal(t, value, cachedData, "Cached value should match the set value")

	// Test getting a non-existent key
	_, found = app.getCachedData("nonExistentKey")
	assert.False(t, found, "Value should not be found for a non-existent key")
}

func TestCacheUpdate(t *testing.T) {
	app := newTestApplication()

	key := "testKey"
	initialValue := "initialValue"
	app.setCache(key, initialValue)

	// Update the key with a new value
	updatedValue := "updatedValue"
	app.setCache(key, updatedValue)

	// Check if the value has been updated in the cache
	cachedData, found := app.getCachedData(key)
	assert.True(t, found, "Value should be found in the cache")
	assert.Equal(t, updatedValue, cachedData, "Cached value should be updated")
}
