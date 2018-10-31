package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now().UTC()
	fmt.Println(now)
	nanoNow := time.Duration(now.UnixNano()) * time.Nanosecond
	till := nanoNow / time.Millisecond
	from := (nanoNow - 24*time.Hour) / time.Millisecond

	fmt.Printf("From in Millisecond: %d\n", int64(from))
	fmt.Printf("Till in Millisecond: %d\n", int64(till))

	d := till * time.Millisecond
	t := time.Unix(int64(d/time.Second), int64(d%time.Second)).UTC()
	fmt.Println(t)

	d = from * time.Millisecond
	t = time.Unix(int64(d/time.Second), int64(d%time.Second)).UTC()
	fmt.Println(t)
}
