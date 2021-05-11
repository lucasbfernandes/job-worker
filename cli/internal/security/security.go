package security

import "errors"

const (
	adminUserAPIToken = "qTMaYIfw8q3esZ6Dv2rQ"

	userAPIToken = "9EzGJOTcMHFMXphfvAuM"
)

// TODO this is mocked because server doesn't have users CRUD and login api
func AuthenticateUser(apiToken string) (*string, error) {
	adminUserBearerToken := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhcGlUb2tlbiI6InFUTWFZSWZ3OHEzZXNaNkR2MnJRIn0.RRjTYWXmLGOI4KYfhgvW2qwgwA9EjP2Xxoo2gKuKHAfsPgBa1XVXdEYKbKdvX8KGcivArUnTDoLxfKUUm6EK9gFofft4o47Yj1hgStSzZIE-UEFOALaOVdJ20orsfFY3lxI90vdCCFTmNgWmzocgZUhScuM4Xn7BQFrmZ82sMLP4wn2GufCIbeL2oz3QephrQJ3aDSP7DueQ_7wY3wYhE2o69m9VGYxVcTbghMIygD1uEIWAQPg6ceApCb0Nke_CYLp1SXDtFBwnR9l8fmYVR40-qmuLyvw-jg3j459jgShuDrNxZvZlZg4qKRSQdBzhJ6kNXyewQ9cB9WYaxmKjMQ"
	userBearerToken := "eyJhbGciOiJSUzUxMiIsInR5cCI6IkpXVCJ9.eyJhcGlUb2tlbiI6IjlFekdKT1RjTUhGTVhwaGZ2QXVNIn0.ZhnpBKFKxcFrPR5FwnpxYxPymScXbq_CBI-x7vuws7wXQBjZT5Z9mtVbb6Mw7pii5GtJTtoTZEAJYzFYWx8akybWBip1cu3hxGo-ZRpgBh7ZOmmR32dln79xKfkUpq2q3u_jN2Gk3VJVUjrihyIeuA2yNuAVWcF_9cGoFmYIhBVpQ0OrmueAYNSNVrDB9wAj8xABePXVTsAyn4cR8AUf3OyP1vQawWxuAsD0bgthoEblMcAPbW4BxIrwS4AFIUM1V0-V-tvbJFXPqsV2Ke_0DSuDaYZwQnQXEXT7OKZL_eMisKINHalwMngUy1M0O9o1mnVKChwu4u9WQ9kIyla3ug"

	if apiToken == adminUserAPIToken {
		return &adminUserBearerToken, nil
	}

	if apiToken == userAPIToken {
		return &userBearerToken, nil
	}

	return nil, errors.New("failed authentication - unauthorized")
}
