package routers

import (
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/handler"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/src/book"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/src/user"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/service/grpc/client"
	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/service/grpc/servicecontract"
)

func (se *Serve) initializeRoutes() {
	//======================== REPOSITORIES ========================
	//initiate repository
	usrgrpc := client.ServiceUser()
	bookgrpc := client.ServiceBook()
	r := servicecontract.NewGrpcService(usrgrpc, bookgrpc)

	//======================== ROUTER ========================
	//Setting Services
	//Setting User Service

	//=== USER ===
	s := user.NewService(r)
	h := handler.NewUserHandler(s)

	//=== BOOK ===
	bs := book.NewService(r) //user.NewService(r)
	bh := handler.NewBookHandler(bs)
	//=========================================================

	//======================== ENDPOINT ========================
	//Initialize endpoint route

	//=== USER ===
	se.Router.HandleFunc("/user/create", h.RegisterNewUser).Methods("POST")
	se.Router.HandleFunc("/user/{id}/find", h.FindUserByUserID).Methods("GET")

	//=== BOOK ===
	se.Router.HandleFunc("/book/add", bh.InsertNewBook).Methods("POST")
	se.Router.HandleFunc("/book/{id}/find", bh.FindBookById).Methods("GET")
	//==========================================================

}
