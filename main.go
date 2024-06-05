package main

import (
	"encoding/base64"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

// Estrutura de Dados usando YAMl
type Secrets struct {
	Data map[string]string `yaml:"data"`
}

func main() {
	// Verificar se a entrada existe
	if len(os.Args) < 2 {
		log.Fatalf("Usage: %s <secret.yaml>", os.Args[0])
	}
	inputFile := os.Args[1]

	// Ler o Arquivo YAML
	yamlFile, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatalf("Error reading YAML file: %v", err)
	}

	// Executando Unmarshall do YAML
	var secrets Secrets
	err = yaml.Unmarshal(yamlFile, &secrets)
	if err != nil {
		log.Fatalf("Error unmarshaling YAML data: %v", err)
	}

	// Criando arquivo output.txt
	outputFile := "output.txt"
	out, err := os.Create(outputFile)
	if err != nil {
		log.Fatalf("Error creating output file: %v", err)
	}
	defer out.Close()

	// Iterar todos os dados e decodificar os base64
	for key, value := range secrets.Data {
		decodedValue, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			log.Printf("Error decoding base64 value for key %s: %v", key, err)
			continue
		}
		// Escrever em key=value no arquivo de saída
		_, err = fmt.Fprintf(out, "%s=%s\n", key, decodedValue)
		if err != nil {
			log.Fatalf("Error writing to output file: %v", err)
		}
	}

	log.Printf("Decoded values written to %s", outputFile)
}
