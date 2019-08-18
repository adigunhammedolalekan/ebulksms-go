package client

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

/*{ "SMS": {
	"auth": {
	"username": "@gmail.com",
	"apikey": ""

	},
	"message": {
	"sender": "TestName",
	"messagetext": "Your code is: 898 76",
	"flash": "0"
	},
	"recipients": {
	"gsm": [{
	"msidn": "2347017947774",
	"msgid": "uniqueid"
	}]
	},
	"dndsender": 0
	}
}*/

const apiUrl = "http://api.ebulksms.com:8080/sendsms.json"

// SMS is the rep of the above json structure
type SMS struct {
	Auth       *Auth       `json:"auth"`
	Message    *SmsMessage `json:"message"`
	Recipients struct {
		GSM []GSM `json:"gsm"`
	} `json:"recipients"`
	DndSender int `json:"dndsender"`
}

type Auth struct {
	Username string `json:"username"`
	ApiKey   string `json:"apikey"`
}

type SmsMessage struct {
	Sender      string `json:"sender"`
	MessageText string `json:"messagetext"`
	Flash       string `json:"flash"`
}

type GSM struct {
	Msidn string `json:"msidn"`
	Msgid string `json:"msgid"`
}

type Message struct {
	Text   string
	Phones []string
	Flash  string
	Dnd    int
	Sender string
}

// Response is the api response object
/*{
	"response": {
		"status": "SUCCESS",
		"totalsent": 1,
		"cost": 1
	}
}*/
type Response struct {
	Body struct {
		Status    string `json:"status"`
		Cost      int    `json:"cost"`
		TotalSent int    `json:"totalsent"`
	} `json:"response"`
}

type EbulkSmsClient struct {
	httpClient       *http.Client
	Username, ApiKey string
}

func NewSmsClient(username, apiKey string) (*EbulkSmsClient, error) {
	if username == "" || apiKey == "" {
		return nil, errors.New("username or apikey is missing")
	}
	httpClient := &http.Client{Timeout: 60 * time.Second}
	return &EbulkSmsClient{
		httpClient: httpClient,
		Username:   username,
		ApiKey:     apiKey,
	}, nil
}

func (s *EbulkSmsClient) Send(m *Message) (*Response, error) {
	data := &SMS{}
	data.Auth = &Auth{Username: s.Username, ApiKey: s.ApiKey}
	data.DndSender = m.Dnd
	data.Message = &SmsMessage{
		Sender:      m.Sender,
		MessageText: m.Text,
		Flash:       m.Flash,
	}
	data.Recipients.GSM = make([]GSM, 0)
	for _, gsm := range m.Phones {
		data.Recipients.GSM = append(data.Recipients.GSM, GSM{
			Msidn: gsm,
			Msgid: randString(),
		})
	}

	// main request payload
	type payload struct {
		SMS *SMS `json:"SMS"`
	}
	p := &payload{SMS: data}
	b, err := json.Marshal(p)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", apiUrl, bytes.NewReader(b))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := s.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response := &Response{}
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(response); err != nil {
		return nil, err
	}
	return response, nil
}

func randString() string {
	s := time.Now().String()
	m5 := md5.New()
	m5.Write([]byte(s))
	return fmt.Sprintf("%x", m5.Sum(nil))
}