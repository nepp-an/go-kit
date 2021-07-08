package main

import (
    "fmt"
    "testing"
)

type MockDB struct {
    // FakeFetch is used to provide unique test case results
    FakeFetch func(string) (int, error)
}

func (m *MockDB) Fetch(k string) (int, error) {
   if m.FakeFetch != nil {
       return m.FakeFetch(k)
   }
   return 0, nil
}

func TestIsOver9k(t *testing.T) {
    t.Run("TestOver9000", func(t *testing.T) {
        DB = &MockDB{
            FakeFetch: func(string) (int, error) {
                return 9001, nil
            },
        }
        if !isOver9000(DB) {
            t.Errorf("Test did not return expected results")
        }
    })

    t.Run("TestUnder9000", func(t *testing.T) {
        DB = &MockDB{
            FakeFetch: func(string) (int, error) {
                return 8999, nil
            },
        }
        if isOver9000(DB) {
            t.Errorf("Test did not return expected results")
        }
    })

    t.Run("TestDBError", func(t *testing.T) {
        DB = &MockDB{
            FakeFetch: func(string) (int, error) {
                return 0, fmt.Errorf("unable to connect to database")
            },
        }
        if isOver9000(DB) {
            t.Errorf("Test did not return expected results")
        }
    })
}