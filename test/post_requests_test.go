package test

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"testing"
)

func TestCreateContact(t *testing.T) {
	values := url.Values{
		"name":        {"Tralala Ozoz"},
		"phonenumber": {"+79887766521"},
	}
	u := "http://127.0.0.1:12345/createContact"
	res, err := http.PostForm(u, values)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	b := string(body)

	fmt.Println(b)
}

func TestChangeContact(t *testing.T) {
	values := url.Values{
		"name":        {"Innion Was"},
		"phonenumber": {"+79877665409"},
	}
	u := "http://127.0.0.1:12345/changeContact/id=1"
	res, err := http.PostForm(u, values)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	b := string(body)

	fmt.Println(b)
}

func TestDeleteContact(t *testing.T) {
	u := "http://127.0.0.1:12345/deleteContact/id=1"
	res, err := http.PostForm(u, nil)
	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	defer res.Body.Close()

	b := string(body)

	fmt.Println(b)
}
