package main

import (
	"encoding/hex"
	"fmt"
)

func main() {
	Hello()
	// 1. Gerar as chaves RSA (se necessário)
	if err := GerarChaves(); err != nil {
		fmt.Println("Erro ao gerar chaves:", err)
		return
	}

	// 2. Criar uma blockchain com o bloco gênesis
	blockchain := Blockchain{
		Blocos: []Bloco{CriarBlocoGenesis()}, // Mude para CriarBlocoGenesis
	}

	// 3. Carregar a chave privada previamente gerada
	privKey, err := CarregarChavePrivada("private_key.pem") // Mude para CarregarChavePrivada
	if err != nil {
		fmt.Println("Erro ao carregar chave privada:", err)
		return
	}
	pubKey := &privKey.PublicKey

	// 4. Criar e assinar a primeira transação
	mensagem1 := "João enviou R$100 para Maria"
	assinatura1, err := Assinar(mensagem1, privKey) // Mude para Assinar
	if err != nil {
		fmt.Println("Erro ao assinar a mensagem:", err)
		return
	}
	transacao1 := Transacao{
		Remetente:    "ChavePublicaJoao",
		Destinatario: "ChavePublicaMaria",
		Valor:        100.0,
		Assinatura:   hex.EncodeToString(assinatura1),
	}

	// Validar a transação
	if err := VerificarAssinatura(mensagem1, assinatura1, pubKey); err != nil { // Mude para VerificarAssinatura
		fmt.Println("Erro ao validar a assinatura da transação:", err)
		return
	}
	fmt.Println("Transação 1 validada com sucesso!")

	// 5. Criar e assinar a segunda transação
	mensagem2 := "Maria enviou R$200 para Carlos"
	assinatura2, err := Assinar(mensagem2, privKey) // Mude para Assinar
	if err != nil {
		fmt.Println("Erro ao assinar a mensagem:", err)
		return
	}
	transacao2 := Transacao{
		Remetente:    "ChavePublicaMaria",
		Destinatario: "ChavePublicaCarlos",
		Valor:        200.0,
		Assinatura:   hex.EncodeToString(assinatura2),
	}

	// Validar a transação
	if err := VerificarAssinatura(mensagem2, assinatura2, pubKey); err != nil { // Mude para VerificarAssinatura
		fmt.Println("Erro ao validar a assinatura da transação:", err)
		return
	}
	fmt.Println("Transação 2 validada com sucesso!")

	// 6. Adicionar as transações validadas à blockchain
	blockchain.AdicionarBloco([]Transacao{transacao1, transacao2}) // Mude para AdicionarBloco

	// 7. Exibir a blockchain atualizada
	fmt.Println("Blockchain atual:")
	for _, bloco := range blockchain.Blocos {
		fmt.Printf("Índice: %d\n", bloco.Index)
		fmt.Printf("Timestamp: %s\n", bloco.Timestamp)
		fmt.Printf("Hash: %s\n", bloco.Hash)
		fmt.Printf("Hash Anterior: %s\n", bloco.HashAnterior)
		fmt.Printf("Transações: %v\n", bloco.Transacoes)
		fmt.Println()
	}

	// 8. Validar a blockchain
	if blockchain.ValidarBlockchain() { // Mude para ValidarBlockchain
		fmt.Println("Blockchain válida!")
	} else {
		fmt.Println("Blockchain inválida!")
	}
}
