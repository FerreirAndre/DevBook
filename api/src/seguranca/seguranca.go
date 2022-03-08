package seguranca

import "golang.org/x/crypto/bcrypt"

// recebe uma string(senha) e coloca uma hash
func Hash(senha string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(senha), bcrypt.DefaultCost)
}

// compara a senha com uma senha com hash e retorna se são iguais
func VerificarSenha(senhaComHash, senha string) error {
	return bcrypt.CompareHashAndPassword([]byte(senhaComHash), []byte(senha))
}
