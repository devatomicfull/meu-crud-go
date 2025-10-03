package controller

import (
	"fmt"
	"github.com/devatomicfull/meu-crud-go/src/configuration/rest_erro"
	"github.com/devatomicfull/meu-crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"log"
)

// CreateUser trata requisições POST para criar um novo usuário.
// @param c *gin.Context - contexto da requisição HTTP do Gin.
// Espera receber os dados do usuário no corpo (body) da requisição em formato JSON.
//
// Exemplo de JSON esperado no body:
//
//		{
//		  "name": "João Silva",
//		  "email": "joao@email.com",
//		  "password": "123456"
//	   "age": 12
//		}
//
// Exemplo de uso na rota:
// POST /createUser
func CreateUser(c *gin.Context) {
	var userRequest request.UserRequest

	// Faz o binding e validação do JSON
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_erro.NewBadRequestError(
			fmt.Sprintf("Existem alguns campos incorretos: %s", err.Error()))
		c.JSON(restErr.Code, restErr)
		return
	}

	log.Println(userRequest)
	fmt.Println(userRequest)

	// Sucesso: retorna usuário recebido
	c.JSON(201, gin.H{
		"message": "Usuário criado com sucesso",
		"user":    userRequest,
	})

}
