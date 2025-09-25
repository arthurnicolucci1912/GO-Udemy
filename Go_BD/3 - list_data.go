package main // Define o pacote como 'main', indicando que este é um programa executável.

import (
	"database/sql" // Importa a biblioteca padrão do Go para interagir com bancos de dados.
	"log"          // Importa a biblioteca para fazer logs (mensagens de erro, informações, etc.).

	// O '_' no início indica que o pacote é importado apenas por seus efeitos colaterais.
	// Ele registra o driver do SQLite para que a biblioteca 'database/sql' possa usá-lo.
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// A função 'main' é o ponto de entrada do programa.

	// Abre a conexão com o banco de dados. O primeiro parâmetro é o nome do driver ('sqlite3')
	// e o segundo é o caminho para o arquivo do banco de dados ('/test.db').
	// 'sql.Open' retorna um ponteiro para o banco de dados ('db') e um erro ('err').
	db, err := sql.Open("sqlite3", "./test.db")

	// Verifica se ocorreu um erro ao tentar abrir a conexão.
	// Se 'err' não for 'nil' (nulo), significa que houve um problema.
	if err != nil {
		// 'log.Fatal(err)' imprime a mensagem de erro e encerra a execução do programa.
		log.Fatal(err)
	}

	// 'defer' adia a execução de 'db.Close()' para o final da função 'main'.
	// Isso garante que a conexão com o banco de dados será fechada, mesmo que ocorra
	// um erro inesperado no meio do programa.
	defer db.Close()

	// Executa uma consulta SQL para selecionar todas as colunas 'id' e 'name' da tabela 'users'.
	// O resultado da consulta é armazenado na variável 'linhas' e um possível erro em 'err'.
	linhas, err := db.Query("SELECT id, name FROM users")
	if err != nil {
		log.Fatal(err) // Se a consulta falhar, o programa é encerrado.
	}

	// 'defer linhas.Close()' garante que os recursos associados à consulta 'linhas'
	// sejam liberados no final da função 'main'.
	defer linhas.Close()

	// Inicia um loop que percorre cada linha retornada pela consulta.
	// 'linhas.Next()' retorna 'true' se houver uma próxima linha e 'false' quando não houver mais.
	for linhas.Next() {
		// Declara duas variáveis para armazenar os valores de 'id' e 'name' da linha atual.
		var id int
		var name string

		// 'linhas.Scan()' copia os valores da linha atual para as variáveis 'id' e 'name'.
		// O '&' indica que estamos passando o endereço de memória das variáveis, permitindo
		// que 'Scan' altere seus valores.
		err = linhas.Scan(&id, &name)
		if err != nil {
			log.Fatal(err) // Se houver um erro ao ler a linha, o programa é encerrado.
		}

		// 'log.Printf()' imprime a saída formatada no console.
		// '%d' é um marcador de posição para um número inteiro (id) e '%s' para uma string (name).
		log.Printf("User: %d, %s", id, name)
	}

	// Após o loop terminar, esta linha verifica se ocorreu algum erro durante a iteração
	// (por exemplo, problemas de rede ou de conexão).
	if err = linhas.Err(); err != nil {
		log.Fatal(err) // Se um erro for encontrado, o programa é encerrado.
	}
}
