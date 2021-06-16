package main

import (
	"context"
	"errors"
	"math/rand"
	"time"

	"github.com/dimitarkovachev/golang-at-ocado/proj/sort/gen"
)

func newSortingService() gen.SortingRobotServer {
	s := &sortingService{}
	s.initChannel()
	s.initRoutineHelper()

	return s
}

type sortingService struct {
	c            chan bool
	items        []*gen.Item
	itemSelected *gen.Item
}

var randSource = rand.NewSource(time.Now().UnixNano())
var random = rand.New(randSource)

func (s *sortingService) LoadItems(ctx context.Context, reqPayload *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {
	return s.loadItems(reqPayload)
}

func (s *sortingService) SelectItem(ctx context.Context, reqPayload *gen.SelectItemRequest) (*gen.SelectItemResponse, error) {
	return s.selectItem()
}

func (s *sortingService) MoveItem(ctx context.Context, reqPayload *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {
	return s.moveItem(reqPayload)
}

func (s *sortingService) loadItems(reqPayload *gen.LoadItemsRequest) (*gen.LoadItemsResponse, error) {
	s.block()
	defer s.unblock()

	s.items = append(s.items, reqPayload.Items...)

	return &gen.LoadItemsResponse{}, nil
}

func (s *sortingService) selectItem() (*gen.SelectItemResponse, error) {
	s.block()
	defer s.unblock()

	if s.itemSelected != nil {
		return nil, errors.New("item already selected in hand")
	}

	if len(s.items) == 0 {
		return nil, errors.New("no items to select")
	}

	itemsCount := len(s.items)
	randomItemIndex := 0
	if itemsCount > 0 {
		randomItemIndex = random.Intn(itemsCount - 1)
	}

	s.itemSelected = s.items[randomItemIndex]
	s.items[randomItemIndex] = s.items[itemsCount-1]
	s.items = s.items[:itemsCount-1]

	return &gen.SelectItemResponse{Item: s.itemSelected}, nil
}

func (s *sortingService) moveItem(reqPayload *gen.MoveItemRequest) (*gen.MoveItemResponse, error) {
	s.block()
	defer s.unblock()

	if s.itemSelected == nil {
		return nil, errors.New("no item in hand")
	}

	s.itemSelected = nil

	return &gen.MoveItemResponse{}, nil
}

func (s *sortingService) block() {
	blocked := true

	for blocked {
		blocked = <-s.c
		if blocked {
			s.c <- blocked
			continue
		}
	}

	s.c <- true
}

func (s *sortingService) unblock() {
	<-s.c
	s.c <- false
}

func (s *sortingService) initChannel() {
	s.c = make(chan bool)
}

func (s *sortingService) initRoutineHelper() {
	go func(c chan bool) {
		c <- false
	}(s.c)

	go func(c chan bool) {
		for {
			cv := <-c
			c <- cv
		}
	}(s.c)
}
