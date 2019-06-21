package main

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/gin-gonic/gin"
)

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

func deletePokemon(c *gin.Context) {

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	fmt.Println(reflect.TypeOf(idInt))

	for i, pokemon := range pokemons {
		if pokemon.PokedexNumber == idInt {
			pokemons = append(pokemons[:i], pokemons[i+1])
			c.JSON(200, gin.H{
				"msg": "Pokemon deleted",
			})
			return
		}
	}

	c.JSON(400, gin.H{
		"msg": "Not found",
	})
}

func levelUpPokemon(c *gin.Context) {

	id := c.Param("id")
	idInt, _ := strconv.Atoi(id)
	fmt.Println(reflect.TypeOf(idInt))

	for i, pokemon := range pokemons {
		if pokemon.PokedexNumber == idInt {
			pokemon.Level += 20
			c.SecureJSON(http.StatusOK, pokemon)
			pokemons[i] = pokemon
			return
		}
	}

	c.JSON(400, gin.H{
		"msg": "not found",
	})
}
