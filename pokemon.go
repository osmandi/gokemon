package main

type Type string

const (
	TypeGrass    Type = "grass"
	TypeDark     Type = "dark"
	TypePsychic  Type = "psychic"
	TypeFighting Type = "fighting"
	TypeIce      Type = "ice"
	TypeGround   Type = "ground"
	TypeElectric Type = "electric"
	TypeBug      Type = "bug"
	TypeWater    Type = "water"
	TypeRock     Type = "rock"
	TypeFairy    Type = "fairy"
	TypeFlying   Type = "flying"
	TypeFire     Type = "fire"
	TypePoison   Type = "poison"
	TypeNormal   Type = "normal"
	TypeGhost    Type = "ghost"
	TypeDragon   Type = "dragon"
	TypeSteel    Type = "steel"
)

type Pokemon struct {
	PokedexNumber int    `json:"pokedex_number"`
	Name          string `json:"name"`
	EvolvesTo     string `json:"evolves_to,omitempty"`
	Types         []Type `json:"types"`
	Level         int    `json:"level"`
}
