package main

import (
	"errors"
	"fmt"
	"net/http"
)

type SimpleDataStore struct {
	userData map[string]string
}

func (sds SimpleDataStore) UserNameForID(userID string) (string, bool) {
	name, ok := sds.userData[userID]
	return name, ok
}

func NewSimpleDataStore() SimpleDataStore {
	return SimpleDataStore{userData: map[string]string{
		"1": "Renata",
		"2": "Guilherme",
		"3": "Martin",
	}}
}

type DataStore interface {
	UserNameForID(userID string) (string, bool)
}

func LogOutPut(message string) {
	fmt.Println(message)
}

type Logger interface {
	Log(message string)
}

type LoggerAdapter func(message string)

func (lg LoggerAdapter) Log(message string) {
	lg(message)
}

type SimpleLogic struct {
	l  Logger
	ds DataStore
}

func NewSimpleLogic(l Logger, ds DataStore) SimpleLogic {
	return SimpleLogic{
		l:  l,
		ds: ds,
	}
}

func (sl SimpleLogic) SayHello(userID string) (string, error) {
	sl.l.Log("in SayHello for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Hello " + name, nil
}

func (sl SimpleLogic) SayGoodbye(userID string) (string, error) {
	sl.l.Log("in SayGoodbye for " + userID)
	name, ok := sl.ds.UserNameForID(userID)
	if !ok {
		return "", errors.New("unknown user")
	}
	return "Goodbye " + name, nil
}

// implicitly implements Logic
type Logic interface {
	SayHello(userID string) (string, error)
}

type Controller struct {
	logger Logger
	logic  Logic
}

func (c Controller) SayHello(w http.ResponseWriter, r *http.Request) {
	c.logger.Log("In SayHello")
	userID := r.URL.Query().Get("user_id")
	message, err := c.logic.SayHello(userID)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	w.Write([]byte(message))
}

func NewController(logger Logger, logic Logic) Controller {
	return Controller{
		logger: logger,
		logic:  logic,
	}
}
func main() {
	l := LoggerAdapter(LogOutPut)
	ds := NewSimpleDataStore()
	logic := NewSimpleLogic(l, ds)
	c := NewController(l, logic)
	http.HandleFunc("/hello", c.SayHello)
	http.ListenAndServe(":8080", nil)
}
