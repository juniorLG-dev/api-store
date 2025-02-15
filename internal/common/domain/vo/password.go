package vo

import (
	"encoding/hex"
	"crypto/sha256"
)

type Password struct {
	Value string
}

func NewPassword(pass string) *Password {
	password := Password{Value: pass}
	password.Value = password.encryptPassword(pass)
	return &password
}

func (p *Password) CheckPassword(pass string) bool {
	return p.encryptPassword(pass) == p.Value
}

func (p *Password) encryptPassword(pass string) string {
	hasher := sha256.New()
	hasher.Write([]byte(pass))
	hashBytes := hasher.Sum(nil)
	hash := hex.EncodeToString(hashBytes)
	return hash
}