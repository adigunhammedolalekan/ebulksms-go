package client

import "testing"

func TestNewSmsClient(t *testing.T) {
	_, err := NewSmsClient("", "")
	if err == nil {
		t.Fatal("error should not be nil for empty username and key")
	}
}

func TestEbulkSmsClient_Send(t *testing.T) {
	c, err := NewSmsClient("testUser@gmail.com", "testKey")
	if err != nil {
		t.Fatal(err)
	}
	m := &Message{
		Text:   "Hola! Your code is: 88097",
		Phones: []string{"phone1", "phone2"},
		Flash:  "0",
		Dnd:    0,
		Sender: "Hola Text",
	}
	r, err := c.Send(m)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(r)
}
