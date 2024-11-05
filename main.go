package main

import (
	"fmt"

	"github.com/caio86/go-deck/internal/cartas"
)

func main() {
	baralho := cartas.NewBaralho()
	baralho.Embaralhar()

	carta, _ := baralho.PegarCartaAleatoria()

	fmt.Printf("Carta escolhida: %s de %s \n", carta.Valor, carta.Naipe)
	fmt.Println("\nRestanto no baralho:")
	fmt.Println(baralho)

	baralho.Sort()
	fmt.Println(baralho)
}
