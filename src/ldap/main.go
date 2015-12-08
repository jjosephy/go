package main

import (
    //"crypto/x509"
    "fmt"
    "github.com/dgrijalva/jwt-go"
    "github.com/nmcclain/ldap"
    "log"
    //"golang.org/x/oauth2"
    //"time"
    //"io/ioutil"
    //"os"
    //"fmt"
    //"encoding/pem"
)

var (
	ldapServer string   = "nordstrom.net"
	ldapPort   uint16   = 389 //636 //389 //3268 //389 //636 //3268s

	baseDN     string   = "dc=*,dc=*"
	filter     string   = "(&(objectClass=user)(sAMAccountName=*)(memberOf=CN=*,OU=*,DC=*,DC=*))"
	Attributes []string = []string{"memberof"}
	user       string   = "q4vy@nordstrom.net"
	passwd     string   = "Walking*1out"
)


/*
func ExampleConfig() {
    conf := &oauth2.Config{
        ClientID:     "CLIENTID",
        ClientSecret: "CLIENTSECRET",
        Scopes:       []string{"non-expiring"},
        RedirectURL: "https://provider.com/redir",
        Endpoint: oauth2.Endpoint{
            AuthURL:  "http://localhost:8443/auth",
            TokenURL: "http://localhost:8443/token",
        },
    }

	// Redirect user to consent page to ask for permission
	// for the scopes specified above.
	url := conf.AuthCodeURL("state", oauth2.AccessTypeOffline)
	fmt.Printf("Visit the URL for the auth dialog: %v", url)

	// Use the authorization code that is pushed to the redirect URL.
	// NewTransportWithCode will do the handshake to retrieve
	// an access token and initiate a Transport that is
	// authorized and authenticated by the retrieved token.
	var code string
	if _, err := fmt.Scan(&code); err != nil {
		log.Fatal(err)
	}
	tok, err := conf.Exchange(oauth2.NoContext, code)
	if err != nil {
		log.Fatal(err)
	}

	client := conf.Client(oauth2.NoContext, tok)
	client.Get("...")
}

func ExampleToken() {

    conf := &oauth2.Config{
        ClientID:     "CLIENTID",
        ClientSecret: "CLIENTSECRET",
        Scopes:       []string{"non-expiring"},
        RedirectURL: "https://provider.com/redir",
        Endpoint: oauth2.Endpoint{
            AuthURL:  "http://localhost:8443/auth",
            TokenURL: "http://localhost:8443/token",
        },
    }

    tok, err := conf.PasswordCredentialsToken(oauth2.NoContext, "jason", "pass")

    if err != nil {
        log.Fatalf("Error trying to create token %v", err)
        return
    }

    log.Printf("token %v", tok)
}

*/


var (
    jwtTestDefaultKey []byte
    defaultKeyFunc    jwt.Keyfunc = func(t *jwt.Token) (interface{}, error) { return jwtTestDefaultKey, nil }
)

