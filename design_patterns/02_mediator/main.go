package main

import "fmt"

type Forum interface {
	Send(from, message string)
	SendFromTo(from, to, message string)
	Add(c Colleague)
}

type PoliticalForum struct {
	colleagues map[string]Colleague
}

func NewPoliticalForum() *PoliticalForum {
	return &PoliticalForum{make(map[string]Colleague)}
}

func (pf *PoliticalForum) Add(colleague Colleague) {
	pf.colleagues[colleague.GetName()] = colleague
	colleague.SetForum(pf);
}

func (pf *PoliticalForum) Send(sender, message string) {
	for name, colleague := range pf.colleagues {
		if name != sender {
			colleague.Receive(sender, message)
		}
	}
}

func (pf *PoliticalForum) SendFromTo(sender, receiver, message string) {
	colleague := pf.colleagues[receiver]
	if colleague != nil {
		colleague.Receive(sender, message)
	} else {
		fmt.Println(sender, "is not a member of this political forum!")
	}
}

type Colleague interface {
	GetName() string
	SetForum(forum Forum)
	Send(to, message string)
	SendAll(message string)
	Receive(from, message string)
}

type MemberOfParliament struct {
	name string
	forum Forum
}

func NewMemberOfParliament(name string) *MemberOfParliament {
	return &MemberOfParliament{name:name}
}
func (mop *MemberOfParliament) GetName() string {
	return mop.name
}
func (mop *MemberOfParliament) SetForum(forum Forum) {
	mop.forum = forum
}

func (mop *MemberOfParliament) Send(receiver, message string) {
	mop.forum.SendFromTo(mop.name, receiver, message);
}
func (mop *MemberOfParliament) SendAll(message string) {
	mop.forum.Send(mop.name, message);
}

func (mop *MemberOfParliament) Receive(sender, message string) {
	fmt.Printf("MP %v received a message from %v: %v\n",
		mop.name, sender, message)
}

type PrimeMinister struct {
	*MemberOfParliament
}
func NewPrimeMinister(name string) *PrimeMinister {
	return &PrimeMinister{NewMemberOfParliament(name)}
}

func (pm *PrimeMinister) Receive(sender, message string) {
	fmt.Printf("Prime Minister %v received a message from %v: %v\n",
	pm.name, sender, message)
}

func main(){
	forum := NewPoliticalForum()
	pm := NewPrimeMinister("Prime")
	mp1 := NewMemberOfParliament("MP1")
	mp2 := NewMemberOfParliament("MP2")
	forum.Add(pm)
	forum.Add(mp1)
	forum.Add(mp2)
	pm.SendAll("Hello everyone")
	mp1.Send("MP2", "Hello back")
	mp2.Send("Prime", "Dito")
	pm.Send("MP1", "Bullsh!t")
	mp1.SendAll("I second")
}