package integration

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/golang-migrate/migrate/source/file"
)

type user struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (s *Suite) TestNewUsers() {
	type sceneries struct {
		description string
		statCode    int
		user        user
	}

	testSceneries := []sceneries{
		{description: "Must save user", statCode: 201, user: user{Name: "Zeca Baleiro", Email: "zeca@baleiro.com"}},
		{description: "Don't send any user", statCode: 400, user: user{}},
	}

	for _, scenery := range testSceneries {
		s.Run(scenery.description, func() {
			body, _ := json.Marshal(scenery.user)

			resp, err := http.Post(fmt.Sprintf("%s/new", s.Srv.URL), "application/json", bytes.NewReader(body))

			s.NoError(err, "Expected no error during request")
			s.Equal(scenery.statCode, resp.StatusCode)

			defer resp.Body.Close()

			if (user{} == scenery.user) {
				return
			}

			_, err = io.ReadAll(resp.Body)
			s.NoError(err, "Expected no error during read of response body")

			seeInDatabase(s, "users", map[string]string{
				"name":  scenery.user.Name,
				"email": scenery.user.Email,
			})
		})
	}
}
