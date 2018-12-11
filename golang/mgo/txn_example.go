package mgo

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
	"log"
)

const (
	DbUrl       = "mongodb://iris:irispassword@35.221.94.114:27217" // mongodb://myuser:mypass@localhost:40001
	DbDatabase  = "sync-iris"
	TCollection = "mgo_txn"
)

var (
	session *mgo.Session
)

type block struct {
	Height int64  `bson:"height"`
	Hash   string `bson:"hash"`
}

type tx struct {
	Hash        string `bson:"hash"`
	BlockHeight int64  `bson:"block_height"`
	Content     string `bson:"content"`
}

func init() {
	dialInfo, err := mgo.ParseURL(DbUrl + "/" + DbDatabase)
	if err != nil {
		panic(err)
	}
	s, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	session = s
}

func newCollection(collectionName string) *mgo.Collection {
	database := session.DB(DbDatabase)
	return database.C(collectionName)
}

func BulkInsertTxn(blocks []block, txs []tx) {
	var (
		blockInsertOps []txn.Op
		txInsertOps    []txn.Op
	)

	for _, v := range blocks {
		op := txn.Op{
			C:      "block",
			Id:     bson.NewObjectId(),
			Assert: txn.DocMissing,
			Insert: v,
		}
		blockInsertOps = append(blockInsertOps, op)
	}

	for _, v := range txs {
		op := txn.Op{
			C:      "tx",
			Id:     bson.NewObjectId(),
			Assert: txn.DocMissing,
			Insert: v,
		}
		txInsertOps = append(txInsertOps, op)
	}

	runner := txn.NewRunner(newCollection(TCollection))
	runner.ChangeLog(newCollection("log"))
	ops := append(blockInsertOps, txInsertOps...)
	if err := runner.Run(ops, bson.NewObjectId(), nil); err != nil {
		panic(err)
	} else {
		log.Println("transaction of bulk insert success")
	}
}
