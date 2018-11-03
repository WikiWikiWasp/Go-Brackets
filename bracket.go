package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	// argument handling
	// arguments:
	// 1 - # of players
	// 2 - bracket title
	args := os.Args[1:] // get all arguments after prog name

	initParticipants(initBracket(args))

	// TODO: Type of bracket (Round Robin, etc)

	// menu
	runMenu()

}

func initBracket(args []string) int {
	//fmt.Println(len(os.Args), os.Args)

	var numPpl int
	var err error
	if numPpl, err = strconv.Atoi(args[0]); err != nil { // convert first arg from str to int if valid arg
		panic(err)
	}

	for numPpl < 2 { // user input validation for correct number of participants
		fmt.Print("\nA minimum of 2 participants are required for the bracket.\nPlease enter valid number of participants: ")
		_, err := fmt.Scanf("%d", &numPpl) // get user input for var numPpl
		if err != nil {                    // validate if input is int
			panic(err)
		}
	}

	brTitle := strings.Join(args[1:], " ") // joins remaining arguments into bracket title

	// input confirmation
	fmt.Printf("\nCreating '%s' bracket with %d participants...\n\n", brTitle, numPpl)

	return numPpl
}

func initParticipants(numPpl int) {
	// entering participants names
	fmt.Println("Please enter participants' names: ")
	names := make([]string, numPpl)
	for i := 0; i < numPpl; i++ {
		reader := bufio.NewReader(os.Stdin)
		names[i], _ = reader.ReadString('\n')
	}

	fmt.Println("\nParticipants: ")
	for i, _ := range names {
		fmt.Print(names[i])
	}
}

func runMenu() {
	var userIn string
	for {
		fmt.Println("\n**************************************")
		fmt.Println("What would you like to do?\n")
		fmt.Println("(1) Add Participant\n(2) Delete Participant\n(3) Edit Participant\n(4) Record Results\n(5) Show Bracket\n(6) Quit")
		fmt.Println("**************************************\n")
		fmt.Scan(&userIn)

		switch userIn {
		case "1":
			fmt.Println("\nAdding Participant...")
			fmt.Print("Number of additional participants: ")
			var numParts int
			fmt.Scan(&numParts)
			addPart(numParts)
			// call func
		case "2":
			fmt.Println("\nDeleting Participant...")
			// call func
		case "3":
			fmt.Println("\nEditing Participant...")
			// call func
		case "4":
			fmt.Println("\nRecording Results...")
			// call func
		case "5":
			// call showBracket func
		case "6":
			fmt.Println("\nQuiting Bracket...")
			return
		default:
			fmt.Println("\nInvalid input. Please try again...")
		}
	}
}

func addPart(numParts int) {

}

func deletePart(numParts int) {

}

func editPart(numParts int) {

}

func recResults() {

}

func showBracket() {

}
