package main

import (
	"GoTest/redis/provider"
	"testing"
)

const COUNT = 100000


func BenchmarkDoHash(t *testing.B)  {
	for i:=0;i<COUNT;i++{
		DoHashStore(provider.Conn)
	}
}



func BenchmarkDoEncodingStore(t *testing.B)  {
	for i:=0;i<COUNT;i++ {
		DoGobEncodingStore(provider.Conn)
	}
}



func BenchmarkDoJsonEncodingStore(t *testing.B)  {
	for i:=0;i<COUNT;i++ {
		DoJsonEncodingStore(provider.Conn)
	}
}