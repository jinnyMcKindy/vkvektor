package main

import (
	"errors"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"io/ioutil"
	"net/http"
)

func getVkPhoto(urlc string) (img []byte, err error) {
	res, err := http.Get(urlc)
	if err != nil {
		return img, err
	}

	root, err := html.Parse(res.Body)
	if err != nil {
		return img, err
	}
	avatarin, ok := scrape.Find(root, scrape.ByClass("pp_img"))
	if !ok || avatarin.DataAtom != atom.Img {
		return img, errors.New("no <a> by class pp_img")
	}
	avapurl := scrape.Attr(avatarin, "src")
	if avapurl == "" {
		return img, errors.New("no href at profile: " + urlc)
	}
	res.Body.Close()

	res, err = http.Get(avapurl)
	if err != nil {
		return img, err
	}
	defer res.Body.Close()

	return ioutil.ReadAll(res.Body)
}
