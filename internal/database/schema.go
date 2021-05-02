package database

import "github.com/hashicorp/go-memdb"

var DBSchema = &memdb.DBSchema{
	Tables: map[string]*memdb.TableSchema{
		"job": {
			Name: "job",
			Indexes: map[string]*memdb.IndexSchema{
				"id": {
					Name:    "id",
					Unique:  true,
					Indexer: &memdb.StringFieldIndex{Field: "Id"},
				},
			},
		},
	},
}
