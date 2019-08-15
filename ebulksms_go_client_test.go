package client

import "testing"

func TestNewSmsClient(t *testing.T) {
	_, err := NewSmsClient("", "")
	if err == nil {
		t.Fatal("error should not be nil for empty username and ket")
	}
}

func TestEbulkSmsClient_Send(t *testing.T) {
	c, err := NewSmsClient("adigunadunfe@gmail.com", "661edb281c3111a871317976887ea6496bdac014")
	if err != nil {
		t.Fatal(err)
	}
	m := &Message{
		Text:   "Hola! Your code is: 88 97",
		Phones: []string{"07035452307", "07017947774"},
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
