package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"math/rand"

	"github.com/google/uuid"
	"github.com/go-faker/faker/v4"
 	"go.opentelemetry.io/otel"

	sqlc "nnaka2992/jaguer_o11y_20250307/sql"
)

 func genUser() (string, string, string, string) {
	return faker.Username(), faker.Name(), faker.Email(), faker.Password()
}

func genItem() (string, sql.NullString, int64) {
	return faker.Word(), sql.NullString{String: faker.Sentence(), Valid: true}, int64(rand.Intn(10000-500+1) + 500)
}

func postUserAddHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer(instrumentationName).Start(context.Background(), "postUserAddHandler")
	defer span.End()

	// validate request before processing
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	id, name, email, password := genUser()
	user, err := sqlc.DB.CreateUser(ctx, sqlc.CreateUserParams{id, name, email, password})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return response to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{"id": "%s", "name": "%s", "email": "%s"}`, user.ID, user.Name, user.Email)))
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer(instrumentationName).Start(context.Background(), "getUserHandler")
	defer span.End()

	// get user id from request
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// get user id from request
	id := r.URL.Query().Get ("id")
	if id == "" {
		http.Error(w, "Invalid user id", http.StatusBadRequest)
		return
	}
	// get user from database
	user, err := sqlc.DB.GetUserByID(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return response to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"id": "%s", "name": "%s", "email": "%s"}`, user.ID, user.Name, user.Email)))
}

func postItemAddHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer(instrumentationName).Start(context.Background(), "postItemAddHandler")
	defer span.End()

	// validate request before processing
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	name, description, price := genItem()
	item, err := sqlc.DB.CreateItem(ctx, sqlc.CreateItemParams{name, description, price})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return response to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf(`{"id": "%s", "name": "%s", "description": "%s", "price": %d}`, item.ID, item.Name, item.Description, item.Price)))
}

func getItemHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer(instrumentationName).Start(context.Background(), "getItemHandler")
	defer span.End()


	if r.Method != http.MethodGet {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	// check if item id is set.
	// if not, return all items
	id := r.URL.Query().Get("id")
	if id == "" {
		// get all items from database by N+1 query
		ids, err := sqlc.DB.GetItemIds(ctx)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		var items []sqlc.Item
		for _, id := range ids {
			item, err := sqlc.DB.GetItemByID(ctx, id)
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			items = append(items, item)
		}
		// return response to client
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(fmt.Sprintf(`{"items": %v}`, items)))
	}
	// get item from database
	uuid, err := uuid.Parse(id)
	if err != nil {
		http.Error(w, "Invalid item id", http.StatusBadRequest)
		return
	}
	item, err := sqlc.DB.GetItemByID(ctx, uuid)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// return response to client
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(`{"id": "%s", "name": "%s", "description": "%s", "price": %d}`, item.ID, item.Name, item.Description, item.Price)))
}
