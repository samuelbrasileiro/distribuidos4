package pokedexRPC

type PokedexRPC struct {}

func (p *PokedexRPC) GetPokemon(number int, reply *Pokemon) error {
	if number < 0 || number >= len(Pokedex) {
		return fmt.Errorf("Número inválido")
	}
	*reply = Pokedex[number]
	return nil
}