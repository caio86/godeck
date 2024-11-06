package baralho

type Naipe string

type Valor string

const (
	Copas   Naipe = "Copas"
	Espadas Naipe = "Espadas"
	Ouros   Naipe = "Ouros"
	Paus    Naipe = "Paus"

	Dois   Valor = "2"
	Tres   Valor = "3"
	Quatro Valor = "4"
	Cinco  Valor = "5"
	Seis   Valor = "6"
	Sete   Valor = "7"
	Oito   Valor = "8"
	Nove   Valor = "9"
	Dez    Valor = "10"
	Valete Valor = "Valete"
	Dama   Valor = "Dama"
	Rei    Valor = "Rei"
	As     Valor = "As"
)

type Carta struct {
	naipe Naipe
	valor Valor
}

func NewCarta(naipe Naipe, valor Valor) (Carta, error) {
	carta := Carta{
		naipe,
		valor,
	}
	if err := carta.validate(); err != nil {
		return Carta{}, err
	}
	return carta, nil
}

func (c Carta) validate() error {
	// TODO: make validation
	return nil
}

func (c Carta) Naipe() Naipe {
	return c.naipe
}

func (c Carta) Valor() Valor {
	return c.valor
}
