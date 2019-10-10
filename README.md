# GraphQL-meets-Protocbufs in Go
GraphQL meets Protocol Buffers Sample project. A simple project that takes a protobuf model and generates the serialization functions and GraphQL schemes that allow users to query data through an API with GraphQL.

We have an API written in Golang and want to extend it with GraphQL support. The API receives events with protocol-buffers from Google and distributed these events to users. Now we want the API to expose this data to the outside world and allow users to query this data with GraphQL. Protocol buffers is a mechanism for serializing structured data. You write your data definition and generate source code for serialize- and deserialization. The generated code can easily write and read your structured data to and from a variety of data streams and using a variety of languages. The data streams can be transferred or stored in a fast binary way using a variety of languages. GraphQL is a query language for your API, and a server-side runtime for executing queries by using a type system you define for your data.

The problem is that it would be easy for small models, but when the size and amount of models increases so does the effort to support GraphQL. After writing the protocol buffers and generating the source code for the models, we need to write the GraphQL scheme to be able to execute the queries on the data. With only a few models this is manageable. When the number of models increases, so does the line of codes we need to write to support GraphQL. The situation gets easily out of hand and maintaining the code would be a nightmare. So why not directly generate the support for GraphQL.

First, we need to install the protocol buffer compiler. This allows us to generate the source code for data streams in this case Go.  To install the compiler on Linux we executed the following commands. For other platforms or more information see https://developers.google.com/protocol-buffers.

``` shell script
wget https://github.com/protocolbuffers/protobuf/releases/download/v3.9.0/protoc-3.9.0-linux-x86_64.zip
unzip protoc-3.9.0-linux-x86_64.zip -d protoc3
sudo mv protoc3/bin/* /usr/local/bin/
sudo mv protoc3/include/* /usr/local/include/
rm protoc-3.9.0-linux-x86_64.zip
```

Then we need some Go dependencies. We want to have Go support for Protocol Buffers and the Protobuf GraphQL extension.

```shell script
go get -u github.com/golang/protobuf/protoc-gen-go
go get -u github.com/bi-foundation/protobuf-graphql-extension/protoc-gen-gogoopsee
```

Now we are ready to write some code. We start writing the protobuf definition. Just a simple example of a person with a phone number. We add the option to generate the GraphQL schemes.

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

The protoc compiler generates two files, models.pb.go and modelspb_test.go. The file contains for each object and variable serialization methods and also GraphQL schemes. The source contains data definitions,  functions to read and write binary data, and GraphQL type definitions. Note that just by enabling the graphql plugin in the gogoopsee_out flag we get the GraphQL type definitions. The plugin extends the protobuf generator and uses the information gathered to generate GraphQL types for each object. 

Now we need to be able to query the generated models. We write a function that takes the query and the data and produces the filtered data. A query scheme is created that wraps the data in a query. It links the query to the data whereon it's executed. It assigns the data of the person to the GraphQL Person type. 

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
 
As we want users to query the data, build a simple API that handles a user request to retrieve the filtered data. In main we start an HTTP listener on port 8080. When a post is received with the query, the queryHandler retrieves the latest data and executes the query on the data.  

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

That’s it. Now we can test it. We start the API and post a GraphQL query. We execute the following curl to retrieve the name and the number of a person:

```shell script
curl -X POST http://localhost:8080/query -d "{ name phone { number } }"
```

Response:
```json
{"person":{"name":"Jaap Joosten","phone":{"number":"053218622189"}}}
```


