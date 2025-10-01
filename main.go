package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/devatomicfull/meu-crud-go/src/controller/routes"
)

func main() {
	// Carrega variáveis de ambiente a partir do arquivo .env.
	// Exemplo: no arquivo .env pode existir TEST=123
	// O pacote godotenv lê esse valor e o insere nas variáveis de ambiente do sistema.
	err := godotenv.Load()
	if err != nil {
		// log.Fatal encerra o programa imediatamente caso ocorra um erro.
		// Aqui, se o arquivo .env não for encontrado ou der erro, a aplicação não inicia.
		log.Fatal("Erro ao carregar .env")
	}

	// Recupera a variável de ambiente "TEST" do sistema (definida no .env).
	// Os valores de variáveis de ambiente são sempre strings.
	test := os.Getenv("TEST")

	// Exibe no terminal o valor da variável "TEST", apenas para validação.
	fmt.Println("Teste:", test)

	// Cria uma nova instância do servidor Gin com configuração padrão.
	// gin.Default() já inclui logger e middleware de recuperação de erros.
	router := gin.Default()

	// Inicializa as rotas da aplicação, passando o RouterGroup do Gin.
	// Aqui é onde você conecta suas rotas definidas no pacote "routes".
	routes.InitRoutes(&router.RouterGroup)

	// Inicia o servidor HTTP na porta 8080.
	// Se houver falha ao iniciar, o programa será encerrado com log.Fatal().
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
