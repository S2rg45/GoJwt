package main

import (
    "io"
    "log"
    "net/http"
    "time"
	"github.com/gin-gonic/gin"
    "github.com/auth0/go-jwt-middleware"
    "github.com/dgrijalva/jwt-go"
)

const (
    APP_KEY = "golangcode.com"
)

var ( 
	router 									   = gin.Default() 
)

func main() {

	mapUrls()
    route.Run(":8080")
}

func mapUrls(){
	route.POST("/token", TokenHandler)
	route.POST("/", AuthMiddleware(http.HandlerFunc(ExampleHandler)))
}

// TokenHandler is our handler to take a username and password and,
// if it's valid, return a token used for future requests.
func TokenHandler(w http.ResponseWriter, r *http.Request) {

    w.Header().Add("Content-Type", "application/json")
    r.ParseForm()

    // Check the credentials provided - if you store these in a database then
    // this is where your query would go to check.
    username := r.Form.Get("username")
    password := r.Form.Get("password")
    if username != "myusername" || password != "mypassword" {
        w.WriteHeader(http.StatusUnauthorized)
        io.WriteString(w, `{"error":"invalid_credentials"}`)
        return
    }

    // We are happy with the credentials, so build a token. We've given it
    // an expiry of 1 hour.
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user": username,
        "exp":  time.Now().Add(time.Minute * time.Duration(2)).Unix(),
        "iat":  time.Now().Unix(),
    })
    tokenString, err := token.SignedString([]byte(APP_KEY))
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        io.WriteString(w, `{"error":"token_generation_failed"}`)
        return
    }
    io.WriteString(w, `{"token":"`+tokenString+`"}`)
    return
}

// AuthMiddleware is our middleware to check our token is valid. Returning
// a 401 status to the client if it is not valid.
func AuthMiddleware(next http.Handler) http.Handler {
    if len(APP_KEY) == 0 {
        log.Fatal("HTTP server unable to start, expected an APP_KEY for JWT auth")
    }
    jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
        ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
            return []byte(APP_KEY), nil
        },
        SigningMethod: jwt.SigningMethodHS256,
    })
    return jwtMiddleware.Handler(next)
}

func ExampleHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    io.WriteString(w, `{"status":"ok"}`)
}



