package main

import (
    "crypto/tls"
    "fmt"
    "time"
)

type Server struct {
    Addr     string
    Port     int
    Protocol string
    Timeout  time.Duration
    MaxConns int
    TLS      *tls.Config
}

type Option func(*Server)

func Protocol(p string) Option {
    return func(s *Server) {
        s.Protocol = p
    }
}
func Timeout(timeout time.Duration) Option {
    return func(s *Server) {
        s.Timeout = timeout
    }
}
func MaxConns(maxconns int) Option {
    return func(s *Server) {
        s.MaxConns = maxconns
    }
}
func TLS(tls *tls.Config) Option {
    return func(s *Server) {
        s.TLS = tls
    }
}

func NewServer(addr string, port int, options ...Option) (*Server, error) {

    srv := Server{
        Addr:     addr,
        Port:     port,
        Protocol: "tcp",
        Timeout:  30 * time.Second,
        MaxConns: 1000,
        TLS:      nil,
    }
    for _, option := range options {
        option(&srv)
    }
    //...
    return &srv, nil
}

func main() {
    s3, _ := NewServer("0.0.0.0", 8080, Timeout(300*time.Second), MaxConns(1200))
    fmt.Println(s3)

}