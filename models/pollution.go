package models

import (
	"log"
)

type pollution struct {
	Site             string
	County           string `bson:"county" json:"county"`
	PM25             string
	DataCreationDate string
	ItemUnit         string
}

// Pollution can export to other file
type Pollution struct {
	DataSlice []pollution
}

var dbConfig = DBConfig{DB: "pollution", Collection: "recent"}

// Get pollution slice
func (p *Pollution) Get() []pollution {
	return p.DataSlice
}

// InsertPollution can insert pollution slice
func (p *Pollution) InsertPollution() {
	log.Println("Save DB start")

	insertSlice := make([]interface{}, len(p.DataSlice))
	for i, pp := range p.DataSlice {
		insertSlice[i] = pp
	}

	getDBInstance().insert(dbConfig, insertSlice)
	log.Println("Save DB end")
}

// FindAllPollution can find all pollution data from db
func (p *Pollution) FindAllPollution() {
	getDBInstance().findAll(dbConfig, nil, nil, &p.DataSlice)
}

// FindPollution can find all pollution data from db
func (p *Pollution) FindPollution(query interface{}) {
	getDBInstance().findAll(dbConfig, query, nil, &p.DataSlice)
}

// RemoveAllPollution can remove all pollution data from db
func (p *Pollution) RemoveAllPollution() {
	getDBInstance().removeAll(dbConfig, nil)
}
