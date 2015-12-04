package authentication

import (
    "github.com/dgrijalva/jwt-go"
    "time"
)

type AuthenticationProvider interface {
    AuthenticateUser(name string, pwd string) (string, error)
    ValidateToken(token string) (bool, error)
}

type SimpleAuthProvider struct {
    SigningKey []byte
}

func (p *SimpleAuthProvider) AuthenticateUser(name string, pwd string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)
    token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
    tokenString, ex := token.SignedString(p.SigningKey)

    if ex != nil {
        return "", ex
    }

    return tokenString, nil
}

func (p *SimpleAuthProvider) ValidateToken(token string) (bool, error) {
    return true, nil
}

/*
var (
    ldapServer string   = "nordstrom.net"
    ldapPort   uint16   = 389 //3268
    baseDN     string   = "dc=*,dc=*"
    filter     string   = "(&(objectClass=user)(sAMAccountName=*)(memberOf=CN=*,OU=*,DC=*,DC=*))"
    Attributes []string = []string{"memberof"}
)

// auth and create token
l, err := ldap.DialTLS("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort), nil)
if err != nil {
    httperror.AuthConnectToLDAPFailure(w)
    return
}
defer l.Close()
// l.Debug = true

dname := fmt.Sprint(uname[1], "@nordstrom.net")
err = l.Bind(dname, pwd[1])
if err != nil {
    httperror.Unauthorized(w)
    return
}

*/
