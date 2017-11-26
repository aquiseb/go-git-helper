package main

import (
	"flag"
	"flgs"
	log "logger"
)

// If user makes wrong usage of flags, we trigger ShowHelp()
func validateFlag(flag *flag.FlagSet) {
	flag.Usage = func() {
		flgs.ShowHelp()
	}
}

func main() {

	// Define add
	var add bool
	flag.BoolVar(&add, "a", false, "")
	flag.BoolVar(&add, "add", false, "")

	// Define commit
	var commit bool
	flag.BoolVar(&commit, "c", false, "")
	flag.BoolVar(&commit, "commit", false, "")
	qCommit := "⚡ What's your commit message?\n"
	rCommit := ""

	// Define status
	var status bool
	flag.BoolVar(&status, "s", false, "")
	flag.BoolVar(&status, "status", false, "")

	// Check if the current flag is valid
	validateFlag(flag.CommandLine)

	// Initialize the logger (to get date)
	log.InitLogger()

	// Parse the command-line flags
	flag.Parse()

	// If bool commit is true, execute git commit
	if add {
		flgs.Add()
	}

	// If bool commit is true, execute git commit
	if commit {
		flgs.Commit(qCommit, rCommit)
	}

	// If bool status is true, execute git status
	if status {
		flgs.Status()
	}

}
