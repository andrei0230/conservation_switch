package main

import (
	"flag"
	"fmt"
	"os"
)

func turnOn(fileName string) {
	os.WriteFile(fileName, []byte("1"), 0777)
	fmt.Println("Conservation mode on.")
}

func turnOff(fileName string) {
	os.WriteFile(fileName, []byte("0"), 0777)
	fmt.Println("Conservation mode off.")
}

func main() {
	const file string = "/sys/bus/platform/drivers/ideapad_acpi/VPC2004:00/conservation_mode"

	var On bool
	var Off bool

	flag.BoolVar(&On, "on", false, "turns on conservation mode")
	flag.BoolVar(&Off, "off", false, "turns on conservation mode")

	flag.Parse()

	if On {
		turnOn(file)
	}

	if Off {
		turnOff(file)
	}
}
