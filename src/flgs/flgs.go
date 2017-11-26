package flgs

import (
	"fmt"
	"header"
	log "logger"
	"os/exec"
	"userinput"
)

func ShowHelp() {
	fmt.Println(`

Usage: ccli [OPTIONS]

Options:
	-a, --add		Add.
	-c, --commit    Add and commit.

	-h, --help      Print the help log.
	
`)
}

func Add() {

	// Add Command
	aCommand := exec.Command("git", "add", ".")

	// We use _ as we don't care about that value. If we need to use it, give it a name
	out, err := aCommand.Output()
	if err != nil {
		fmt.Printf(string(out))
		log.Fatal(err)
	}
	log.Println("Added")
}

func Commit(q string, r string) {

	// Tell QAR that user input is required
	// The question and answer strings are provided from parent
	required := true
	answer := userinput.QAR(q, r, required)

	// Show the fancy ASCII header
	header.Show()

	// Do git add
	Add()

	// Commit Command
	cCommand := exec.Command("git", "commit", "-m", answer)

	// We use _ as we don't care about that value. If we need to use it, give it a name
	out, err := cCommand.Output()
	fmt.Printf(string(out))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Committed::" + answer)
}

func Status() {
	sCommand := exec.Command("git", "status")
	out, err := sCommand.Output()
	fmt.Printf(string(out))
	if err != nil {
		log.Fatal(err)
	}
	// log.Printf("%s", out)
}
