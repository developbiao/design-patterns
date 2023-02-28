package main

import "fmt"

type Cooker struct{}

func (c *Cooker) MakeChiken() {
	fmt.Println("Skewer Make chiken")
}

func (c *Cooker) MakeChuaner() {
	fmt.Println("Skewer Make Chuaner")
}

type Command interface {
	Make()
}

type CommandCookChiken struct {
	cooker *Cooker
}

func (cmd *CommandCookChiken) Make() {
	cmd.cooker.MakeChiken()
}

type CommandCookChuaner struct {
	cooker *Cooker
}

func (cmd *CommandCookChuaner) Make() {
	cmd.cooker.MakeChuaner()
}

type WaiterGirl struct {
	CmdList []Command
}

func (w *WaiterGirl) Notify() {
	if w.CmdList == nil {
		return
	}
	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}

func main() {
	cooker := new(Cooker)
	cmdChuaner := CommandCookChuaner{cooker}
	cmdChiken := CommandCookChiken{cooker}

	// Waiter
	waiterGirl := new(WaiterGirl)
	waiterGirl.CmdList = append(waiterGirl.CmdList, &cmdChuaner)
	waiterGirl.CmdList = append(waiterGirl.CmdList, &cmdChiken)

	// Notify
	waiterGirl.Notify()

	fmt.Println("Main func done!")
}
