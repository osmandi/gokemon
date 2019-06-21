package main

import (
	"github.com/gin-gonic/gin"
)

var pokemons []Pokemon

func init() {

	pokemon1 := Pokemon{1, "Pikachu", "", []Type{TypeDark}, 89}
	pokemon2 := Pokemon{2, "Pikachu2", "", []Type{TypeDark}, 89}
	pokemon3 := Pokemon{3, "Pikachu3", "", []Type{TypeDark}, 89}
	pokemon4 := Pokemon{4, "Pikachu4", "", []Type{TypeDark}, 89}

	pokemons = append(pokemons, pokemon1, pokemon2, pokemon3, pokemon4)
}

func main() {
	r := gin.Default()

	r.GET("/", getPokemons)
	r.GET("/:id", getPokemon)
	r.POST("/create", createPokemon)
	r.Run()

}
