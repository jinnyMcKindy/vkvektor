package main

import (
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
	"time"
)

// ========== mongo config

type mongoDB struct {
	Host     string
	Port     string
	Addrs    string
	Database string
	Username string
	Password string
	Info     *mgo.DialInfo
	Session  *mgo.Session
}

func (mongo *mongoDB) setDefault() {
	mongo.Port = os.Getenv("MONGO_PORT")
	mongo.Host = os.Getenv("MONGO_HOST")
	mongo.Addrs = mongo.Host + ":" + mongo.Port
	mongo.Database = os.Getenv("MONGO_NAME")
	mongo.Username = os.Getenv("MONGO_USER")
	mongo.Password = os.Getenv("MONGO_PASSWORD")
	mongo.Info = &mgo.DialInfo{
		Addrs:    []string{mongo.Addrs},
		Timeout:  50 * time.Second,
		Database: mongo.Database,
		Username: mongo.Username,
		Password: mongo.Password,
	}
	err := mongo.setSession()
	if err != nil {
		panic("db connection does not exist")
	}
}

func (mongo *mongoDB) setSession() (err error) {
	mongo.Session, err = mgo.DialWithInfo(mongo.Info)
	if err != nil {
		mongo.Session, err = mgo.Dial(mongo.Host)
	}
	return err
}

func (mongo *mongoDB) drop() (err error) {
	session := mongo.Session.Clone()
	defer session.Close()
	err = session.DB(mongo.Database).C("datas").DropCollection()
	return err
}

func (mongo *mongoDB) init() (err error) {
	err = mongo.drop()
	if err != nil {
		log.Println(err)
	}

	session := mongo.Session.Clone()
	defer session.Close()
	session.EnsureSafe(&mgo.Safe{})

	// ========== datas
	collection := session.DB(mongo.Database).C("datas")
	index := mgo.Index{
		Key:        []string{"rectangle", "vector"},
		Unique:     false,
		Background: true,
		Sparse:     true,
	}
	err = collection.EnsureIndex(index)
	if err != nil {
		return err
	}

	index = mgo.Index{
		Key:        []string{"url"},
		Unique:     true,
		Background: true,
		Sparse:     true,
	}
	err = collection.EnsureIndex(index)
	if err != nil {
		return err
	}

	index = mgo.Index{
		Key:    []string{"$text:img"},
		Unique: false,
	}
	err = collection.EnsureIndex(index)

	return err
}

// ========== res

func (mongo *mongoDB) insertRes(r res) (err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	err = session.DB(mongo.Database).C("datas").Insert(&r)
	return err
}

func (mongo *mongoDB) getDatas() (ress []res, err error) {
	session := mongo.Session.Clone()
	defer session.Close()

	err = session.DB(mongo.Database).C("datas").Find(bson.M{}).All(&ress)
	return ress, err
}
