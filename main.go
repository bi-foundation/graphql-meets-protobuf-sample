package main

import (
	"encoding/json"
	"fmt"
	"github.com/FactomProject/graphql-meets-protobuf-sample/models"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/query", queryHandler).Methods(http.MethodPost)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func queryHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading request body: %v", err), http.StatusBadRequest)
		return
	}

	data, err := getData()
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading data: %v", err), http.StatusBadRequest)
		return
	}

	result, err := Query(string(body), data)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing query: %v", err), http.StatusBadRequest)
		return
	}

	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error reading writing response: %v", err), http.StatusBadRequest)
		return
	}
}

func Query(filtering string, person *models.Person) (interface{}, error) {
	schema, err := queryScheme(person)
	if err != nil {
		return nil, fmt.Errorf("failed to create schema: %v", err)
	}

	// inject filtering in query
	query := fmt.Sprintf(`{ person %s }`, filtering)
	params := graphql.Params{Schema: schema, RequestString: query}
	result := graphql.Do(params)

	if len(result.Errors) > 0 {
		return nil, fmt.Errorf("failed to execute graphql operation: %v", result.Errors)
	}

	return result.Data, nil
}

func queryScheme(person interface{}) (graphql.Schema, error) {
	return graphql.NewSchema(graphql.SchemaConfig{
		Query: graphql.NewObject(graphql.ObjectConfig{
			Name: "Query",
			Fields: graphql.Fields{
				"person": &graphql.Field{
					Type: models.GraphQLPersonType,
					Resolve: func(p graphql.ResolveParams) (interface{}, error) {
						return person, nil
					},
				},
			},
		}),
	})
}

func getData() (*models.Person, error) {
	// retrieve the latest data
	data, err := ioutil.ReadFile("data.bin")
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %v", err)
	}

	person := &models.Person{}
	err = proto.Unmarshal(data, person)
	if err != nil {
		return nil, fmt.Errorf("failed to read data: %v", err)
	}

	return person, nil
}
