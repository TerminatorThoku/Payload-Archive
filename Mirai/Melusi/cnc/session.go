package main

import (
	"sync"
	"time"
)

var Sessions []Session
var sessionMutex sync.Mutex

type Session struct {
	user   string
	expire int64
}

func addSession(user string) {

	sessionMutex.Lock()
	Sessions = append(Sessions, Session{
		user:   user,
		expire: time.Now().Add(120 * time.Second).Unix(),
	})
	sessionMutex.Unlock()
}

func sessionExists(user string) bool {
	for _, session := range Sessions {
		if session.user == user {
			if session.expire > time.Now().Unix() {
				return true
			}

			sessionRemove(user)
		}
	}

	return false
}

func sessionKeepAlive(user string) {
	for i, session := range Sessions {
		if session.user == user {
			sessionMutex.Lock()
			Sessions[i].expire = time.Now().Add(120 * time.Second).Unix()
			sessionMutex.Unlock()
		}
	}
}

func sessionRemove(user string) {

	for i, session := range Sessions {
		if session.user == user {
			sessionMutex.Lock()
			Sessions = append(Sessions[:i], Sessions[i+1:]...)
			sessionMutex.Unlock()
			return
		}
	}
}
