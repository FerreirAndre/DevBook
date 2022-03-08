package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"criadoEm,omitempty"`
}

//Preparar chama os metodos para verificar se os dados estão no formato correto para a API
func (usuario *Usuario) Preparar(etapa string) error {
	usuario.formatar(etapa)
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

// verifica o formato dos dados inseridos
func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório, não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório, não pode estar em branco")
	}

	if usuario.Email == "" {
		return errors.New("o email é obrigatório, não pode estar em branco")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("o formato do email é invalido")
	}

	if usuario.Senha == "" && etapa == "cadastro" {
		return errors.New("o senha é obrigatória, não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usuario.Senha)
		if erro != nil {
			return erro
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}
