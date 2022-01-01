package main

import (
  "bytes"
  "testing"
)

// Test_CountWords tests the count function set to count words
func Test_CountWords(t *testing.T) {
  b := bytes.NewBufferString("word1 word2 word3 word4\n")

  exp := 4

  res := count(b, false)

  if res != exp {
    t.Errorf("Expected %d, got %d instead.\n", exp, res)
  }
}

// Test_CountLines tests the count function set to count lines
func Test_CountLines(t *testing.T) {
  b := bytes.NewBufferString("word1 word2 word3\nline2\nline3 word1")

  exp := 3

  res := count(b, true)

  if res != exp {
    t.Errorf("Expected %d, got %d instead.\n", exp, res)
  }
}
