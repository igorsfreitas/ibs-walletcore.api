package database

import (
	"database/sql"
	"testing"

	"github.com/igorsfreitas/ibs-walletcore.api/internal/entity"
	_ "github.com/mattn/go-sqlite3"
	"github.com/stretchr/testify/suite"
)

type ClientDBTestSuite struct {
	suite.Suite
	db       *sql.DB
	clientDB *ClientDB
}

func (s *ClientDBTestSuite) SetupSuite() {
	db, err := sql.Open("sqlite3", ":memory:")
	s.Nil(err)
	s.db = db

	db.Exec(`CREATE TABLE clients (id varchar(255), name varchar(255), email varchar(255), created_at date)`)
	s.clientDB = NewClientDB(db)
}

func (s *ClientDBTestSuite) TearDownSuite() {
	defer s.db.Close()
	s.db.Exec(`DROP TABLE clients`)
}

func TestClientDBTestSuite(t *testing.T) {
	suite.Run(t, new(ClientDBTestSuite))
}

func (s *ClientDBTestSuite) TestGetClient() {
	client, _ := entity.NewClient("Teste", "teste@teste.com.br")
	s.clientDB.SaveClient(client)

	clientDB, err := s.clientDB.GetClient(client.ID)
	s.Nil(err)
	s.Equal(client.ID, clientDB.ID)
	s.Equal(client.Name, clientDB.Name)
	s.Equal(client.Email, clientDB.Email)
}

func (s *ClientDBTestSuite) TestSaveClient() {
	client := &entity.Client{
		ID:    "123",
		Name:  "Teste",
		Email: "teste@teste.com",
	}
	err := s.clientDB.SaveClient(client)
	s.Nil(err)
}
