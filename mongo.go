package main

import (
	"fmt"
	mgo "gopkg.in/mgo.v2"
)

type Conf struct {
	Key string `bson:"key"`
	Value interface{} `bson:"value"`
}

func main() {
	connect := "mongodb://snsgame:NijAL15AmqeMAP27j6V@mongo.master.snsgame/SNS"
	session, err := mgo.Dial(connect)
	if err != nil {
		panic(err)
	}
	collect := session.DB("SNS").C("config")
	data := collect.Find(nil).Iter()

	var onep Conf

	for data.Next(&onep) {
		fmt.Println(onep.Key)
		fmt.Println(onep.Value)
	}
}