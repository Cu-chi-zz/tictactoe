package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var jeu = [3][3]string{{"0", "0", "0"}, {"0", "0", "0"}, {"0", "0", "0"}}

func main() {
	displayGame()
	tourCounter := 1
	for tourCounter <= 9 {
		scanner := bufio.NewScanner(os.Stdin)
		if tourCounter == 1 {
			fmt.Println("Joueur 1:")
			scanner.Scan()
			placeChoose, err := strconv.Atoi(scanner.Text())
			if err != nil || placeChoose > 9 {
				fmt.Println("Merci d'entrer un nombre valide")
				continue
			}

			if checkIfAlreadyPlayed(placeChoose) {
				fmt.Println("Cette case est déjà jouée")
				continue
			}
			place(1, placeChoose)
			tourCounter++
		} else if tourCounter == 2 {
			fmt.Println("Joueur 2:")
			scanner.Scan()
			placeChoose, err := strconv.Atoi(scanner.Text())
			if err != nil || placeChoose > 9 {
				fmt.Println("Merci d'entrer un nombre valide")
				continue
			}

			if checkIfAlreadyPlayed(placeChoose) {
				fmt.Println("Cette case est déjà jouée")
				continue
			}
			place(2, placeChoose)
			tourCounter++
		} else if tourCounter == 3 {
			fmt.Println("Joueur 1:")
			scanner.Scan()
			placeChoose, err := strconv.Atoi(scanner.Text())
			if err != nil || placeChoose > 9 {
				fmt.Println("Merci d'entrer un nombre valide")
				continue
			}

			if checkIfAlreadyPlayed(placeChoose) {
				fmt.Println("Cette case est déjà jouée")
				continue
			}
			place(1, placeChoose)
			tourCounter++
		} else if tourCounter == 4 {
			fmt.Println("Joueur 2:")
			scanner.Scan()
			placeChoose, err := strconv.Atoi(scanner.Text())
			if err != nil || placeChoose > 9 {
				fmt.Println("Merci d'entrer un nombre valide")
				continue
			}

			if checkIfAlreadyPlayed(placeChoose) {
				fmt.Println("Cette case est déjà jouée")
				continue
			}
			place(2, placeChoose)
			tourCounter++
		} else if tourCounter == 5 {
			fmt.Println("Joueur 1:")
			scanner.Scan()
			placeChoose, err := strconv.Atoi(scanner.Text())
			if err != nil || placeChoose > 9 {
				fmt.Println("Merci d'entrer un nombre valide")
				continue
			}

			if checkIfAlreadyPlayed(placeChoose) {
				fmt.Println("Cette case est déjà jouée")
				continue
			}
			place(1, placeChoose)
			tourCounter++
		} else if tourCounter == 6 {
			fmt.Println("Joueur 2:")
			scanner.Scan()
			placeChoose, err := strconv.Atoi(scanner.Text())
			if err != nil || placeChoose > 9 {
				fmt.Println("Merci d'entrer un nombre valide")
				continue
			}

			if checkIfAlreadyPlayed(placeChoose) {
				fmt.Println("Cette case est déjà jouée")
				continue
			}
			place(2, placeChoose)
			tourCounter++
		} else if tourCounter == 7 {
			fmt.Println("Joueur 1:")
			scanner.Scan()
			placeChoose, err := strconv.Atoi(scanner.Text())
			if err != nil || placeChoose > 9 {
				fmt.Println("Merci d'entrer un nombre valide")
				continue
			}

			if checkIfAlreadyPlayed(placeChoose) {
				fmt.Println("Cette case est déjà jouée")
				continue
			}
			place(1, placeChoose)
			tourCounter++
		} else if tourCounter == 8 {
			fmt.Println("Joueur 2:")
			scanner.Scan()
			placeChoose, err := strconv.Atoi(scanner.Text())
			if err != nil || placeChoose > 9 {
				fmt.Println("Merci d'entrer un nombre valide")
				continue
			}

			if checkIfAlreadyPlayed(placeChoose) {
				fmt.Println("Cette case est déjà jouée")
				continue
			}
			place(2, placeChoose)
			tourCounter++
		} else if tourCounter == 9 {
			fmt.Println("Joueur 1:")
			scanner.Scan()
			placeChoose, err := strconv.Atoi(scanner.Text())
			if err != nil || placeChoose > 9 {
				fmt.Println("Merci d'entrer un nombre valide")
				continue
			}

			if checkIfAlreadyPlayed(placeChoose) {
				fmt.Println("Cette case est déjà jouée")
				continue
			}
			place(1, placeChoose)
			tourCounter++
		}

		displayGame()

		winner := winCheck()
		if winner != 0 {
			if winner == 1 {
				fmt.Println("Le joueur 1 a gagné !")
			} else if winner == 2 {
				fmt.Println("Le joueur 2 a gagné !")
			}

			break
		} else if winner == 0 && tourCounter == 10 {
			fmt.Println("Personne n'a gagné !")
		}
	}
}

