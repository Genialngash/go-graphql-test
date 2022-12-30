package customemiddleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/Genialngash/graphql-go-test/graph/model"
	"github.com/Genialngash/graphql-go-test/postgress"
	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/pkg/errors"
)

const CurrentUserKey = "currentUser"

func AuthMiddleware(repo postgress.UsersRepo) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := parseToken(r)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				usr, err := repo.GetUserById(claims["jti"].(string))
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}

				ctxt := context.WithValue(r.Context(), CurrentUserKey, usr)
				next.ServeHTTP(w, r.WithContext(ctxt))
			} else {
				next.ServeHTTP(w, r)
				return
			}

		})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{},
	Filter:    stripBearerPrefixFromToken,
}

func stripBearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"

	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}
	return token, nil

}

var authExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, authExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(os.Getenv("JWT_SECRET"))
		return t, nil
	})
	return jwtToken, errors.Wrap(err, "Parse Token error")

}

func GetCurrentUserFromCtxt(ctxt context.Context) (*model.User, error) {
	if ctxt.Value(CurrentUserKey) == nil {
		return nil, errors.New("no user in context")
	}

	user, ok := ctxt.Value(CurrentUserKey).(*model.User)

	if !ok || user.ID == "" {

		return nil, errors.New("No user in context")
	}

	return user, nil
}
