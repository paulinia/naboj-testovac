package main

import "crypto/sha256"

type User struct {
	name      string
	password  [sha256.Size224]byte
	points    int
	avialable []int
	submits   []Submit
}
