package main

import "fmt"

// Doctor reciver
type Doctor struct{}

func (d *Doctor) treatEye() {
	fmt.Println("Doctor do treat eye")
}

func (d *Doctor) treatNose() {
	fmt.Println("Doctor do treat Nose")
}

func (d *Doctor) treatMouth() {
	fmt.Println("Doctor do treat mouth")
}

// Abstract command
type Command interface {
	Treat()
}

// Treat eye order
type CommandTreatEye struct {
	doctor *Doctor
}

func (cmd *CommandTreatEye) Treat() {
	cmd.doctor.treatEye()
}

// Treat nose order
type CommandTreatNose struct {
	doctor *Doctor
}

func (cmd *CommandTreatNose) Treat() {
	cmd.doctor.treatNose()
}

// Treat Mouth
type CommandTreatMouth struct {
	doctor *Doctor
}

func (cmd *CommandTreatMouth) Treat() {
	cmd.doctor.treatMouth()
}

// Nurse invoker
type Nurse struct {
	CmdList []Command // Collection  command
}

// Send sick orders
func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}
	for _, cmd := range n.CmdList {
		cmd.Treat() // execte sick order bind command (here, the diagnosis method of the doctor whose medical record has been bound will be called)
	}

}

func main() {
	// A sick man example
	doctor := new(Doctor)
	// treat eye
	cmdEye := CommandTreatEye{doctor}
	// treat nose
	cmdNose := CommandTreatNose{doctor}

	cmdMouth := CommandTreatMouth{doctor}

	// Nurse
	nurse := new(Nurse)
	nurse.CmdList = append(nurse.CmdList, &cmdEye)
	nurse.CmdList = append(nurse.CmdList, &cmdNose)
	nurse.CmdList = append(nurse.CmdList, &cmdMouth)

	// Notify Execute
	nurse.Notify()

	fmt.Println("Main func done!")
}
