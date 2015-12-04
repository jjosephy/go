package handler

import (
    //"github.com/dgrijalva/jwt-go"
    "fmt"
    "github.com/nmcclain/ldap"
    "interview/httperror"
    "io/ioutil"
    "net/http"
    "strings"
    //"time"
)

var (
    ldapServer string   = "nordstrom.net"
    ldapPort   uint16   = 3268
    baseDN     string   = "dc=*,dc=*"
    filter     string   = "(&(objectClass=user)(sAMAccountName=*)(memberOf=CN=*,OU=*,DC=*,DC=*))"
    Attributes []string = []string{"memberof"}
)


func TokenHandler(signingKey []byte) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        // TODO: validate the version
        switch r.Method {
            case "POST":
                b, err := ioutil.ReadAll(r.Body);
                if len(b) == 0 || err != nil {
                    httperror.NoRequestBody(w)
                    return
                }

                parts := strings.Split(string(b), "&")
                if len(parts) != 2 {
                    httperror.InvalidRequestBody(w)
                    return
                }

                uname := strings.Split(parts[0], "=")
                if len(uname) != 2 {
                    httperror.InvalidRequestBody(w)
                    return
                }

                pwd := strings.Split(parts[1], "=")
                if len(pwd) != 2 {
                    httperror.InvalidRequestBody(w)
                    return
                }

                // auth and create token
                l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))
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


            default:
                w.WriteHeader(http.StatusMethodNotAllowed)
            return
        }

        /*
        token := jwt.New(jwt.SigningMethodHS256)
        token.Claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
        tokenString, ex := token.SignedString(signingKey)

        if ex != nil {
            w.Write([]byte("Error occurred trying to create token"))
            w.WriteHeader(http.StatusUnauthorized)
            return
        }

        w.Write([]byte(tokenString))
        */
    }
}
