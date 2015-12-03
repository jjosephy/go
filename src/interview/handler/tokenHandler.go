package handler

import (
    //"github.com/dgrijalva/jwt-go"
    "fmt"
    "github.com/nmcclain/ldap"
    "interview/httperror"
    "io/ioutil"
    "net/http"
    //"time"
)

var (
    ldapServer string   = "nordstrom.net"
    ldapPort   uint16   = 3268
    baseDN     string   = "dc=*,dc=*"
    filter     string   = "(&(objectClass=user)(sAMAccountName=*)(memberOf=CN=*,OU=*,DC=*,DC=*))"
    Attributes []string = []string{"memberof"}
    user       string   = "q4vy@nordstrom.net"
    passwd     string   = "Walking*1out"
)


func TokenHandler(signingKey []byte) http.HandlerFunc {
    return func (w http.ResponseWriter, r *http.Request) {
        // TODO: validate the version
        switch r.Method {
            case "POST":
                //w.Write([]byte(r.Body))
                b, err := ioutil.ReadAll(r.Body);

                if err == nil {
                    s := fmt.Sprint(string(b), "\n")
                    w.Write([]byte(s))
                    return
                } else {
                    w.Write([]byte("Error strying to read body"))
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

                err = l.Bind(user, passwd)
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
