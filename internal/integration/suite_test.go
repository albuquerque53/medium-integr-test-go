package integration

import (
	"database/sql"
	"fmt"
	"net/http/httptest"
	"strings"
	"testing"
	"usersapi/internal/application/context"
	"usersapi/internal/application/server"
	"usersapi/internal/infra/db"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type Suite struct {
	suite.Suite                  // obrigatório para funcionamento da suite
	DBConn      *sql.DB          // conexão com o DB que usaremos nos testes
	Srv         *httptest.Server // servidor de teste necessário para chamarmos as rotas
	Migration   *db.Migration    // instância de migration
}

func (s *Suite) SetupSuite() {
	var err error

	s.DBConn = db.ConectToDatabase()

	s.Migration, err = db.CreateMigration(s.DBConn, "../infra/db/migrations")
	require.NoError(s.T(), err)
}

func (s *Suite) TearDownSuite() {
	s.DBConn.Close()
}

// SetupTest irá ser chamado sempre que um teste individual
// começar a rodar
func (s *Suite) SetupTest() {
	err := s.Migration.Up()
	require.NoError(s.T(), err)

	ctx := &context.Context{}
	ctx.SetDBConnection(s.DBConn)

	app := server.SetupServer(ctx)
	s.Srv = httptest.NewServer(app)
}

// SetupTest irá ser chamado sempre que um teste individual
// finalizar a execução
func (s *Suite) TearDownTest() {
	err := s.Migration.Down()
	require.NoError(s.T(), err)

	s.Srv.Close()
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(Suite))
}

func seeInDatabase(s *Suite, table string, search map[string]string) {
	var keys []string
	var values []any
	q := fmt.Sprintf("SELECT COUNT(id) AS total FROM %s WHERE ", table)

	for key, value := range search {
		keys = append(keys, fmt.Sprintf("%s=?", key))
		values = append(values, any(value))
	}

	var row struct {
		Total int `json:"total"`
	}

	err := s.DBConn.QueryRow(q+strings.Join(keys, " and "), values...).Scan(&row.Total)

	s.NoError(err, "Expected no error during seeInDatabase() scan")

	if row.Total > 0 {
		return
	}

	s.Fail(fmt.Sprintf("Could not found registry in %s for %v", table, search))
}
