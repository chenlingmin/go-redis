package database

import "go-redis/interface/resp"

type CmdLine = [][]byte

type Database interface {
	Exec(client resp.Connection, cmdLine [][]byte) resp.Reply
	AfterClientClose(c resp.Connection)
	Close()
}

type DataEntity struct {
	Data interface{}
}
