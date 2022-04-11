package autenticacao

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// Recebe o Id do usuario e retorna um token com suas permissões.
func CriarToken(usuarioId uint64) (string, error) {
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioId
	// Secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("secret")) // Secret
}
