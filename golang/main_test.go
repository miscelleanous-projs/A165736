package main

import (
	"math/big"
	"testing"
)

func TestA165736(t *testing.T) {
	assert := func(condition bool) {
		if !condition {
			t.Fatal("Assertion failed")
		}
	}

	assert(a165736(big.NewInt(  1)).Cmp(big.NewInt(1)) == 0)
	assert(a165736(big.NewInt(  6)).Cmp(big.NewInt(7447238656)) == 0)
	assert(a165736(big.NewInt( 10)).Cmp(big.NewInt(0)) == 0)
	assert(a165736(big.NewInt( 11)).Cmp(big.NewInt(9172666611)) == 0)
	assert(a165736(big.NewInt( 16)).Cmp(big.NewInt(290415616)) == 0)
	assert(a165736(big.NewInt( 19)).Cmp(big.NewInt(609963179)) == 0)
	assert(a165736(big.NewInt( 20)).Cmp(big.NewInt(0)) == 0)
	assert(a165736(big.NewInt( 23)).Cmp(big.NewInt(1075718247)) == 0)
	assert(a165736(big.NewInt( 30)).Cmp(big.NewInt(0)) == 0)
	assert(a165736(big.NewInt( 40)).Cmp(big.NewInt(0)) == 0)
	assert(a165736(big.NewInt(100)).Cmp(big.NewInt(0)) == 0)
}
