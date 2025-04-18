package goprlib

type StateFilter int8

const (
	StateOpen StateFilter = iota
	StateClosed
	StateMerged
	StateAll
)
