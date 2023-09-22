package testing

import (
	"github.com/kafka_example/chat_service/config"
	"github.com/kafka_example/chat_service/genproto/chat_service"
	"github.com/kafka_example/chat_service/pkg/db"
	"github.com/kafka_example/chat_service/storage/postgres"
	"github.com/kafka_example/chat_service/storage/repo"
	"github.com/stretchr/testify/suite"
)

type UserTest struct {
	suite.Suite
	Repo    repo.ChatService
	CleanUp func()
}

func (s *UserTest) SetUpTesting() {
	conndb, cleanUp := db.ConnectToDbSuiteTest(config.Load())
	s.Repo = postgres.NewStorage(conndb)
	s.CleanUp = cleanUp
}

func (s *UserTest) TestCreateUser() {
	userInfo := chat_service.UserReq{
		FirstName: "firstName",
		LastName:  "lastname",
		UserName:  "username",
		Bio:       "bio",
		Phone:     "phone",
		Image:     "image",
	}
	user, err := s.Repo.AddUser(&userInfo)
	s.Nil(err)
	s.NotEmpty(user)
	s.Equal(userInfo.FirstName, user.FirstName)
}

func (s *UserTest) TestGetUser() {
	userId := chat_service.UserId{
		UserId: 23,
	}
	user, err := s.Repo.FindUser(&userId)
	s.Nil(err)
	s.NotEmpty(user)
}

func (s *UserTest) UpdateUser() {
	userInfo := chat_service.UserRes{
		Id:        10,
		FirstName: "firstName",
		LastName:  "lastname",
		UserName:  "username",
		Bio:       "bio",
		Phone:     "phone",
		Image:     "image",
	}
	user, err := s.Repo.Update(&userInfo)
	s.Nil(err)
	s.NotEmpty(user)
}

func (s *UserTest) TestUpdateUser() {
	s.UpdateUser()
}
