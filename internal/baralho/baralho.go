package baralho

import (
	"fmt"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type Baralho []Carta

func NewBaralho() Baralho {
	naipes := []Naipe{Copas, Espadas, Ouros, Paus}
	valores := []Valor{Dois, Tres, Quatro, Cinco, Seis, Sete, Oito, Nove, Dez, Valete, Dama, Rei, As}
	baralho := make(Baralho, 0, 52)

	for _, naipe := range naipes {
		for _, valor := range valores {
			baralho = append(baralho, Carta{naipe: naipe, valor: valor})
		}
	}

	return baralho
}

func (b *Baralho) Embaralhar() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	r.Shuffle(len(*b), func(i, j int) {
		(*b)[i], (*b)[j] = (*b)[j], (*b)[i]
	})
}

func (b *Baralho) Sort() {
	ordemDosNaipes := map[Naipe]int{
		Copas:   0,
		Espadas: 1,
		Ouros:   2,
		Paus:    3,
	}
	ordemDosValores := map[Valor]int{
		Dois:   0,
		Tres:   1,
		Quatro: 2,
		Cinco:  3,
		Seis:   4,
		Sete:   5,
		Oito:   6,
		Nove:   7,
		Dez:    8,
		Valete: 9,
		Dama:   10,
		Rei:    11,
		As:     12,
	}

	sort.Slice(*b, func(i, j int) bool {
		if ordemDosNaipes[(*b)[i].naipe] != ordemDosNaipes[(*b)[j].naipe] {
			return ordemDosNaipes[(*b)[i].naipe] < ordemDosNaipes[(*b)[j].naipe]
		}
		return ordemDosValores[(*b)[i].valor] < ordemDosValores[(*b)[j].valor]
	})
}

func (b Baralho) String() string {
	var cartas []string
	for _, carta := range b {
		cartas = append(cartas, fmt.Sprintf("%s de %s", carta.valor, carta.naipe))
	}

	return strings.Join(cartas, "\n")
}

func (b *Baralho) PegarCartaAleatoria() (Carta, error) {
	if len(*b) == 0 {
		return Carta{}, fmt.Errorf("baralho vazio")
	}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	index := r.Intn(len(*b))

	carta := (*b)[index]

	*b = append((*b)[:index], (*b)[index+1:]...)

	return carta, nil
}
