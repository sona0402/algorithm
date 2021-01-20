package main

type St interface {
	Put(k string, value interface{})
	Get(k string) interface{}
	Del(k string) bool
	Keys() STIterator

	Contains(k string) bool
	IsEmpty() bool
	Size() int
}
