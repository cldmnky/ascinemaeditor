//go:build !(desktop || dev || production || bindings)

package main

func runWailsIfTagged() bool { return false }
