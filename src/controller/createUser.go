package controller

import "github.com/gin-gonic/gin"

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
}
