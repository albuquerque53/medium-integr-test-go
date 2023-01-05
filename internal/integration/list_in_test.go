package integration

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/golang-migrate/migrate/source/file"
)

func (s *Suite) TestListUsers() {
	type sceneries struct {
		description string
		expctUsers  []user
	}

	testSceneries := []sceneries{
		{
			description: "Must return one user",
			expctUsers: []user{
				{Name: "James Hetfield", Email: "hetfield@napster.com"},
			},
		},
		{
			description: "Must return many users",
			expctUsers: []user{
				{Name: "James Hetfield", Email: "hetfield@napster.com"},
				{Name: "Lars Ulrich", Email: "ulrich@napster.com"},
			},
		},
		{
			description: "Must return no users", expctUsers: []user{},
		},
	}

	for _, scenery := range testSceneries {
		s.Run(scenery.description, func() {
			for _, expctUser := range scenery.expctUsers {
				s.DBConn.Query("INSERT INTO users(name, email) VALUES(?, ?)", expctUser.Name, expctUser.Email)
			}

			resp, err := http.Get(fmt.Sprintf("%s/list", s.Srv.URL))

			s.NoError(err, "Expected no error during request")
			s.Equal(200, resp.StatusCode)

			defer resp.Body.Close()

			b, err := io.ReadAll(resp.Body)
			s.NoError(err, "Expected no error during read of response body")

			var r []user

			err = json.Unmarshal(b, &r)
			s.NoError(err, "Expected no error during unmarshal of response body to struct")

			for i := 0; i < len(scenery.expctUsers); i++ {
				expctUser := scenery.expctUsers[i]
				user := r[i]

				s.Equal(expctUser.Name, user.Name)
				s.Equal(expctUser.Email, user.Email)
			}
		})
	}
}
