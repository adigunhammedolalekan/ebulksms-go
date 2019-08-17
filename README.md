### eBulksms-go-client
A simple package to send SMS through http://ebulksms.com API.

### Why?
I use it in some projects and i don't want to keep writing http-client wrapper plus it might be useful to someone out there!.

### You want to use it?
`go get github.com/adigunhammedolalekan/ebulksms-go`

```Go
// create a client
c, err := NewSmsClient("yourUsername", "yourApiKey")
	if err != nil {
		log.Fatal(err)
	}
    // prepare your message
	m := &Message{
		Text:   "Hola! Your code is: 88097",
		Phones: []string{"phone1", "phone2"},
		Flash:  "0",
		Dnd:    0,
		Sender: "Hola Text",
	}
    // send and process response
	r, err := c.Send(m)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(r)
```