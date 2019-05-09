package main

import (
	"testing"
)

var conn = ConnectRedis()
const COUNT = 10000


func BenchmarkDoHash(t *testing.B)  {
	for i:=0;i<COUNT;i++{
		DoHashStore(conn)
	}
}



func BenchmarkDoEncodingStore(t *testing.B)  {
	for i:=0;i<COUNT;i++ {
		DoGobEncodingStore(conn)
	}
}



func BenchmarkDoJsonEncodingStore(t *testing.B)  {
	for i:=0;i<COUNT;i++ {
		DoJsonEncodingStore(conn)
	}
}