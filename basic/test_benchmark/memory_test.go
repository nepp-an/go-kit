package main

import (
    "github.com/aws/aws-sdk-go/aws"
    "sync"
    "testing"
)
type Adapter struct {
    pool sync.Pool
}
func NewAdapter() *Adapter {
    return &Adapter{pool:sync.Pool{New: func() interface{} {
        return make([]byte, 4096)
    }}}
}
func BenchmarkMakeByte(b *testing.B) {
    pool := NewAdapter()
    for i := 0; i < b.N; i++ {
        buffer := pool.pool.Get().([]byte)

        pool.pool.Put(buffer)
        buffer = nil
    }


}

func getByte() {
    buffer := aws.NewWriteAtBuffer(make([]byte, 0,0))
    buffer.Bytes()
}