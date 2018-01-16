package main

import (
	"time"

	gocql "github.com/gocql/gocql"
	"github.com/mitchellh/mapstructure"
)

const (
	ChannelStop = iota
	UserStop
	MessageStop
)

type Channel struct {
	ID   gocql.UUID `cql:"id,omitempty" json:"id"`
	Name string     `cql:"name" json:"name"`
}

type User struct {
	ID   gocql.UUID `cql:"id,omitempty" json:"id"`
	Name string     `cql:"name" json:"name"`
}

type ChannelMessage struct {
	ID        gocql.UUID `cql:"id,omitempty" json:"id"`
	ChannelId string     `cql:"channelid" json:"channelid"`
	Body      string     `cql:"body" json:"body"`
	Author    string     `cql:"author" json:"author"`
	CreatedAt string     `cql:"createdat" json:"createdat"`
}

func addChannel(client *Client, data interface{}) {
	var channel Channel
	err := mapstructure.Decode(data, &channel)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	uuid, err := gocql.RandomUUID()
	if err != nil {
		client.send <- Message{"error", err.Error()}
	}
	channel.ID = uuid
	go func() {
		if err = client.session.Query(`INSERT INTO channel (id,name) VALUES (?, ?)`, uuid, channel.Name).Exec(); err != nil {
			client.send <- Message{"error", err.Error()}
		}
		client.send <- Message{"channel add", channel}
	}()
}

func subscribeChannel(client *Client, data interface{}) {
	stop := client.NewStopChannel(ChannelStop)
	result := make(chan Channel)
	var id gocql.UUID
	var name string
	iter := client.session.Query(`SELECT id,name FROM channel`).Iter()
	go func() {
		for iter.Scan(&id, &name) {
			channel := Channel{
				ID:   id,
				Name: name,
			}
			result <- channel
		}
	}()
	go func() {
		for {
			select {
			case <-stop:
				if err := iter.Close(); err != nil {
					client.send <- Message{"error", "err.Error()"}
					return
				}
			case change := <-result:
				client.send <- Message{"channel add", change}
			}
		}
	}()
}

func unsubscribeChannel(client *Client, data interface{}) {
	client.StopForKey(ChannelStop)
}

func editUser(client *Client, data interface{}) {
	var user User
	err := mapstructure.Decode(data, &user)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	client.userName = user.Name
	user.ID = client.userId
	go func() {
		if err = client.session.Query(`UPDATE user set name = ? where id = ?`, user.Name, client.userId).Exec(); err != nil {
			client.send <- Message{"error", err.Error()}
		}
		client.send <- Message{"user edit", user}
	}()
}

func subscribeUser(client *Client, data interface{}) {
	stop := client.NewStopChannel(UserStop)
	result := make(chan User)
	var id gocql.UUID
	var name string
	iter := client.session.Query(`SELECT id,name FROM user`).Iter()
	go func() {
		for iter.Scan(&id, &name) {
			user := User{
				ID:   id,
				Name: name,
			}
			result <- user
		}
	}()
	go func() {
		for {
			select {
			case <-stop:
				if err := iter.Close(); err != nil {
					client.send <- Message{"error", "err.Error()"}
					return
				}
			case change := <-result:
				client.send <- Message{"user add", change}
			}
		}
	}()
}

func unsubscribeUser(client *Client, data interface{}) {
	client.StopForKey(UserStop)
}

func addChannelMessage(client *Client, data interface{}) {
	var channelMessage ChannelMessage
	err := mapstructure.Decode(data, &channelMessage)
	if err != nil {
		client.send <- Message{"error", err.Error()}
		return
	}
	uuid, err := gocql.RandomUUID()
	if err != nil {
		client.send <- Message{"error", err.Error()}
	}
	channelMessage.ID = uuid
	channelMessage.CreatedAt = time.Now().Format("02/01/2006 03:04:05 PM")
	channelMessage.Author = client.userName
	go func() {
		if err = client.session.Query(`INSERT INTO message (id,author,body,channelid,createat) VALUES (?, ?, ?, ?, ?)`,
			uuid, channelMessage.Author, channelMessage.Body, channelMessage.ChannelId, channelMessage.CreatedAt).Exec(); err != nil {
			client.send <- Message{"error", err.Error()}
		}
		client.send <- Message{"message add", channelMessage}
	}()
}

func subscribeChannelMessage(client *Client, data interface{}) {
	stop := client.NewStopChannel(MessageStop)
	result := make(chan ChannelMessage)
	var id gocql.UUID
	var author string
	var body string
	var createAt string
	var iter *gocql.Iter
	go func() {
		eventData := data.(map[string]interface{})
		val, ok := eventData["channelId"]
		if !ok {
			return
		}
		channelId, ok := val.(string)
		if !ok {
			return
		}
		iter = client.session.Query(`SELECT id,author,body,createat FROM message where channelId = ? order by createat`, channelId).Iter()
		for iter.Scan(&id, &author, &body, &createAt) {
			channelMessage := ChannelMessage{
				ID:        id,
				Author:    author,
				Body:      body,
				ChannelId: channelId,
				CreatedAt: createAt,
			}
			result <- channelMessage
		}
	}()
	go func() {
		for {
			select {
			case <-stop:
				if err := iter.Close(); err != nil {
					client.send <- Message{"error", err.Error()}
					return
				}
			case change := <-result:
				client.send <- Message{"message add", change}
			}
		}
	}()
}

func unsubscribeChannelMessage(client *Client, data interface{}) {
	client.StopForKey(MessageStop)
}
