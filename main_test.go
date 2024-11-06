package main

import (
	"os"
	"testing"
)

func dbTest() (mongo *mongoDB, err error) {
	os.Setenv("MONGO_NAME", "test")
	os.Setenv("MONGO_USER", "jaime")
	os.Setenv("MONGO_PASSWORD", "123456789")
	os.Setenv("MONGO_HOST", "localhost")
	os.Setenv("MONGO_PORT", "27017")
	mongo = &mongoDB{}
	mongo.setDefault()

	err = mongo.init()
	if err != nil {
		return mongo, err
	}

	return mongo, err
}

func TestStartVkScrap(t *testing.T) {
	m, err := dbTest()
	if err != nil {
		t.Error(err)
		return
	}
	err = startVkScrap(m)
	if err != nil {
		t.Error(err)
		return
	}
}
