#! /bin/bash
#---
# Excerpted from "Powerful Command-Line Applications in Go",
# published by The Pragmatic Bookshelf.
# Copyrights apply to this code. It may not be used to create training material,
# courses, books, articles, and the like. Contact us if you are in doubt.
# We make no guarantees that this code is fit for any purpose.
# Visit https://pragprog.com/titles/rggo for more book information.
#---

FHASH=`md5sum $1`
while true; do
  NHASH=`md5sum $1`
  if [ "$NHASH" != "$FHASH" ]; then 
    ./mdp -file $1
    FHASH=$NHASH
  fi
  sleep 5
done

