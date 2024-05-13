package services

import (
	"api_golang_ia/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type IaServices struct{}

func (ia IaServices) GetWords() []models.Word {
	// Obtenha a chave da API do OpenAI da variável de ambiente
	openaiAPIKey := os.Getenv("OPENAI_API_KEY_GOLANG")

	// Verifica se a chave foi definida
	if openaiAPIKey == "" {
		fmt.Println("A chave OPENAI_API_KEY_GOLANG não está definida.")
		return nil
	}

	// Define a URL da API da OpenAI
	apiURL := "https://api.openai.com/v1/chat/completions"

	// Mensagens de exemplo para treinar o modelo de linguagem
	configMessages := []map[string]string{
		{
			"role":    "system",
			"content": "Sua missão é retornar uma lista de palavras em inglês com quatro alternativas e uma correta, o formato retornado será assim: [{\"word\": \"cat\", \"translation\": \"gato\", \"options\": [\"cachorro\", \"rato\", \"pássaro\", \"peixe\"]}] somente o json e nenhum outro texto",
		},
		{
			"role":    "user",
			"content": "Retorne para mim o campo 'word' em inglês, o campo 'translation' em pt-br e o campo 'options' com quatro alternativas em pt-br",
		},
	}

	// Corpo da solicitação contendo a pergunta
	requestBody, err := json.Marshal(map[string]interface{}{
		"model": "gpt-3.5-turbo",
		"messages": configMessages,
	})
	if err != nil {
		fmt.Println("Erro ao converter mensagem para JSON::", err)
		return nil
	}

	// Criando uma solicitação POST para a API do OpenAI
	req, err := http.NewRequest("POST", apiURL, bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println("Erro ao criar a solicitação HTTP:", err)
		return nil
	}

	// Definindo o cabeçalho de autorização com a chave da API
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+openaiAPIKey)

	// Enviando a solicitação HTTP
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Erro ao enviar a solicitação HTTP:", err)
		return nil
	}
	defer resp.Body.Close()

	// Lendo a resposta da API
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Erro ao ler a resposta da API:", err)
		return nil
	}

	var result map[string]interface{}
	json.Unmarshal([]byte(responseBody), &result)

	var words []models.Word

	fmt.Println("=========================")
	fmt.Println(result)
	fmt.Println("=========================")

	// Processando a resposta da API
	if choices, ok := result["choices"].([]interface{}); ok && len(choices) > 0 {
		if firstChoice, ok := choices[0].(map[string]interface{}); ok {
			if message, ok := firstChoice["message"].(map[string]interface{}); ok {
				response := message["content"].(string)

				fmt.Println("=========================")
				fmt.Println(response)
				fmt.Println("=========================")

				err := json.Unmarshal([]byte(response), &words)
				if err != nil {
					fmt.Println("Erro ao fazer o Unmarshal do JSON:", err)
					return nil
				}
				return words
			}
		}
	}

	fmt.Println("Não foi possível processar a resposta da API.")
	return nil
}