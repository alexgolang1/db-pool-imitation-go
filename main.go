package main

import (
	"fmt"
	"time"
)

type Conn struct {
	ID int
}

type ConnectionPool struct {
	AvaliableConns chan *Conn
}

func NewPool(maxConns int) *ConnectionPool {

	pool := &ConnectionPool{
		AvaliableConns: make(chan *Conn, maxConns),
	}

	for i := 1; i < maxConns; i++ {
		pool.AvaliableConns <- &Conn{ID: i}
	}

	return pool
}

func (p ConnectionPool) Get() *Conn {

	return <-p.AvaliableConns

}

func (p ConnectionPool) Put(c *Conn) {

	p.AvaliableConns <- c

}

func main() {

	pool := NewPool(3)

	for i := 1; i <= 5; i++ {

		go func() {

			conn := pool.Get()

			fmt.Println("working...")

			pool.Put(conn)

		}()

	}

	time.Sleep(time.Second * 3)

}
