package goangelapi

import (
	"errors"

	"github.com/dragonzurfer/goangelapi/smartapigo"
)

func GetClientSession(client ClientInterface) (smartapigo.UserSession, error) {
	var session smartapigo.UserSession

	// retry exponentially since without session we can't do anything
	sessionInterface, err := ExponentialRetry(func() (interface{}, error) {
		session, err := client.GenerateSession()
		return session, err
	})
	// check if hit max retries
	if err != nil {
		return session, err
	}

	// cast interface to type
	session = sessionInterface.(smartapigo.UserSession)
	session.UserProfile, err = client.GetUserProfile()
	if err != nil {
		return session, errors.New("unable to get user profile:" + err.Error())
	}

	return session, nil
}
