package controller

import "github.com/gin-gonic/gin"

// FindUserById trata requisições GET para buscar um usuário pelo ID.
// @param c *gin.Context - contexto da requisição HTTP do Gin.
// O ID do usuário deve ser fornecido como parâmetro de rota ("/getUserById/:userId").
func FindUserById(c *gin.Context) {
	// Aqui será implementada a lógica para buscar o usuário no banco
	// Exemplo:
	// userId := c.Param("userId")
	// Buscar no banco de dados...
}

// FindUserByEmail trata requisições GET para buscar um usuário pelo e-mail.
// @param c *gin.Context - contexto da requisição HTTP do Gin.
// O e-mail do usuário deve ser fornecido como parâmetro de rota ("/getUserByEmail/:email").
func FindUserByEmail(c *gin.Context) {
	// Aqui será implementada a lógica para buscar o usuário no banco
	// Exemplo:
	// email := c.Param("email")
	// Buscar no banco de dados...
}
