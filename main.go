package main

import (
	"fmt"
	"os"

	"github.com/kikiluvv/clockr/config"
	"github.com/kikiluvv/clockr/export"
	"github.com/kikiluvv/clockr/gui"
	"github.com/kikiluvv/clockr/session"
)

func main() {
	if len(os.Args) < 2 {
		printUsage()
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "in":
		session.ClockIn()
	case "out":
		session.ClockOut()
	case "status":
		session.Status()
	case "export":
		export.ExportHours()
	case "break":
		if len(os.Args) < 3 {
			fmt.Println("Usage: clockr break [start|end]")
			return
		}
		switch os.Args[2] {
		case "start":
			session.BreakStart()
		case "end":
			session.BreakEnd()
		default:
			fmt.Println("Unknown break command:", os.Args[2])
			fmt.Println("Usage: clockr break [start|end]")
		}
	case "summary":
		session.ShowSummary() // you'll need to implement this in session package
	case "adjust":
		if len(os.Args) < 5 {
			fmt.Println("Usage: clockr adjust <date> in|out <HH:MM>")
			return
		}
		date := os.Args[2]
		field := os.Args[3]
		newTime := os.Args[4]
		session.AdjustSession(date, field, newTime)
	case "config":
		if len(os.Args) < 3 {
			fmt.Println("Usage: clockr config [set|show]")
			return
		}
		switch os.Args[2] {
		case "show":
			config.ShowConfig()
		case "set":
			if len(os.Args) < 5 {
				fmt.Println("Usage: clockr config set <field> <value>")
				return
			}
			field := os.Args[3]
			value := os.Args[4]
			config.SetConfig(field, value)
		default:
			fmt.Println("Unknown config command:", os.Args[2])
			fmt.Println("Usage: clockr config [set|show]")
		}
	case "gui":
		gui.Start()
	default:
		fmt.Println("Unknown command:", cmd)
		printUsage()
	}

}

func printUsage() {
	fmt.Println("Usage: clockr <command> [options]")
	fmt.Println("\nCommands:")
	fmt.Println("  gui                     		Open GUI")
	fmt.Println("  in                     		Clock in for work")
	fmt.Println("  out                    		Clock out from work")
	fmt.Println("  status                 		Show current session status")
	fmt.Println("  export                 		Export hours for the current pay period")
	fmt.Println("  break start            		Start a break")
	fmt.Println("  break end              		End the current break")
	fmt.Println("  summary                		Show quick summary of hours worked (day/week/pay period)")
	fmt.Println("  adjust <date> in|out <HH:MM>     	Adjust clock in/out time for a session")
	fmt.Println("  config show            		Show current config")
	fmt.Println("  config set <field> <value>       	Set a config field (pay_period/start/end)")
	fmt.Println("\nExamples:")
	fmt.Println("  clockr in")
	fmt.Println("  clockr break start")
	fmt.Println("  clockr adjust 2025-08-13 in 09:00")
	fmt.Println("  clockr config set pay_period biweekly")
}
