package main
/*
import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-faker/faker/v4"


	sql "./sql"
)

func genUser() (string, string, int) {
	f := faker.New()
	return f.Name(), f.Email(), f.Number().Between(1, 100)
}



func postUserAddHandler(w http.ResponseWriter, r *http.Request) {
	// Validate request
	if r.Method != "POST" {
		m := fmt.Sprintf("%d Bad Request: %s", http.StatusBadRequest, "HTTP Method is not valid")
		httpError(w, m)
		return
	}

	if r.Header.Get("Content-Type") != "application/json" {
		m := fmt.Sprintf("%d Bad Request: %s", http.StatusBadRequest, "Content Type is not valid")
		httpError(w, m)
		return
	}

	// Read request body
	params, err := readJson(r)
	if err != nil {
		httpError(w, err.Error())
		return
	}

	if params["Name"] == nil || params["Email"] == nil || params["Age"] == nil {
		m := fmt.Sprintf("%d Internal Server Error: %s", http.StatusInternalServerError, "Invalid Input")
		httpError(w, m)
		return
	}

 	ctx := r.Context()
	age, err := strconv.Atoi(params["Age"].(string))
	if err != nil {
		m := fmt.Sprintf("%d Internal Server Error: %s", http.StatusInternalServerError, "Invalid Input")
		httpError(w, m)
		return
	}
	u, err := query.CreateUser(ctx, db.CreateUserParams{
		Name:  params["Name"].(string),
		Email: params["Email"].(string),
		Age:   int32(age),
	})
	if err != nil {
		m := fmt.Sprintf("%d Internal Server Error: %s", http.StatusInternalServerError, err)
		httpError(w, m)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf("User Created: %v\n", u)))
}
*/
