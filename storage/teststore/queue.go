// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package teststore

import (
	"container/list"
	"fmt"

	"storj.io/storj/storage"
)

//Queue is a non-threadsafe FIFO queue implementing DistQueue
type Queue struct {
	*list.List
}

//NewQueue returns a queue suitable for testing
func NewQueue() Queue {
	return Queue{list.New()}
}

//Enqueue add a FIFO element, for DistQueue
func (q Queue) Enqueue(value storage.Value) error {
	q.PushBack(value)
	return nil
}

//Dequeue removes a FIFO element, for DistQueue
func (q Queue) Dequeue() (storage.Value, error) {
	for q.Len() > 0 {
		e := q.Front() // First element
		q.Remove(e)    // Dequeue
		return e.Value.(storage.Value), nil
	}
	return nil, fmt.Errorf("queue empty")
}
