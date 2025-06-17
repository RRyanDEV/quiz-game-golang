package model

type Question struct {
	// Estrutura das Perguntas
	Text    string
	Options []string
	Answer  int
}

/*
Métodos de Estruturas devem ser definidos no mesmo pacote
*/

type GameState struct {
	// Estrutura do game
	Name      string
	Points    int
	Questions []Question
}
