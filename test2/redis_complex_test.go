package main

import (
	"GoTest/redis/provider"
	"testing"
)

const COMPLEX_COUNT = 100


func BenchmarkDoComplexGobEncoding(t *testing.B)  {
	for i:=0;i<COMPLEX_COUNT;i++ {
		DoComplexGobEncodingStore(provider.Conn)
	}
}



func BenchmarkDoComplexJsonEncoding(t *testing.B)  {
	for i:=0;i<COMPLEX_COUNT;i++ {
		DoComplexJSONStore(provider.Conn)
	}
}