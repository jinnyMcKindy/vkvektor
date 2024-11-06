package main

import (
	"log"
	"testing"
)

func TestGetVkPhoto(t *testing.T) {
	img, err := getVkPhoto("https://vk.com/id101734124")
	if err != nil {
		t.Error(err)
		return
	}
	if len(img) < 1 {
		t.Error("empty temp link")
	}
	log.Println(img)
}
