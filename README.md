# implement-gRpc-microservice-orchestrator
![implement-gRpc-microservice-orchestrator](https://entgo.io/images/assets/ent-grpc.jpg)


## Dependency repositories
- Service User  : https://github.com/fajarcandraaa/implement-gRpc-microservice-service-user
- Service Book  : https://github.com/fajarcandraaa/implement-gRpc-microservice-service-book
- Protobank     : https://github.com/fajarcandraaa/implement-gRpc-microservice-protobank


## How to start ?
- `STEP 1 : `We need to clone [Service User](https://github.com/fajarcandraaa/implement-gRpc-microservice-service-user)
- `STEP 2 : `Setup the environment in `.env` like `env.example` according to the settings on your localhost
- `STEP 3 : `Run command on project's terminal : `go mod tidy` and `go mod vendor` to constructs a directory named vendor in the main module's root directory that contains copies of all packages needed to support builds and tests of packages in the main module
- `STEP 4 : `Run command `go run main.go` as usualy
- `STEP 5 : `Clone [Service Book](https://github.com/fajarcandraaa/implement-gRpc-microservice-service-book)
- `STEP 6 : `Repeat `STEP 2` until `STEP 4`
- `STEP 7 : `Dont forget to do `STEP 2` until `STEP 3` in this [Orchestrator Services](https://github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator)
- `STEP 8 : `Execute command `go run main.go` (again) in this [Orchestrator Services](https://github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator)


## List of Endpoint :
- URL : {host}/user/create
    - Method        : POST
    - Body          : { "email" : "", "name" : "", "username" : "", "password" : "" }
    - Body type     : json

- URL : {host}/user/:id/find
    - Method        : GET
    - Path params   : id

- URL : {host}/book/add
    - Method        : POST
    - Body          : { "author" : "", "title" : "" }
    - Body type     : json

- URL : {host}/book/:id/find
    - Method        : GET
    - Path params   : id