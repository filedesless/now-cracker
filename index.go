package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"
	"time"
	"io/ioutil"
	"os"
)

type thing struct {
	time time.Time
	pass string
	hash string
}

var things = make([]thing, 0)

func filterOutOldThings(t time.Time) {
	new_things := make([]thing, 0)
	for _, thing := range things {
		if thing.time.Add(time.Second * 2).After(t) {
			new_things = append(new_things, thing)
		}
	}
	things = new_things
}

func addNewThing(t time.Time) string {
	pass := strconv.Itoa(1000 + rand.Intn(9000))
	hash := md5.Sum([]byte(pass))
	hexDigest := hex.EncodeToString(hash[:])
	things = append(things, thing{t, pass, hexDigest})
	return hexDigest
}	

func Handler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	filterOutOldThings(now)
	switch r.Method {
	case "POST":
		givenHash := r.URL.Path[1:]
		body, _ := ioutil.ReadAll(r.Body)
		givenPass := string(body)
		for _, thing := range things {
			if thing.hash == givenHash && thing.pass == givenPass {
				fmt.Fprintln(w, os.Getenv("FLAG"))
			}
		}
	case "GET":
		hash := addNewThing(now)
		fmt.Fprintln(w, hash)
	}
}
