package main

import (
	"bufio"
	"fmt"
	"os"

	model "github.com/RRyanDEV/quiz-game-golang/models"
)

func main() {
	game := &model.GameState{Points: 0}
	go game.ProccessCSV()
	game.Init()
	game.Run()

	fmt.Printf("Fim de Jogo!! Você fez %d pontos\n", game.Points)

	/* Quando rodar no terminar, ele solicita que o usuário feche manualmente*/
	fmt.Println("\nPressione a tecla ENTER para fechar o programa...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
