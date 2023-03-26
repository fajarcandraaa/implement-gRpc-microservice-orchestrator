package faker

// import (
// 	"time"

// 	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/internal/entity/userentity"
// )

// // FakeUser initiate fake data user
// func FakeUser() *userentity.User {
// 	t := time.Now()
// 	fakeUser := &userentity.User{
// 		ID:        UserID001,
// 		Firstname: UserFirstName001,
// 		Lastname:  UserLastName001,
// 		Phone:     UserPhoneNumber001,
// 		Avatar:    UserAvatar001,
// 		Email:     UserEmail001,
// 		Username:  UserUsername001,
// 		Password:  UserID001,
// 		Status:    UserID001,
// 		CreatedAt: t,
// 		UpdatedAt: t,
// 	}

// 	return fakeUser
// }

// // FakeUserInput initiate fake data for payload create or register new user
// func FakeUserInput() *userentity.UserRequest {
// 	fakeUser := &userentity.UserRequest{
// 		Firstname: UserFirstName001,
// 		Lastname:  UserLastName001,
// 		Phone:     UserPhoneNumber001,
// 		Avatar:    UserAvatar001,
// 		Email:     UserEmail001,
// 		Username:  UserUsername001,
// 		Password:  UserID001,
// 		Status:    UserID001,
// 	}

// 	return fakeUser
// }

// // FakeUserUpdate initiate fake payload for update data user
// func FakeUserUpdate() *userentity.UserData {
// 	t := time.Now()
// 	fakeUser := &userentity.UserData{
// 		ID:        UserID001,
// 		Firstname: UserFirstName002,
// 		Lastname:  UserLastName002,
// 		Phone:     UserPhoneNumber001,
// 		Avatar:    UserAvatar003,
// 		Email:     UserEmail001,
// 		Username:  UserUsername002,
// 		Status:    UserStatus001,
// 		UpdatedAt: t.String(),
// 	}

// 	return fakeUser
// }
