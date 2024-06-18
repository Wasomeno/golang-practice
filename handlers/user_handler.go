package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func (queries queries) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	type createUserPayload struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	payload := createUserPayload{}

	err := decoder.Decode(&payload)
	if err != nil {
		respondWithError(w, 400, "Error when parsing request body")
		return
	}

	createdUser, err := queries.CreateUser(r.Context(), tutorial.CreateUserParams{
		ID: pgtype.UUID{Bytes: uuid.New(), Valid: true},
		CreatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		UpdatedAt: pgtype.Timestamp{
			Time:  time.Now(),
			Valid: true,
		},
		Name: payload.Name,
	})

	log.Println(tutorial.CreateUserParams{
		ID: pgtype.UUID{},
		CreatedAt: pgtype.Timestamp{
			Time: time.Now(),
		},
		UpdatedAt: pgtype.Timestamp{
			Time: time.Now(),
		},
		Name: payload.Name,
	})
	if err != nil {
		respondWithError(w, 500, "Error when creating user")
		return
	}
	respondWithJSON(w, 200, createdUser)
}


func (queries queries) handleGetUser(w http.ResponseWriter, r *http.Request) {
	urlQuery := r.URL.Query()

	userQuery := urlQuery.Get("name")

	user, err := queries.GetUser(r.Context(), userQuery)

	log.Println(err)

	if err != nil {
		respondWithError(w, 500, "Error when getting user")
		return
	}

	

	respondWithJSON(w, 200, user)
}