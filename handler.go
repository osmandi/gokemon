package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

// TODO: Implementar Put, Delete
// TODO: Usar SQlite
// TODO: Hacer lo mismo en Gorilla, Echo y Beego

// Handler
func getPokemons(c *gin.Context) {

	// Will output  :   while(1);["lena","austin","foo"]
	c.SecureJSON(http.StatusOK, pokemons)

}

// Get a Pokemon
func getPokemon(c *gin.Context) {

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

func createPokemon(c *gin.Context) {
	var pokemon Pokemon

	c.BindJSON(&pokemon)

	pokemons = append(pokemons, pokemon)

	c.JSON(200, pokemon)
}
