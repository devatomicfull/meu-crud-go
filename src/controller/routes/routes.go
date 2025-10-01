package routes

import (
	"github.com/devatomicfull/meu-crud-go/src/controller"
	"github.com/gin-gonic/gin"
)

// InitRoutes inicializa as rotas do módulo de usuários.
// Recebe um *gin.RouterGroup como parâmetro, que permite agrupar
// endpoints relacionados e aplicar middlewares compartilhados.
func InitRoutes(rg *gin.RouterGroup) {
	// GET /getUserById/:userId
	// Endpoint para buscar um usuário pelo ID.
	// ":userId" é um parâmetro de rota dinâmico que será extraído na handler.
	rg.GET("/getUserById/:userId", controller.FindUserById)

	// GET /getUserByEmail/:email
	// Endpoint para buscar um usuário pelo email.
	// ":email" é um parâmetro de rota dinâmico.
	rg.GET("/getUserByEmail/:email", controller.FindUserByEmail)

	// POST /createUser/:userId
	// Endpoint para criar um novo usuário.
	// Pode receber dados do usuário no corpo da requisição (JSON).
	rg.POST("/createUser/:userId", controller.CreateUser)

	// PUT /updateUser/:userId
	// Endpoint para atualizar um usuário existente.
	// ":userId" identifica o usuário a ser atualizado.
	rg.PUT("/updateUser/:userId", controller.UpdateUser)

	// DELETE /deleteUser/:userId
	// Endpoint para deletar um usuário pelo ID.
	rg.DELETE("/deleteUser/:userId", controller.DeleteUser)
}

/*
Notas técnicas:

1. *gin.RouterGroup
   - Permite agrupar rotas relacionadas.
   - Útil para aplicar middlewares ou prefixos de URL de forma centralizada.
   - Exemplo:
     router := gin.Default()
     apiGroup := router.Group("/api")
     routes.InitRoutes(apiGroup)

2. Parâmetros de rota (":param")
   - ":userId" ou ":email" são placeholders dinâmicos.
   - Gin os extrai automaticamente e você pode acessar via:
     c.Param("userId") ou c.Param("email") dentro do handler.

3. Estrutura de rotas
   - GET → obter dados.
   - POST → criar novos recursos.
   - PUT → atualizar recursos existentes.
   - DELETE → remover recursos.

4. Boas práticas
   - Idealmente, cada rota deve ter um handler (função) definido para processar requisições.
   - Handlers podem ser funções separadas ou métodos de um controller.
*/
