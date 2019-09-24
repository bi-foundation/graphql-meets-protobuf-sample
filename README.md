# GraphQL-meets-Protocbufs
GraphQL meets Protocol Buffers Sample project. A simple project that takes a protobuf model and generates the serialization functions and GraphQL schemes that allow users to query data through an API with GraphQL.

We have an API written in golang that communicates with protocol-buffers from Google. Now we want the API to expose this data to the outside world and allow users to query this data with GraphQL. Protocol buffers is a mechanism for serializing structured data. You write your data definition and generate source code for serialize and deserialization. The datastreams can be transferred or stored in a fast binary way using a variety of languages. GraphQL is a query language for your API, and a server-side runtime for executing queries by using a type system you define for your data.
 
The problem is that it would be easy for small models, but when the size and amount of models increases so does effort to support GraphQL. After writing the protocol buffers and generating the source code for the models, we need to write the GraphQL scheme to be able to execute the queries on the data. With only a few models this is manageable. When the number of models increase, so does the line of codes we need to write to support GraphQL. The situation get easily out of hand and maintaining the code would be a nightmare. So why not directly generate the support for GraphQL.

First we need to install the protocol buffer compiler. For more information https://developers.google.com/protocol-buffers

For linux x86-64 we do:
``` shell script
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.9.0/protoc-3.9.0-linux-x86_64.zip
unzip protoc-3.9.0-linux-x86_64.zip -d protoc3
sudo mv protoc3/bin/* /usr/local/bin/
sudo mv protoc3/include/* /usr/local/include/
rm protoc-3.9.0-linux-x86_64.zip
```

Then we need some golang dependencies.

```shell script
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/bi-foundation/protobuf-graphql-extension/protoc-gen-gogoopsee
```

Now we are ready to write some code. We start writing the protobuf. Just a simple one as an example. We add the option to generate the GraphQL schemes.

```protobuf
syntax = "proto3";
package models;

import "github.com/bi-foundation/protobuf-graphql-extension/graphqlproto/graphql.proto";

option (graphqlproto.graphql) = true;

message Person {
    string name = 1;
    int32 id = 2;
    string email = 3;
    PhoneNumber phone = 4;
}

message PhoneNumber {
    string number = 1;
    PhoneType type = 2;
}

enum PhoneType {
    MOBILE = 0;
    HOME = 1;
    WORK = 2;
}
```
 

Now we are ready to generate some source code:
```shell script
protoc --gogoopsee_out=plugins=grpc+graphql,Mopsee/protobuf/opsee.proto=github.com/opsee/protobuf/opseeproto,Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:./models --proto_path=$GOPATH/src:. *.proto
```

The protoc program generates two files, models.pb.go and modelspb_test.go. The file contains for each object and variable serialization methods and also GraphQL schemes.

Now we need to be able to query the generated models. We write a function that takes the query and the data and produce a result. A query scheme is created that wraps the data in a query. It assigns the person data to the GraphQL Person type.

```go
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
```
 
As we want users to query the data, build a simple API that handles a user request to retrieve the filtered data. In main we start an HTTP listener on port 8080. When a post 

```go
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
```

That’s it. Now we can test it. We start the API and do post with a GraphQL query. We execute the following curl to retrieve the name and the number of a person:

```shell script
curl -X POST http://localhost:8080/query -d "{ name phone { number } }"
```

Response:
```json
{"person":{"name":"Jaap Joosten","phone":{"number":"053218622189"}}}
```


