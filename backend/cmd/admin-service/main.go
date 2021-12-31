package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"rides/models"
	"strings"
)

// Creates an Organization via the command line.
// TODO: Make less stupid and more pragmatic
func createOrg() {
	org := models.Organization{}
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("----Create a New Organization----")
	fmt.Print("* ID: ")
	org.ID, _ = reader.ReadString('\n')
	org.ID = strings.TrimSuffix(org.ID, "\n")

	fmt.Print("* Name: ")
	org.Name, _ = reader.ReadString('\n')
	org.Name = strings.TrimSuffix(org.Name, "\n")

	fmt.Print("* Password: ")
	org.Password, _ = reader.ReadString('\n')
	org.Password = strings.TrimSuffix(org.Password, "\n")

	org.Events = []uint16{}
	org.Drivers = []uint16{}

	err := models.CreateOrg(org)
	if err != nil {
		log.Fatal(err)
	}
}

// Prints help for admin-service
func printHelp() {
	fmt.Println("usage: admin-service <option>")
	fmt.Println("\toptions:")
	fmt.Println("\t\t-org required\t Creates a new organization")
}

// Admin Service allows interaction with backend via command line.
func main() {
	createOrgPtr := flag.Bool("org", false, "create a new organization")
	flag.Parse()

	if !(*createOrgPtr) {
		printHelp()
	}

	if *createOrgPtr {
		createOrg()
	}
}
