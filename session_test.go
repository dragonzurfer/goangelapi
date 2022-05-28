package goangelapi_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/dragonzurfer/goangelapi"
	"github.com/dragonzurfer/goangelapi/smartapigo"
	"github.com/stretchr/testify/mock"
)

func TestGetClientSessionSuccess(t *testing.T) {
	client := goangelapi.GetClient()
	session, err := goangelapi.GetClientSession(client)
	if err != nil {
		t.Fatalf("%s", err.Error())
	}
	fmt.Println(session.ClientCode)
}

type mockFailClient struct {
	mock.Mock
}

func (m *mockFailClient) GenerateSession() (smartapigo.UserSession, error) {
	args := m.Called()
	res := args.Get(0)
	err := args.Get(1)
	return res.(smartapigo.UserSession), err.(error)
}

func (m *mockFailClient) GetUserProfile() (smartapigo.UserProfile, error) {
	return smartapigo.UserProfile{}, errors.New("Test Error")
}

func (m *mockFailClient) GetPositions() (smartapigo.Positions, error) {
	var x smartapigo.Positions
	return x, errors.New("Test Error")
}
func (m *mockFailClient) GetClosedPositions() (smartapigo.Positions, error) {
	var x smartapigo.Positions
	return x, errors.New("Test Error")
}
func (m *mockFailClient) GetOpenPositions() (smartapigo.Positions, error) {
	var x smartapigo.Positions
	return x, errors.New("Test Error")
}

func TestGetClientSessionFail(t *testing.T) {
	testclient := new(mockFailClient)
	err_message := "Test Error"
	testclient.On("GenerateSession").Return(smartapigo.UserSession{}, errors.New(err_message))
	goangelapi.SetExpRetryMaxSleepDuration("10")
	goangelapi.SetExpRetrySleepDuration("1")
	_, err := goangelapi.GetClientSession(testclient)
	if err == nil {
		t.Fatalf("Session Fail test succeeding not hitting api rate limit")
	}

	if err.Error() != err_message {
		t.Errorf("want error message %s got %s", err_message, err.Error())
	}

}
