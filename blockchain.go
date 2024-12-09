package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

type Transacao struct {
	Remetente    string
	Destinatario string
	Valor        float64
	Assinatura   string
}

type Bloco struct {
	Index        int
	Timestamp    string
	Transacoes   []Transacao
	HashAnterior string
	Hash         string
}

type Blockchain struct {
	Blocos []Bloco
}

func calcularHash(bloco Bloco) string {
	dados := fmt.Sprintf("%d%s%v%s", bloco.Index, bloco.Timestamp, bloco.Transacoes, bloco.HashAnterior)
	hash := sha256.Sum256([]byte(dados))
	return hex.EncodeToString(hash[:])
}

func CriarBlocoGenesis() Bloco {
	blocoGenesis := Bloco{
		Index:        0,
		Timestamp:    time.Now().String(),
		Transacoes:   []Transacao{},
		HashAnterior: "0",
	}
	blocoGenesis.Hash = calcularHash(blocoGenesis)
	return blocoGenesis
}

func (bc *Blockchain) AdicionarBloco(transacoes []Transacao) {
	ultimoBloco := bc.Blocos[len(bc.Blocos)-1]
	novoBloco := Bloco{
		Index:        ultimoBloco.Index + 1,
		Timestamp:    time.Now().String(),
		Transacoes:   transacoes,
		HashAnterior: ultimoBloco.Hash,
	}
	novoBloco.Hash = calcularHash(novoBloco)
	bc.Blocos = append(bc.Blocos, novoBloco)
}

func (bc *Blockchain) ValidarBlockchain() bool {
	for i := 1; i < len(bc.Blocos); i++ {
		blocoAtual := bc.Blocos[i]
		blocoAnterior := bc.Blocos[i-1]

		if blocoAtual.Hash != calcularHash(blocoAtual) {
			return false
		}

		if blocoAtual.HashAnterior != blocoAnterior.Hash {
			return false
		}
	}
	return true
}
