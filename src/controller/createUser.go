package controller

import (
	"fmt"
	"log"

	"github.com/devatomicfull/meu-crud-go/src/configuration/validation"
	"github.com/devatomicfull/meu-crud-go/src/controller/model/request"
	"github.com/devatomicfull/meu-crud-go/src/controller/model/response"
	"github.com/gin-gonic/gin"
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
		log.Printf("Error trying to marshal object, %s", err.Error())
		restErr := validation.ValidateUserError(err)

		c.JSON(restErr.Code, restErr)
		return
	}

	fmt.Println(userRequest)

	// Monta resposta simulada
	userResponse := response.UserResponse{
		ID:    "test",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}

	// Retorna resposta de sucesso
	c.JSON(201, userResponse)

}
