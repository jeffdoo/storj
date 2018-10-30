// Copyright (C) 2018 Storj Labs, Inc.
// See LICENSE for copying information.

package teststore

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"storj.io/storj/storage"
)

func TestQueue(t *testing.T) {
	q := NewQueue()
	q.Enqueue(storage.Value("hello"))
	q.Enqueue(storage.Value("world"))
	out, err := q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, out, storage.Value("hello"))
	out, err = q.Dequeue()
	assert.Nil(t, err)
	assert.Equal(t, out, storage.Value("world"))
	out, err = q.Dequeue()
	assert.Nil(t, out)
	assert.NotNil(t, err)
}
