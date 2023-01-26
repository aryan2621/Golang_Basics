package main

import (
	"fmt"
	e "lesson-8/example"
	"log"
)

var pf = fmt.Printf
var pl = fmt.Println

type Contact struct {
	FirstName string
	LastName  string
	Phone     string
}

type Business struct {
	Name    string
	Address string
	Contact Contact
}

func (b Business) Info() {
	fmt.Printf("Name: %s\n Address: %s\n Contact: %s\n %s\n %s\n", b.Name, b.Address, b.Contact.FirstName, b.Contact.LastName, b.Contact.Phone)
}

type TSp float64
type TBs float64
type ML float64

func TspToML(t TSp) ML {
	return ML(t * 4.929)
}
func TBsToML(t TBs) ML {
	return ML(t * 14.787)
}

func (tsp TSp) TspToML() ML {
	return ML(tsp * 4.929)
}
func (tbs TBs) TBsToML() ML {
	return ML(tbs * 14.787)
}

func main() {
	con1 := Contact{"John", "Doe", "1234567890"}
	biz1 := Business{"Acme", "123 Main St", con1}
	biz1.Info()

	ml1 := ML(TSp(1) * 4.929)
	fmt.Printf("3 TSPs = %.2f ML \n", ml1)

	ml2 := ML(TBs(1) * 14.787)
	fmt.Printf("3 TBs = %.2f ML \n", ml2)

	fmt.Println("2TSP + 4TSP=", TspToML(2)+TspToML(4))
	fmt.Println("2TB + 4TB=", TBsToML(2)+TBsToML(4))

	pf("2TSP + 4TSP=%.2f ML \n", TspToML(2)+TspToML(4))
	pf("3 TSP =%.2f ml\n", TspToML(3))
	pf("2 TBP =%.2f ml\n", TBsToML(2))
	pl()

	pf("3 TSP =%.2f ml\n", TSp(3).TspToML())
	pf("2 TBP =%.2f ml\n", TBs(2).TBsToML())

	date := e.Date{}
	err := date.SetDay(31)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetMonth(10)
	if err != nil {
		log.Fatal(err)
	}
	err = date.SetYear(2021)
	if err != nil {
		log.Fatal(err)
	}
	pl(date)
	pf("Day: %d Month: %d Year: %d", date.Day(), date.Month(), date.Year())
}
