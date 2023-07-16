package main

import (
	"fmt"
	"net"
	"encoding/json"
)

type Pokemon struct {
	Name  string
	Type  string
	Level int
}

var Pokedex = []Pokemon{
	{"Bulbasaur", "Grass/Poison", 5},
	{"Ivysaur", "Grass/Poison", 16},
	{"Venusaur", "Grass/Poison", 32},
	{"Charmander", "Fire", 5},
	{"Charmeleon", "Fire", 16},
	{"Charizard", "Fire/Flying", 36},
	{"Squirtle", "Water", 5},
	{"Wartortle", "Water", 16},
	{"Blastoise", "Water", 36},
	{"Caterpie", "Bug", 3},
	{"Metapod", "Bug", 7},
	{"Butterfree", "Bug/Flying", 10},
	{"Weedle", "Bug/Poison", 3},
	{"Kakuna", "Bug/Poison", 7},
	{"Beedrill", "Bug/Poison", 10},
	{"Pidgey", "Normal/Flying", 3},
	{"Pidgeotto", "Normal/Flying", 18},
	{"Pidgeot", "Normal/Flying", 36},
	{"Rattata", "Normal", 3},
	{"Raticate", "Normal", 20},
	{"Spearow", "Normal/Flying", 3},
	{"Fearow", "Normal/Flying", 20},
	{"Ekans", "Poison", 10},
	{"Arbok", "Poison", 22},
	{"Pikachu", "Electric", 10},
	{"Raichu", "Electric", 28},
	{"Sandshrew", "Ground", 10},
	{"Sandslash", "Ground", 22},
	{"Nidoran♀", "Poison", 5},
	{"Nidorina", "Poison", 16},
	{"Nidoqueen", "Poison/Ground", 36},
	{"Nidoran♂", "Poison", 5},
	{"Nidorino", "Poison", 16},
	{"Nidoking", "Poison/Ground", 36},
	{"Clefairy", "Fairy", 10},
	{"Clefable", "Fairy", 28},
	{"Vulpix", "Fire", 10},
	{"Ninetales", "Fire", 28},
	{"Jigglypuff", "Normal/Fairy", 3},
	{"Wigglytuff", "Normal/Fairy", 20},
	{"Zubat", "Poison/Flying", 5},
	{"Golbat", "Poison/Flying", 22},
	{"Oddish", "Grass/Poison", 5},
	{"Gloom", "Grass/Poison", 21},
	{"Vileplume", "Grass/Poison", 36},
	{"Paras", "Bug/Grass", 5},
	{"Parasect", "Bug/Grass", 24},
	{"Venonat", "Bug/Poison", 10},
	{"Venomoth", "Bug/Poison", 31},
	{"Diglett", "Ground", 10},
	{"Dugtrio", "Ground", 26},
	{"Meowth", "Normal", 10},
	{"Persian", "Normal", 28},
	{"Psyduck", "Water", 10},
	{"Golduck", "Water", 33},
	{"Mankey", "Fighting", 10},
	{"Primeape", "Fighting", 28},
	{"Growlithe", "Fire", 10},
	{"Arcanine", "Fire", 34},
	{"Poliwag", "Water", 5},
	{"Poliwhirl", "Water", 25},
	{"Poliwrath", "Water/Fighting", 36},
	{"Abra", "Psychic", 8},
	{"Kadabra", "Psychic", 16},
	{"Alakazam", "Psychic", 36},
	{"Machop", "Fighting", 10},
	{"Machoke", "Fighting", 28},
	{"Machamp", "Fighting", 36},
	{"Bellsprout", "Grass/Poison", 5},
	{"Weepinbell", "Grass/Poison", 21},
	{"Victreebel", "Grass/Poison", 36},
	{"Tentacool", "Water/Poison", 5},
	{"Tentacruel", "Water/Poison", 30},
	{"Geodude", "Rock/Ground", 10},
	{"Graveler", "Rock/Ground", 25},
	{"Golem", "Rock/Ground", 36},
	{"Ponyta", "Fire", 16},
	{"Rapidash", "Fire", 40},
	{"Slowpoke", "Water/Psychic", 18},
	{"Slowbro", "Water/Psychic", 37},
	{"Magnemite", "Electric/Steel", 10},
	{"Magneton", "Electric/Steel", 30},
	{"Farfetch'd", "Normal/Flying", 36},
	{"Doduo", "Normal/Flying", 10},
	{"Dodrio", "Normal/Flying", 34},
	{"Seel", "Water", 22},
	{"Dewgong", "Water/Ice", 34},
	{"Grimer", "Poison", 10},
	{"Muk", "Poison", 28},
	{"Shellder", "Water", 10},
	{"Cloyster", "Water/Ice", 34},
	{"Gastly", "Ghost/Poison", 8},
	{"Haunter", "Ghost/Poison", 25},
	{"Gengar", "Ghost/Poison", 36},
	{"Onix", "Rock/Ground", 36},
	{"Drowzee", "Psychic", 12},
	{"Hypno", "Psychic", 34},
	{"Krabby", "Water", 10},
	{"Kingler", "Water", 33},
	{"Voltorb", "Electric", 19},
	{"Electrode", "Electric", 40},
	{"Exeggcute", "Grass/Psychic", 20},
	{"Exeggutor", "Grass/Psychic", 36},
	{"Cubone", "Ground", 15},
	{"Marowak", "Ground", 28},
	{"Hitmonlee", "Fighting", 30},
	{"Hitmonchan", "Fighting", 30},
	{"Lickitung", "Normal", 30},
	{"Koffing", "Poison", 10},
	{"Weezing", "Poison", 35},
	{"Rhyhorn", "Ground/Rock", 20},
	{"Rhydon", "Ground/Rock", 42},
	{"Chansey", "Normal", 30},
	{"Tangela", "Grass", 30},
	{"Kangaskhan", "Normal", 40},
	{"Horsea", "Water", 5},
	{"Seadra", "Water", 20},
	{"Goldeen", "Water", 5},
	{"Seaking", "Water", 30},
	{"Staryu", "Water", 15},
	{"Starmie", "Water/Psychic", 36},
	{"Mr. Mime", "Psychic/Fairy", 30},
	{"Scyther", "Bug/Flying", 30},
	{"Jynx", "Ice/Psychic", 30},
	{"Electabuzz", "Electric", 30},
	{"Magmar", "Fire", 30},
	{"Pinsir", "Bug", 30},
	{"Tauros", "Normal", 30},
	{"Magikarp", "Water", 5},
	{"Gyarados", "Water/Flying", 20},
	{"Lapras", "Water/Ice", 40},
	{"Ditto", "Normal", 30},
	{"Eevee", "Normal", 15},
	{"Vaporeon", "Water", 30},
	{"Jolteon", "Electric", 30},
	{"Flareon", "Fire", 30},
	{"Porygon", "Normal", 30},
	{"Omanyte", "Rock/Water", 40},
	{"Omastar", "Rock/Water", 40},
	{"Kabuto", "Rock/Water", 40},
	{"Kabutops", "Rock/Water", 40},
	{"Aerodactyl", "Rock/Flying", 40},
	{"Snorlax", "Normal", 50},
	{"Articuno", "Ice/Flying", 50},
	{"Zapdos", "Electric/Flying", 50},
	{"Moltres", "Fire/Flying", 50},
	{"Dratini", "Dragon", 20},
	{"Dragonair", "Dragon", 30},
	{"Dragonite", "Dragon/Flying", 55},
	{"Mewtwo", "Psychic", 70},
	{"Mew", "Psychic", 30},
}


