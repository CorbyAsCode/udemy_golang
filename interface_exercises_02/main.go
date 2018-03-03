package main

import (
	"fmt"
)

type customs interface {
	printEntry()
	greeting()
	action()
}

type holiday struct {
	date string
}

type person struct {
	country string
}

func (h holiday) greeting() {
	if h.date == "3/17" {
		fmt.Println("The luck of the Irish be with ya!")
	} else if h.date == "12/25" {
		fmt.Println("Merry Christmas!")
	} else if h.date == "7/4" {
		fmt.Println("Happy 4th of July!")
	}
}

func (h holiday) action() {
	if h.date == "3/17" {
		fmt.Println("Have a green parade.")
	} else if h.date == "12/25" {
		fmt.Println("Open presents")
	} else if h.date == "7/4" {
		fmt.Println("Go shoot some fireworks")
	}
}

func (h holiday) printEntry() {
	fmt.Printf("Today is %s\n", h.date)
}

func (p person) greeting() {
	if p.country == "Canada" {
		fmt.Println("Take off hoser!")
	} else if p.country == "Britain" {
		fmt.Println("Would you like some tea and crumpets?")
	} else if p.country == "Australia" {
		fmt.Println("G'day mate!")
	}
}

func (p person) action() {
	if p.country == "Canada" {
		fmt.Println("Go play ice hockey")
	} else if p.country == "Britain" {
		fmt.Println("Go play soccer")
	} else if p.country == "Australia" {
		fmt.Println("Get lost in the Outback")
	}
}

func (p person) printEntry() {
	fmt.Printf("I'm from %s\n", p.country)
}

func whatWeDo(c customs) {
	c.printEntry()
	c.greeting()
	c.action()
	fmt.Println()

}

func main() {
	c := holiday{date: "12/25"}
	j := holiday{date: "7/4"}
	p := holiday{date: "3/17"}
	ca := person{country: "Canada"}
	br := person{country: "Britain"}
	au := person{country: "Australia"}
	whatWeDo(c)
	whatWeDo(j)
	whatWeDo(p)
	whatWeDo(ca)
	whatWeDo(br)
	whatWeDo(au)
}
