package authentication

import (
	"crypto/rsa"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	"gitlab.com/fabstao/rokobooking/models"

	jwt "github.com/dgrijalva/jwt-go"
)

var (
	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
)

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	privateBytes, err := ioutil.ReadFile("./roko.key")
	checkError(err)
	publicBytes, err := ioutil.ReadFile("./roko.pub")
	checkError(err)
	privateKey, err = jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	checkError(err)
	publicKey, err = jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	checkError(err)

}

// GenerateJWT :
func GenerateJWT(user models.User) (string, error) {
	claims := models.Claim{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 4).Unix(),
			Issuer:    "Fabs",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)
	result, err := token.SignedString(privateKey)
	//checkError(err)
	return result, err
}

// ValidateToken :
/*
func ValidateToken(w http.ResponseWriter, r *http.Request, _ httprouter.Params) (interface{}, error) {
	token, err := request.ParseFromRequestWithClaims(r, request.OAuth2Extractor, &models.Claim{}, func(token *jwt.Token) (interface{}, error) {
		return publicKey, nil
	})
	if err != nil {
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				fmt.Println(w, "<h3>Token expirado</h3>")
				return nil, err
			case jwt.ValidationErrorSignatureInvalid:
				fmt.Println(w, "<h3>Token con firma inválida</h3>")
				return nil, err
			default:
				fmt.Println(w, "<h3>Token inválido</h3>")
				return nil, err
			}
		default:
			fmt.Println(w, "<h3>Token inválido</h3>")
			return nil, err
		}
	}

	if token.Valid {
		w.WriteHeader(http.StatusAccepted)
		w.Write([]byte("{\"status\":\"ok\""))
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("{\"error\":\"Invalid token\"}"))
	}
}
*/

// ValidateToken : versión sin Request
func ValidateToken(tokenString string, user models.User) (interface{}, error) {
	claims := models.Claim{}
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		//return []byte("AllYourBase"), nil
		return publicKey, nil
	})
	//fmt.Println(token.Claims)
	if err != nil {
		fmt.Println("Error: ", err)
		switch err.(type) {
		case *jwt.ValidationError:
			vErr := err.(*jwt.ValidationError)
			switch vErr.Errors {
			case jwt.ValidationErrorExpired:
				fmt.Println("Token expirado")
				return nil, err
			case jwt.ValidationErrorSignatureInvalid:
				fmt.Println("Token con firma inválida")
				return nil, err
			default:
				fmt.Println("Token inválido")
				return nil, err
			}
		default:
			fmt.Println("Token inválido")
			return nil, err
		}
	}

	if token.Valid {
		return struct{ Status string }{Status: "ok"}, nil
	}
	return struct{ Error string }{Error: "token inválido"}, nil
}
