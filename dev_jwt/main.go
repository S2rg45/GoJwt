package main

import (
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
	"net/http"
	"log"
)

var (
	mySigningKey = []byte("MySuperSecret")
)

func homePage(w http.ResponseWriter, r *http.Request){
	validToken, err :=  GenerateJWT()
	if err != nil {
		fmt.Fprintf(w, err.Error())
	}
	fmt.Fprintf(w, validToken)
}

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["autorized"] = true
	claims["user"] = "Eliot forbes"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)
	if err != nil {
		fmt.Errorf("Something went wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func handleRequests(){
	http.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func main() {
	fmt.Println("My simple client")

	handleRequests()
	// tokenString, err := GenerateJWT()
	// if err != nil {
	// 	fmt.Println("Error generating token string")
	// }
	// fmt.Println(tokenString)
}