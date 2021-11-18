package modelos

import (
	"errors"
	"strings"
	"time"
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
func (usuario *Usuario) Preparar() error {
	usuario.formatar()
	if erro := usuario.validar(); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar() error {
	if usuario.Nome == "" {
		return errors.New("o nome é obrigatório, não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("o nick é obrigatório, não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("o email é obrigatório, não pode estar em branco")
	}
	if usuario.Senha == "" {
		return errors.New("o senha é obrigatória, não pode estar em branco")
	}

	return nil
}

func (usuario *Usuario) formatar() {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)
}