func main() {

    l, err := ldap.Dial("tcp", fmt.Sprintf("%s:%d", ldapServer, ldapPort))
    if err != nil {
        log.Fatalf("ERROR: %s\n", err.Error())
        return
    }
    defer l.Close()
    // l.Debug = true

    err = l.Bind(user, passwd)
    if err != nil {
        log.Printf("ERROR: Cannot bind: %s\n", err.Error())
        return
    }

    log.Printf("success")

    /*
    var pemBytes = `-----BEGIN EC PRIVATE KEY-----
    MHcCAQEEIKGOgzn9u8RCSwwJj0sGOog6QGpDNkCuBRNsv76bRXLYoAoGCCqGSM49
    AwEHoUQDQgAEPAYLQF6I4NQ1Q0AjeHqJj7fDX/WwJ6xba5aDQ7V9pIQfq8k+JUME
    RUBF85MS+jPu5Rn+59AP9aPRSybIQsxZrg==
    -----END EC PRIVATE KEY-----`


    var pk = `-----BEGIN CERTIFICATE-----
    MIIDwzCCAqugAwIBAgIJAKEXB1nRlSWvMA0GCSqGSIb3DQEBCwUAMHgxCzAJBgNV
    BAYTAlVTMQswCQYDVQQIDAJXQTEQMA4GA1UEBwwHU2VhdHRsZTELMAkGA1UECgwC
    SkoxCzAJBgNVBAsMAkpKMQswCQYDVQQDDAJKSjEjMCEGCSqGSIb3DQEJARYUampv
    c2VwaHlAaG90bWFpbC5jb20wHhcNMTUxMjAyMTk1MzE0WhcNMTYxMjAxMTk1MzE0
    WjB4MQswCQYDVQQGEwJVUzELMAkGA1UECAwCV0ExEDAOBgNVBAcMB1NlYXR0bGUx
    CzAJBgNVBAoMAkpKMQswCQYDVQQLDAJKSjELMAkGA1UEAwwCSkoxIzAhBgkqhkiG
    9w0BCQEWFGpqb3NlcGh5QGhvdG1haWwuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOC
    AQ8AMIIBCgKCAQEA0ygdIO6u6jeVObVP/dzX8ey4fi9goF79sEajQc3NVDeLWYui
    CWPowcgfZdtZBJuIFffyX89qDDAh2lqxd6vT3GzD2g0R1FVCowp4qsMKXYc4wzSn
    3fnJm+ilQtQ/klYXcRQ0NxtptYb/wT9LMqBEvIbL46QawvND/CbiNBqyjUYxE35D
    ZLQRx+4Ec8LsfeulJS2UrOi7Z70D4w7XxlZLRvVaSPMgbn4AWlqjsA+w0H+FvUr5
    nlz8wBpR/IUCizCKZXggS3HF46iaCKkUoZCCvmn0tJFfLTh3z5sapQ9M7G4sAghs
    lxdZygDzOl5RkkFYsE2ZwOsd207wIOqSL+Ag0QIDAQABo1AwTjAdBgNVHQ4EFgQU
    sR4okoMfeNd/kwX5kQWZCxGtPq0wHwYDVR0jBBgwFoAUsR4okoMfeNd/kwX5kQWZ
    CxGtPq0wDAYDVR0TBAUwAwEB/zANBgkqhkiG9w0BAQsFAAOCAQEAq4hn0FYamdLp
    visa0+BMHlEt36pYxH/EUrJ7UnOiRwrdK5t9lxIiwLyh9k6zo1B1ezX9AQhbHdpK
    qy+jN5Ee5o4YfIwnN5npQoLh432/q/xkyufuaTHIjyMbtsOwqF/HxISFL4KJ8pYa
    Jmd/6Mbcp5W2R0q5b90biYmo/SkPDNf7YGd6A7weGHALyUNhzW+gz99qn5SKX9pO
    wkdOtfHlzpgWeQOU2z7unZM2LaFayVS/RzLua1g3NWI4o83cOTpzcwdkFNN8nwru
    An+7GCH1hgB2NmNp26cTKn6sGrz1gkJXxxZMj4ocGjRXf6SZ0lK09NrXGYaodimw
    mkUNuhtVPg==
    -----END CERTIFICATE-----`
    */
/*
    block, _ := pem.Decode([]byte(pemBytes))
    privatekey, err := x509.ParseECPrivateKey(block.Bytes)

    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }

    fmt.Printf("Public Key :  : %x\n\n", privatekey.PublicKey)


    fmt.Printf("Private Key D :  : %x\n\n", privatekey.D)
*/
    //ParseECPrivateKey(der []byte) (key *ecdsa.PrivateKey, err error)

    /*
    var e error

    if jwtTestDefaultKey, e = ioutil.ReadFile("cert.pem"); e != nil {
        panic(e)
    }

    // Create the token
	token := jwt.New(jwt.SigningMethodHS256)
	// Set some claims
	token.Claims["foo"] = "bar"
	token.Claims["exp"] = time.Now().Add(time.Hour * 72).Unix()
	// Sign and get the complete encoded token as a string
	tokenString, ex := token.SignedString(jwtTestDefaultKey)

    if ex != nil {
        log.Printf("Error %v", ex)
        return
    }

    log.Println(tokenString)
    log.Println("")

    token, err := jwt.Parse(tokenString, defaultKeyFunc)

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return defaultKeyFunc(token.Header["kid"]), nil
    })
    */

    /*
    log.Println("h %v", token.Valid)

    if token.Valid {
        fmt.Println("You look nice today")
    } else if ve, ok := err.(*jwt.ValidationError); ok {
        if ve.Errors&jwt.ValidationErrorMalformed != 0 {
            fmt.Println("That's not even a token")
        } else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
        // Token is either expired or not active yet
            fmt.Println("Timing is everything")
        } else {
            fmt.Println("Couldn't handle this token:", err)
        }
    } else {
        fmt.Println("Couldn't handle this token:", err)
    }
    */

    //ExampleToken()
    //ExampleConfig()

    /*

    */

    /*
    search := ldap.NewSearchRequest(
        baseDN,
        ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
        filter,
        Attributes,
        nil)

    sr, err := l.Search(search)
    if err != nil {
        log.Fatalf("ERROR: %s\n", err.Error())
        return
    }

    log.Printf("Search: %s -> num of entries = %d\n", search.Filter, len(sr.Entries))
    sr.PrettyPrint(0)
    */
}
