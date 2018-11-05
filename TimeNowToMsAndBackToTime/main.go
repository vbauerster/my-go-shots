package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("UTC now is: %s\n", now.UTC())
	nanoNow := time.Duration(now.UnixNano()) * time.Nanosecond
	msNow := nanoNow / time.Millisecond

	fmt.Printf("Now in Milliseconds: %d\n", int64(msNow))

	d := msNow * time.Millisecond
	t := time.Unix(int64(d/time.Second), int64(d%time.Second)).UTC()
	fmt.Printf("Milliseconds back to %T: %[1]s\n", t)
}
