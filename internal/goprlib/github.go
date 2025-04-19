package goprlib

type StateFilter int8

const (
	StateOpen StateFilter = iota
	StateClosed
	StateMerged
	StateAll
)

func GetStateFromString(stringState string) StateFilter {
	var state StateFilter

	switch stringState {
	case "open":
		state = StateOpen
	case "closed":
		state = StateClosed
	case "merged":
		state = StateMerged
	case "all":
		state = StateAll
	default:
		state = StateOpen
	}

	return state
}

func (stateValue StateFilter) ToString() string {
	var stringState string

	switch stateValue {
	case StateAll:
		stringState = ""
	case StateClosed:
		stringState = "closed"
	case StateMerged:
		stringState = "merged"
	case StateOpen:
		stringState = "open"
	}

	return stringState
}
