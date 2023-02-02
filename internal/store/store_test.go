package store

import (
	"gotest.tools/v3/assert"
	"testing"
)

type TestItem struct {
    Id int
}

func setupStore() []*Store {
    items := []TestItems{
        {Id: 1},
        {Id: 2},
        {Id: 3},
        {Id: 4},
        {Id: 5},
    }
    return NewStore[TestItem](items)
}

func TestAdd(t *testing.T) {
    s := &Store[TestItem]{}
    s.Add(TestItem{id: 1337})
    items := s.ListAll()
    assert.Equal(t, len(items), 1)
    assert.Equal(t, items[0].Id, 1337)
}

func TestAddAll(t *testing.T) {
    v := []TestItem{
        {Id: 11},
        {Id: 22},
        {Id: 33}
    }
    s := &Store[TestItem]{}
    s.AddAll(v)
    items := s.ListAll()
    assert.Equal(t, len(items), 3)
    assert.Equal(t, items[0].Id, 11)
    assert.Equal(t, items[1].Id, 22)
    assert.Equal(t, items[2].Id, 33)
}

func TestDeleteOne(t *testing.T) {
    s := setupStore()
}

func TestDeleteAll(t *testing.T) {
    s := setupStore()
}

func TestFindOne(t *testing.T) {
    s := setupStore()
}

func TestFindAll(t *testing.T) {
    s := setupStore()
}
