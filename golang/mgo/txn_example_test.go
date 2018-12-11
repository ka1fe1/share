package mgo

import (
	"gopkg.in/mgo.v2/bson"
	"testing"
)

func TestBulkInsertTxn(t *testing.T) {
	type args struct {
		blocks []block
		txs    []tx
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test bulk insert multiple document transaction",
			args: args{
				blocks: []block{
					{
						Height: 1,
						Hash:   bson.NewObjectId().Hex(),
					},
					{
						Height: 2,
						Hash:   bson.NewObjectId().Hex(),
					},
				},
				txs: []tx{
					{
						Hash:        bson.NewObjectId().Hex(),
						BlockHeight: 1,
						Content:     "1",
					},
					{
						Hash:        bson.NewObjectId().Hex(),
						BlockHeight: 2,
						Content:     "1",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			BulkInsertTxn(tt.args.blocks, tt.args.txs)
		})
	}
}
