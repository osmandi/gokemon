# Gokemon

API Rest de Gokepones en Golang con el Framework Gin. Es un proyecto personal para poner en práctica BackEnd en Go.

Instalar dependencias y ejecutar
```
go get -u github.com/gin-gonic/gin
go run .
```

## Routes

```
localhost:8080        Método GET para obtener todos los pokemones 
localhost:8080/1      Método GET para obtener datos del pokemon con ID 1
localhost:8080/1      Método DELETE para elimnar el pokemón 1
localhost:8080/create Método POST para crear pokemon
```

> Se recomienda utilizar Postman para interactuar con el API.