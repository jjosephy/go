package authentication

import (

)

type SimpleAuthProvider struct {
    SigningKey []byte
}

func (p *SimpleAuthProvider) AuthenticateUser(name string, pwd string) (string, error) {
    return GenerateToken(p.SigningKey)
}

func (p *SimpleAuthProvider) ValidateToken(token string) (bool, error) {
    return ValidateToken(token, p.SigningKey)
}
