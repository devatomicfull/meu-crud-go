package controller

import "github.com/gin-gonic/gin"

// DeleteUser trata requisições DELETE para remover um usuário existente.
// @param c *gin.Context - contexto da requisição HTTP do Gin.
// O ID do usuário a ser removido deve ser fornecido como parâmetro de rota.
//
// Exemplo de rota:
// DELETE /deleteUser/:userId
//
// Exemplo de uso na rota:
// apiGroup.DELETE("/deleteUser/:userId", controller.DeleteUser)
func DeleteUser(c *gin.Context) {
	// Aqui será implementada a lógica para deletar um usuário existente.
	// Exemplo:
	// userId := c.Param("userId")
	// deletar no banco de dados...
	// c.JSON(200, gin.H{"message": "Usuário deletado com sucesso"})
}
