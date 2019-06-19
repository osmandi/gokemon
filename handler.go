package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: Traer datos de int en main para tener un solo Slide
// TODO: Implementar Post, Put, Delete
// TODO: Usar SQlite
// TODO: Hacer lo mismo en Gorilla, Echo y Beego

// Handler
func getPokemons(c *gin.Context) {
	var pokemons []Pokemon
	pokemon1 := Pokemon{1, "Pikachu", "", []Type{TypeDark}, 89}
	pokemon2 := Pokemon{2, "Pikachu2", "", []Type{TypeDark}, 89}
	pokemon3 := Pokemon{3, "Pikachu3", "", []Type{TypeDark}, 89}
	pokemon4 := Pokemon{4, "Pikachu4", "", []Type{TypeDark}, 89}
	pokemons = append(pokemons, pokemon1, pokemon2, pokemon3, pokemon4)

	// Will output  :   while(1);["lena","austin","foo"]
	c.SecureJSON(http.StatusOK, pokemons)

}

// Get a Pokemon
func getPokemon(c *gin.Context) {

	var pokemons []Pokemon
	pokemon1 := Pokemon{1, "Pikachu", "", []Type{TypeDark}, 89}
	pokemon2 := Pokemon{2, "Pikachu2", "", []Type{TypeDark}, 89}
	pokemon3 := Pokemon{3, "Pikachu3", "", []Type{TypeDark}, 89}
	pokemon4 := Pokemon{4, "Pikachu4", "", []Type{TypeDark}, 89}
	pokemons = append(pokemons, pokemon1, pokemon2, pokemon3, pokemon4)

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	fmt.Println(reflect.TypeOf(idInt))

	for _, pokemon := range pokemons {
		if pokemon.PokedexNumber == idInt {
			c.SecureJSON(http.StatusOK, pokemon)
			return
		}
	}

	c.JSON(400, gin.H{
		"msg": "not found",
	})
}
