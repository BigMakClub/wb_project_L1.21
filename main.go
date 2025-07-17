package main

import "fmt"

type OldPrinter struct {
}

func (o *OldPrinter) PrintText(text string) {
	fmt.Println("Печать на старом принтере", text)
}

type ModerPrinter interface {
	Print(text string)
}

func usePrinter(m ModerPrinter) {
	m.Print("Мы смогли сокртаить расходы")
}

type Adapter struct {
	old *OldPrinter
}

func (a *Adapter) Print(text string) {
	a.old.PrintText(text)
}

func main() {

	old := &OldPrinter{}
	adapter := &Adapter{old}

	usePrinter(adapter)
}