func displayGame() {
	fmt.Println(jeu[0][0], jeu[0][1], jeu[0][2], "\n"+jeu[1][0], jeu[1][1], jeu[1][2], "\n"+jeu[2][0], jeu[2][1], jeu[2][2])
}

func place(fromPlayer int, position int) {
	if fromPlayer == 1 {
		if position == 1 {
			jeu[0][0] = "X"
		} else if position == 2 {
			jeu[0][1] = "X"
		} else if position == 3 {
			jeu[0][2] = "X"
		} else if position == 4 {
			jeu[1][0] = "X"
		} else if position == 5 {
			jeu[1][1] = "X"
		} else if position == 6 {
			jeu[1][2] = "X"
		} else if position == 7 {
			jeu[2][0] = "X"
		} else if position == 8 {
			jeu[2][1] = "X"
		} else if position == 9 {
			jeu[2][2] = "X"
		}
	} else {
		if position == 1 {
			jeu[0][0] = "O"
		} else if position == 2 {
			jeu[0][1] = "O"
		} else if position == 3 {
			jeu[0][2] = "O"
		} else if position == 4 {
			jeu[1][0] = "O"
		} else if position == 5 {
			jeu[1][1] = "O"
		} else if position == 6 {
			jeu[1][2] = "O"
		} else if position == 7 {
			jeu[2][0] = "O"
		} else if position == 8 {
			jeu[2][1] = "O"
		} else if position == 9 {
			jeu[2][2] = "O"
		}
	}
}

func checkIfAlreadyPlayed(position int) bool {
	if position == 1 {
		if jeu[0][0] != "0" {
			return true
		}
	} else if position == 2 {
		if jeu[0][1] != "0" {
			return true
		}
	} else if position == 3 {
		if jeu[0][2] != "0" {
			return true
		}
	} else if position == 4 {
		if jeu[1][0] != "0" {
			return true
		}
	} else if position == 5 {
		if jeu[1][1] != "0" {
			return true
		}
	} else if position == 6 {
		if jeu[1][2] != "0" {
			return true
		}
	} else if position == 7 {
		if jeu[2][0] != "0" {
			return true
		}
	} else if position == 8 {
		if jeu[2][1] != "0" {
			return true
		}
	} else if position == 9 {
		if jeu[2][2] != "0" {
			return true
		}
	}
	return false
}

func winCheck() int {
	/*
		X X X
		0 0 0
		0 0 0
	*/
	if jeu[0][0] == "X" && jeu[0][1] == "X" && jeu[0][2] == "X" {
		return 1
	}
	if jeu[0][0] == "O" && jeu[0][1] == "O" && jeu[0][2] == "O" {
		return 2
	}

	/*
		0 0 0
		X X X
		0 0 0
	*/
	if jeu[1][0] == "X" && jeu[1][1] == "X" && jeu[1][2] == "X" {
		return 1
	}
	if jeu[1][0] == "O" && jeu[1][1] == "O" && jeu[1][2] == "O" {
		return 2
	}

	/*
		0 0 0
		0 0 0
		X X X
	*/
	if jeu[2][0] == "X" && jeu[2][1] == "X" && jeu[2][2] == "X" {
		return 1
	}
	if jeu[2][0] == "O" && jeu[2][1] == "O" && jeu[2][2] == "O" {
		return 2
	}

	/*
		X 0 0
		X 0 0
		X 0 0
	*/
	if jeu[0][0] == "X" && jeu[1][0] == "X" && jeu[2][0] == "X" {
		return 1
	}
	if jeu[0][0] == "O" && jeu[1][0] == "O" && jeu[2][0] == "O" {
		return 2
	}

	/*
		0 X 0
		0 X 0
		0 X 0
	*/
	if jeu[0][1] == "X" && jeu[1][1] == "X" && jeu[2][1] == "X" {
		return 1
	}
	if jeu[0][1] == "O" && jeu[1][1] == "O" && jeu[2][1] == "O" {
		return 2
	}

	/*
		0 0 X
		0 0 X
		0 0 X
	*/
	if jeu[0][2] == "X" && jeu[1][2] == "X" && jeu[2][2] == "X" {
		return 1
	}
	if jeu[0][2] == "O" && jeu[1][2] == "O" && jeu[2][2] == "O" {
		return 2
	}

	/*
		X 0 0
		0 X 0
		0 0 X
	*/
	if jeu[0][0] == "X" && jeu[1][1] == "X" && jeu[2][2] == "X" {
		return 1
	}
	if jeu[0][0] == "O" && jeu[1][1] == "O" && jeu[2][2] == "O" {
		return 2
	}

	/*
		0 0 X
		0 X 0
		X 0 0
	*/
	if jeu[0][2] == "X" && jeu[1][1] == "X" && jeu[2][0] == "X" {
		return 1
	}
	if jeu[0][2] == "O" && jeu[1][1] == "O" && jeu[2][0] == "O" {
		return 2
	}
	return 0
}
