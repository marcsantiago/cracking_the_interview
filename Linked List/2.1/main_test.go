package main

import (
	"fmt"
	"testing"
)

func TestAppendToTail(t *testing.T) {
	mockData := []int{2, 3, 4, 5, 6, 7, 8, 9, 10}
	ll := New(1)
	for _, n := range mockData {
		ll.AppendToTail(n)
	}
	testData := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var i int
	for {
		if ll.Next == nil {
			break
		}
		if testData[i] != ll.Data {
			t.Errorf("the data should be the same. Test data: %d linked list data %d", testData[i], ll.Data)
		}
		ll = ll.Next
		i++
	}
}

func TestRemoveDups1(t *testing.T) {
	mockData := []int{2, 3, 4, 5, 5, 7, 5, 9, 10}
	ll := New(1)
	for _, n := range mockData {
		ll.AppendToTail(n)
	}
	ll.RemoveDups1()
	fmt.Println(ll)
	testData := []int{1, 2, 3, 4, 5, 7, 9, 10}
	var i int
	for {
		if ll.Next == nil {
			break
		}
		if testData[i] != ll.Data {
			t.Errorf("the data should be the same. Test data: %d linked list data %d", testData[i], ll.Data)
		}
		ll = ll.Next
		i++
	}
}

func TestRemoveDups2(t *testing.T) {
	mockData := []int{9, 2, 3, 4, 5, 5, 7, 5, 9, 10}
	ll := New(1)
	for _, n := range mockData {
		ll.AppendToTail(n)
	}
	ll.RemoveDups2()
	fmt.Println(ll)
	testData := []int{1, 9, 2, 3, 4, 5, 7, 10}
	var i int
	for {
		if testData[i] != ll.Data {
			t.Errorf("the data should be the same. Test data: %d linked list data %d", testData[i], ll.Data)
		}
		ll = ll.Next
		i++
		if ll.Next == nil {
			break
		}
	}
}
