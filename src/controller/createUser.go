package controller

import (
	"github.com/devatomicfull/meu-crud-go/src/configuration/rest_erro"
	"github.com/gin-gonic/gin"
	"log"
)

// CreateUser trata requisições POST para criar um novo usuário.
// @param c *gin.Context - contexto da requisição HTTP do Gin.
// Espera receber os dados do usuário no corpo (body) da requisição em formato JSON.
//
// Exemplo de JSON esperado no body:
//
//	{
//	  "name": "João Silva",
//	  "email": "joao@email.com",
//	  "password": "123456"
//	}
//
// Exemplo de uso na rota:
// POST /createUser
func CreateUser(c *gin.Context) {
	// Aqui será implementada a lógica para criar um novo usuário.
	// Exemplo:
	// var user User
	// if err := c.ShouldBindJSON(&user); err != nil {
	//     c.JSON(400, gin.H{"error": "Dados inválidos"})
	//     return
	// }
	// salvar no banco...
	// c.JSON(201, gin.H{"message": "Usuário criado com sucesso"})
	err := rest_erro.NewBadRequestError("Você chamou a rota de forma incorreta.")
	log.Printf("DEBUG: %+v\n", err)
	c.JSON(err.Code, err)
}
