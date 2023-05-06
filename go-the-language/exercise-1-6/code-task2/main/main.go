package main

import (
	"fmt"
	"walkintoabar/people"
)

type Sentence struct {
	Speaker   *people.Person
	Receivers []*people.Person
	Text      string
	Mood      people.Mood
}

func main() {
	bob := &people.Person{Name: "Bob", Mood: people.Neutral}
	alice := &people.Person{Name: "Alice", Mood: people.Neutral}
	bartender := &people.Person{Name: "Bartender", Mood: people.Neutral}
	everybody := []*people.Person{bob, alice, bartender}
	sentences := []Sentence{
		{
			Speaker:   bartender,
			Receivers: []*people.Person{alice, bob},
			Text:      "What brings you in tonight?",
			Mood:      people.Excited,
		},
		{
			Speaker:   bob,
			Receivers: []*people.Person{alice},
			Text:      "Will you marry me?",
			Mood:      people.Excited,
		},
		{
			Speaker:   alice,
			Receivers: []*people.Person{bob},
			Text:      "Yes! I will!",
			Mood:      people.Happy,
		},
		{
			Speaker:   bartender,
			Receivers: []*people.Person{bob, alice},
			Text:      "Congratulations to both of you! First round's on me.",
			Mood:      people.Grateful,
		},
	}

	fmt.Printf("Walk into a bar story:\n")
	fmt.Printf("%s and %s walked into a bar.\n", bob.Name, alice.Name)
	bartender.Mood = people.Curious
	alice.Mood = people.Relaxed
	bob.Mood = people.Excited
	bartender.Mood = people.Curious
	for _, receiver := range everybody {
		fmt.Printf("%s is %s.\n", receiver.Name, receiver.Mood)
	}

	for _, sentence := range sentences {
		speaker := sentence.Speaker
		receivers := sentence.Receivers
		text := sentence.Text
		mood := sentence.Mood

		receiverNames := make([]string, len(receivers))
		for i, receiver := range receivers {
			receiverNames[i] = receiver.Name
		}
		if len(receivers) > 0 {
			fmt.Printf("%s said to %s, \"%s\"\n", speaker.Name, receiverNames, text)
		} else {
			fmt.Printf("%s said, \"%s\"\n", speaker, text)
		}
		for _, receiver := range receivers {
			receiver.Mood = mood
			fmt.Printf("%s is %s.\n", receiver.Name, receiver.Mood)
		}
	}
	fmt.Printf("End\n")
}
