package main

import (
	"fmt"
	"strings"
	"io/ioutil"
    "log"
    "net/http"
    "os"
)

/* 
	Goal: playing with Go after going through the Go tour and other ressources by implementing
		a client that will consume an API that I have developed and deployed previously
		to represent the Tic-Tac-Toe Game.
*/

func main() {

	// tic-tac-toe board representation for output.
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	boardGuide := [][]string{
		[]string{"TL", "TM", "TR"},
		[]string{"ML", "M", "MR"},
		[]string{"BL", "BM", "BR"},
	}

	//the array we will use to represent the games main data structure
	boardArray := []string{"+","+","+","+","+","+","+","+","+"}

	//the parameter that is going to be sent to my API 
	var boardStringParameter string


	fmt.Println("Welcome to the unbeatable Tic-Tac-Toe by Sid Barrack!")
	fmt.Println("-----------------------------")
	fmt.Println("You play as X: ")
	fmt.Println("-----------------------------")

	//the game loop
	for {

	//here we use our helper functions to check if there is winner or a Tie
	var winnerX bool = checkWinnerX(boardArray)
	var winnerO bool = checkWinnerO(boardArray)
	var isBoardFull bool = checkIfBoardFull(boardArray)


	//if there is a winner or a tie we break out of the game loop thus ending the game
	if winnerX {
		fmt.Println("!----------------!")
		fmt.Println("You win! (impossible)")
		fmt.Println("!----------------!")
		break
		} else if winnerO {
			fmt.Println("----------------")
			fmt.Println("My API wins woo hoo!")
			fmt.Println("----------------")
			break
			} else if isBoardFull {
				fmt.Println("----------------")
				fmt.Println("this is a draw. I told you it was unbeatable!")
				fmt.Println("----------------")
				break
			}

	fmt.Println("choose a Position to play:")
	fmt.Println("-----------------------------")

	showBoard(boardGuide)

	//getting user input based on the boardGuide
	var position string
	fmt.Scanln(&position)

	//assigning the user's move on the board if it's valid and the position is available
	switch strings.ToUpper(position) {
	case "TL" :
		if boardArray[0] == "+" {
		boardArray[0] = "x"} else {
			fmt.Println("Position Already Taken! ",board)
			continue
		}
	case "TM" :
		if boardArray[1] == "+" {
		boardArray[1] = "x"} else {
			fmt.Println("Position Already Taken!")
			continue
		}
	case "TR" :
		if boardArray[2] == "+" {
		boardArray[2] = "x"} else {
			fmt.Println("Position Already Taken!")
			continue
		}
	case "ML" :
		if boardArray[3] == "+" {
		boardArray[3] = "x"} else {
			fmt.Println("Position Already Taken!")
			continue
		}
	case "M" :
		if boardArray[4] == "+" {
		boardArray[4] = "x"} else {
			fmt.Println("Position Already Taken!")
			continue
		}
	case "MR" :
		if boardArray[5] == "+" {
		boardArray[5] = "x"} else {
			fmt.Println("Position Already Taken!")
			continue
		}
	case "BL" :
		if boardArray[6] == "+" {
		boardArray[6] = "x"} else {
			fmt.Println("Position Already Taken!")
			continue
		}
	case "BM" :
		if boardArray[7] == "+" {
		boardArray[7] = "x"} else {
			fmt.Println("Position Already Taken!")
			continue
		}
	case "BR" :
		if boardArray[8] == "+" {
		boardArray[8] = "x"} else {
			fmt.Println("Position Already Taken!")
			continue
		}
	default :
		fmt.Println("Incorrect Position!")
		continue
		}

	

	//here we join our boardArray to make a string paramater that my API can understand
	boardStringParameter = strings.Join(boardArray,"")

	//and here we call the Tic-Tac-Toe Serverless API that I have deployed previously on Netlify
	//the API is a minimax algorithm implementation and returns the best possible move for O on the board
	var response, err = http.Get("https://hopeful-thompson-102d67.netlify.app/.netlify/functions/getBestOutcome?board="+boardStringParameter)

    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }

	//we read the response that is sent back from the API
    var responseData, error = ioutil.ReadAll(response.Body)
    if error != nil {
        log.Fatal(err)
    }


	//after getting the response data as a string we split it in our boardArray
	if len(string(responseData))==9 {
    boardArray = strings.Split(string(responseData),"")
	} 

	//assigning positions in the 2D board to be displayed
	boardArrayToBoard(boardArray,board)

	//replacing white spaces with + to be able to perform the next API request
	formatboardArrayForNextRequest(boardArray)

	//displaying the board
	showBoard(board)
	
	}

}


func showBoard(board [][]string){
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
		}
}

func formatboardArrayForNextRequest(board []string){
	for i, element := range board {
		if element == " "{
			board[i]="+"
		}
	}
}

func checkWinnerX(board []string) bool{
	if board[0]=="x" && board[1]=="x" && board[2]=="x" ||
	   board[3]=="x" && board[4]=="x" && board[5]=="x" ||
	   board[6]=="x" && board[7]=="x" && board[8]=="x" ||
	   board[0]=="x" && board[4]=="x" && board[8]=="x" ||
	   board[2]=="x" && board[4]=="x" && board[6]=="x" ||
	   board[0]=="x" && board[3]=="x" && board[6]=="x" ||
	   board[2]=="x" && board[5]=="x" && board[8]=="x" {
		return true
	   }
	   return false
}
func checkWinnerO(board []string) bool{
	if board[0]=="o" && board[1]=="o" && board[2]=="o" ||
	   board[3]=="o" && board[4]=="o" && board[5]=="o" ||
	   board[6]=="o" && board[7]=="o" && board[8]=="o" ||
	   board[0]=="o" && board[4]=="o" && board[8]=="o" ||
	   board[2]=="o" && board[4]=="o" && board[6]=="o" ||
	   board[0]=="o" && board[3]=="o" && board[6]=="o" ||
	   board[2]=="o" && board[5]=="o" && board[8]=="o" {
		return true
	   }
	   return false
}

func checkIfBoardFull(board []string) bool {
	for _,element := range board {
		if element != "x" && element != "o" {
			return false
		}
	}
	return true
}



func boardArrayToBoard(boardArray []string , board [][]string){
		for i, position := range boardArray {
			if position == " " {
				continue
			}
			switch i {
			case 0:
				board[0][0] = position
			case 1:
				board[0][1] = position
			case 2:
				board[0][2] = position
			case 3:
				board[1][0] = position
			case 4:
				board[1][1] = position
			case 5:
				board[1][2] = position
			case 6:
				board[2][0] = position
			case 7:
				board[2][1] = position
			case 8:
				board[2][2] = position
			}
		}
	}