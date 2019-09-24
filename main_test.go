package main

import (
	"bytes"
	"github.com/FactomProject/graphql-meets-protobuf-sample/models"
	"github.com/golang/protobuf/proto"
	"io/ioutil"
	"net/http"
	"reflect"
	"testing"
)

func TestQuery(t *testing.T) {
	go main()

	query := []byte("{ name phone { number } }")
	response, err := http.Post("http://localhost:8080/query", "application/json", bytes.NewBuffer(query))
	if err != nil || response == nil {
		t.Fatal(err)
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)

	t.Logf("response: %s", body)

	expected := `{"person":{"name":"Jaap Joosten","phone":{"number":"053218622189"}}}`
	if reflect.DeepEqual(expected, string(body)) {
		t.Fatalf("response assertion failed: %s != %s", expected, body)
	}
}

func TestOutputData(t *testing.T) {
	person := &models.Person{
		Id:    32,
		Name:  "Jaap Joosten",
		Email: "jaap@joosten",
		Phone: &models.PhoneNumber{
			Number: "053218622189",
			Type:   models.PhoneType_HOME,
		},
	}

	data, err := proto.Marshal(person)
	if err != nil {
		t.Fatalf("failed to marshell data: %v", err)
	}

	err = ioutil.WriteFile("data.bin", data, 0644)
	if err != nil {
		t.Fatalf("failed to write data: %v", err)
	}
}
