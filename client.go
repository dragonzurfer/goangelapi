package goangelapi

import (
	"errors"
	"os"
	"strconv"

	"github.com/dragonzurfer/goangelapi/smartapigo"
)

var (
	API_KEY   = os.Getenv("AngelAPIKEY")
	CLIENT_ID = os.Getenv("AngelClientID")
	PASSWORD  = os.Getenv("Password")
)

func setAPIkey(apikey string) {
	API_KEY = apikey
}

func setClientID(clientid string) {
	CLIENT_ID = clientid
}

func setPassword(pass string) {
	PASSWORD = pass
}

func SetCredentials(apikey, clientid, pass string) {
	setAPIkey(apikey)
	setClientID(clientid)
	setPassword(pass)
}

type ClientInterface interface {
	GenerateSession() (smartapigo.UserSession, error)
	GetUserProfile() (smartapigo.UserProfile, error)
	GetPositions() (smartapigo.Positions, error)
	GetOpenPositions() (smartapigo.Positions, error)
	GetClosedPositions() (smartapigo.Positions, error)
}

type Client struct {
	SmartAPIClient *smartapigo.Client
}

func (c *Client) GenerateSession() (smartapigo.UserSession, error) {
	return c.SmartAPIClient.GenerateSession()
}

func (c *Client) GetUserProfile() (smartapigo.UserProfile, error) {
	return c.SmartAPIClient.GetUserProfile()
}

// ** Parms order apikey, clientid, password
func GetClient(params ...string) *Client {
	if len(params) > 0 {
		setAPIkey(params[0])
	}
	if len(params) > 1 {
		setClientID(params[1])
	}
	if len(params) > 2 {
		setPassword(params[2])
	}
	smartapiClient := smartapigo.New(CLIENT_ID, PASSWORD, API_KEY)
	client := &Client{
		SmartAPIClient: smartapiClient,
	}
	// set them back in case len(params) > 0
	SetCredentials(os.Getenv("AngelAPIKEY"), os.Getenv("AngelClientID"), os.Getenv("Password"))

	return client
}

func (c *Client) GetPositions() (smartapigo.Positions, error) {
	positions, err := c.SmartAPIClient.GetPositions()
	return positions, err
}

func (c *Client) GetOpenPositions() (smartapigo.Positions, error) {
	var openPositions smartapigo.Positions
	positions, err := c.SmartAPIClient.GetPositions()
	if err != nil {
		return positions, err
	}

	// open positions have abs(NetQty) > 0
	for _, p := range positions {
		netQty, err := strconv.ParseFloat(p.NetQty, 64)
		if err != nil {
			return openPositions, errors.New("error parsing float64 Angel api position.NetQty which is type string")
		}
		if netQty != 0 {
			openPositions = append(openPositions, p)
		}
	}
	return openPositions, nil
}

func (c *Client) GetClosedPositions() (smartapigo.Positions, error) {
	var closedPosition smartapigo.Positions
	positions, err := c.SmartAPIClient.GetPositions()
	if err != nil {
		return positions, err
	}

	// closed positions have abs(NetQty) = 0
	for _, p := range positions {
		netQty, err := strconv.ParseFloat(p.NetQty, 64)
		if err != nil {
			return closedPosition, errors.New("error parsing float64 Angel api position.NetQty which is type string")
		}
		if netQty == 0 {
			closedPosition = append(closedPosition, p)
		}
	}
	return closedPosition, nil
}
