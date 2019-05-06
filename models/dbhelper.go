package models

import (
	"sync"

	"gopkg.in/mgo.v2"
)

const (
	mongoURL string = "mongodb://127.0.0.1:27017"
)

var dbInstance *dbHelper
var dbSession *mgo.Session
var once sync.Once

type dbHelper struct{}

type DBConfig struct {
	DB         string
	Collection string
}

// getDBInstance function can provide single instance
func getDBInstance() *dbHelper {
	once.Do(func() {
		session, err := mgo.Dial(mongoURL)
		if err != nil {
			panic(err)
		}
		if session != nil {
			session.SetPoolLimit(10)
			dbSession = session
			dbInstance = new(dbHelper)
		}
	})
	return dbInstance
}

func connect(config DBConfig) (*mgo.Session, *mgo.Collection) {
	s := dbSession.Clone()
	c := s.DB(config.DB).C(config.Collection)
	return s, c
}

// Insert function can clone DB session to insert data
func (h *dbHelper) insert(config DBConfig, i []interface{}) {
	session, c := connect(config)
	defer session.Close()

	err := c.Insert(i...)
	if err != nil {
		panic(err)
	}
}

// findOne function can query and return only first result
func (h *dbHelper) findOne(config DBConfig, query, selector, result interface{}) {
	session, c := connect(config)
	defer session.Close()

	if err := c.Find(query).Select(selector).One(result); err != nil {
		panic(err)
	}
}

// findAll function can query and return all result
func (h *dbHelper) findAll(config DBConfig, query, selector, result interface{}) {
	session, c := connect(config)
	defer session.Close()

	if err := c.Find(query).Select(selector).All(result); err != nil {
		panic(err)
	}
}

func (h *dbHelper) removeAll(config DBConfig, selector interface{}) {
	session, c := connect(config)
	defer session.Close()

	if _, err := c.RemoveAll(selector); err != nil {
		panic(err)
	}
}
