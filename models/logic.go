package model

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func (g *GameState) Init() {
	// Inicia o jogo, coletando informações do Usuário
	fmt.Println("Seja bem vindo(a) ao quiz sobre Golang!!")
	fmt.Print("Escreva seu nome: ")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')

	if err != nil {
		panic("Erro ao ler a string")
	}
	g.Name = name
	fmt.Printf("Vamos ao jogo %s", g.Name)
}

func (g *GameState) ProccessCSV() {
	// Função que faz a leitura das informações do arquivo .CSV
	f, err := os.Open("quizgo.csv")
	if err != nil {
		panic("Erro ao ler o arquivo")
	}
	defer f.Close()
	/*DEFER: Tudo que passado à ele, será executado no final da função.*/

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		panic("Erro ao ler csv")
	}

	for index, record := range records {
		if index > 0 {
			question := Question{
				Text:    record[0],
				Options: record[1:5],
				Answer:  toInt(record[5]),
			}

			g.Questions = append(g.Questions, question)
		}
	}
}

func (g *GameState) Run() {
	// Executa o jogo
	for index, question := range g.Questions {
		// Mostra as perguntas ao usuário.
		fmt.Printf("%d. %s\n", index+1, question.Text)

		for j, option := range question.Options {
			// Mostra as Opções ao usuário.
			fmt.Printf("[%d] %s\n", j+1, option)
		}
		fmt.Print("Digite uma Alternative: ")

		var answer int
		var err error
		reader := bufio.NewReader(os.Stdin)

		for {
			// Lê a entrada do usuário
			read, _ := reader.ReadString('\n')

			// Limpa a string de entrada de todos os espaços em branco
			// (espaços, tabs, \n, \r) no início e no fim.
			// Isso torna o código robusto e compatível com Windows, Linux e macOS.
			cleanedInput := strings.TrimSpace(read)

			answer, err = strconv.Atoi(cleanedInput) // Usa a função padrão `strconv.Atoi` para a conversão.
			if err != nil {
				// Se a conversão falhar, informa o usuário e tenta novamente.
				fmt.Println("Entrada inválida. Por favor, digite apenas um número inteiro.")
				fmt.Print("Tente novamente: ")
				continue // Roda o For novamente
			}
			// Se a conversão for bem-sucedida, sai do loop.
			break
		}
		if answer == question.Answer {
			fmt.Println("Parabéns está correto!!")
			g.Points += 10
		} else {
			fmt.Println("Ops! Está errado!")
			fmt.Println("------------------------------")
		}
	}
}

func toInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}
