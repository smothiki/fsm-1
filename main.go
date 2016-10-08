package main

import (
	"fmt"

	"github.com/smothiki/fsm-1/fsm"
)

type Thing struct {
	State fsm.State

	// our machine cache

}

func initjob(subject fsm.Stater, goal fsm.State) bool {
	fmt.Println(subject.CurrentState())
	return true
}

func plan(subject fsm.Stater, goal fsm.State) bool {
	fmt.Println(subject.CurrentState())
	return true
}

func execute(subject fsm.Stater, goal fsm.State) bool {
	fmt.Println(subject.CurrentState())
	return true
}

func close(subject fsm.Stater, goal fsm.State) bool {
	fmt.Println(subject.CurrentState())
	return true
}

// Add methods to comply with the fsm.Stater interface
func (t Thing) CurrentState() fsm.State { return t.State }
func (t Thing) SetState(s fsm.State)    { t.State = s }

func main() {

	some_thing := Thing{State: "idle"} // Our subject
	fmt.Println(some_thing.CurrentState())

	// Establish some rules for our FSM
	var rules fsm.Ruleset
	rules = make(map[fsm.Transition][]fsm.Guard)
	rules.AddRule(fsm.T{"idle", "init"}, initjob)
	rules.AddRule(fsm.T{"init", "planned"}, plan)
	rules.AddRule(fsm.T{"planned", "executed"}, execute)

	fmt.Println(some_thing.CurrentState())

	machine := fsm.Machine{Rules: &rules, Subject: some_thing}
	machine.Transition("init")
}