func handleUDPConnection(conn *net.UDPConn, addr *net.UDPAddr, req []byte, n int) {
	// Converter o número recebido para um Pokémon correspondente
	number := int(req[0])<<24 | int(req[1])<<16 | int(req[2])<<8 | int(req[3])
	index := number % len(Pokedex)
	pokemon := Pokedex[index]

	// Serializar o Pokémon para JSON
	pokemonJSON, err := json.Marshal(pokemon)
	if err != nil {
		fmt.Println("Erro ao serializar o Pokémon:", err)
		return
	}

	// Envie a resposta para o cliente
	_, errResp := conn.WriteToUDP(pokemonJSON, addr)
	if errResp != nil {
		fmt.Println("Erro ao enviar resposta:", errResp)
		return
	}

}

func main() {
	// Define o endpont do servidor UDP
	addr, err := net.ResolveUDPAddr("udp", ":8080")
	if err != nil {
		fmt.Println("Erro ao resolver endereço")
		return
	}

	// Prepara o endpoint UDP para receber requests
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Erro ao iniciar o servidor")
		return
	}

	// Fecha conn
	defer conn.Close()

	fmt.Println("Servidor UDP aguardando conexões...")

	for {
		req := make([]byte, 4)

		// Recebe request (número do cliente)
		n, clientAddr, err := conn.ReadFromUDP(req)
		if err != nil {
			fmt.Println("Erro ao ler dados UDP")
			continue
		}

		go handleUDPConnection(conn, clientAddr, req, n)
	}
}
