package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func CarregarChavePrivada(caminho string) (*rsa.PrivateKey, error) {
	privKeyFile, err := os.ReadFile(caminho)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo da chave privada: %v", err)
	}
	block, _ := pem.Decode(privKeyFile)
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, fmt.Errorf("arquivo não contém uma chave privada válida")
	}
	privKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("erro ao parsear chave privada: %v", err)
	}
	return privKey, nil
}

func CarregarChavePublica(caminho string) (*rsa.PublicKey, error) {
	pubKeyFile, err := os.ReadFile(caminho)
	if err != nil {
		return nil, fmt.Errorf("erro ao ler o arquivo da chave pública: %v", err)
	}
	block, _ := pem.Decode(pubKeyFile)
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, fmt.Errorf("arquivo não contém uma chave pública válida")
	}
	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, fmt.Errorf("erro ao parsear chave pública: %v", err)
	}
	return pubKey.(*rsa.PublicKey), nil
}

// Gera um par de chaves RSA (pública e privada) e as salva em arquivos
func GerarChaves() error {
	// Tamanho da chave (2048 bits)
	chavePrivada, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return fmt.Errorf("erro ao gerar chave privada: %v", err)
	}

	// Codificar a chave privada em PEM
	chavePrivadaPEM := &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(chavePrivada),
	}

	// Salvar a chave privada em um arquivo
	privFile, err := os.Create("private_key.pem")
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de chave privada: %v", err)
	}
	defer privFile.Close()

	err = pem.Encode(privFile, chavePrivadaPEM)
	if err != nil {
		return fmt.Errorf("erro ao salvar chave privada: %v", err)
	}

	// Gerar a chave pública
	chavePublica := &chavePrivada.PublicKey
	chavePublicaBytes, err := x509.MarshalPKIXPublicKey(chavePublica)
	if err != nil {
		return fmt.Errorf("erro ao gerar chave pública: %v", err)
	}

	// Codificar a chave pública em PEM
	chavePublicaPEM := &pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: chavePublicaBytes,
	}

	// Salvar a chave pública em um arquivo
	pubFile, err := os.Create("public_key.pem")
	if err != nil {
		return fmt.Errorf("erro ao criar arquivo de chave pública: %v", err)
	}
	defer pubFile.Close()

	err = pem.Encode(pubFile, chavePublicaPEM)
	if err != nil {
		return fmt.Errorf("erro ao salvar chave pública: %v", err)
	}

	fmt.Println("Chaves geradas e salvas com sucesso!")
	return nil
}

// Função para criptografar uma mensagem usando a chave pública
func Criptografar(mensagem string, chavePublica *rsa.PublicKey) ([]byte, error) {
	mensagemBytes := []byte(mensagem)
	criptografado, err := rsa.EncryptOAEP(sha256.New(), rand.Reader, chavePublica, mensagemBytes, nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao criptografar mensagem: %v", err)
	}
	return criptografado, nil
}

// Função para descriptografar uma mensagem usando a chave privada
func descriptografar(criptografado []byte, chavePrivada *rsa.PrivateKey) (string, error) {
	mensagemBytes, err := rsa.DecryptOAEP(sha256.New(), rand.Reader, chavePrivada, criptografado, nil)
	if err != nil {
		return "", fmt.Errorf("erro ao descriptografar mensagem: %v", err)
	}
	return string(mensagemBytes), nil
}

// Função para assinar digitalmente uma mensagem
func Assinar(mensagem string, chavePrivada *rsa.PrivateKey) ([]byte, error) {
	// Calcula o hash da mensagem
	hash := sha256.Sum256([]byte(mensagem))

	// Usa a constante crypto.SHA256 para assinar
	assinatura, err := rsa.SignPSS(rand.Reader, chavePrivada, crypto.SHA256, hash[:], nil)
	if err != nil {
		return nil, fmt.Errorf("erro ao assinar mensagem: %v", err)
	}
	return assinatura, nil
}

// Função para verificar uma assinatura
// Função para verificar uma assinatura
func VerificarAssinatura(mensagem string, assinatura []byte, chavePublica *rsa.PublicKey) error {
	// Calcula o hash da mensagem
	hash := sha256.Sum256([]byte(mensagem))

	// Usa a constante crypto.SHA256 para verificar
	err := rsa.VerifyPSS(chavePublica, crypto.SHA256, hash[:], assinatura, nil)
	if err != nil {
		return fmt.Errorf("assinatura inválida: %v", err)
	}
	return nil
}
