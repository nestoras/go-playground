package main

import "fmt"

type Receiver interface {
	SetOn(bool)
	IsOn() bool
}

type Light struct {
	isOn bool
}

func (l *Light) SetOn(isOn bool) {
	l.isOn = isOn
}

func (l *Light) IsOn() bool {
	return l.isOn
}

//the interface Command defines the method Execute
type Command interface {
	Execute()
}

type OnCommand struct {
	receiver Receiver
}

func NewOnCommand(receiver Receiver) *OnCommand {
	return &OnCommand{receiver}
}

func (oc OnCommand) Execute() {
	oc.receiver.SetOn(true)
}

type RemoteControl struct {
	command Command
}

func (rc *RemoteControl) PressButton() {
	rc.command.Execute()
}

func main(){
	remote := RemoteControl{}
	light := Light{}
	lightOnCmd := NewOnCommand(&light)
	remote.command = lightOnCmd
	remote.PressButton()
	if !light.isOn {
		fmt.Println("light should be on")
	}
	fmt.Println("light is on")
}