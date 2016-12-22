package main

import (
	"bytes"
	"encoding/gob"
)

type P struct {
	X string
	Y int64
	Z int
}

type Q struct {
	X string
	Y int32
}

func main()  {
	var buff bytes.Buffer

	encoder := gob.NewEncoder(&buff)
	decoder := gob.NewDecoder(&buff)

	sent := P{"satendra", 26, 0}
	encoder.Encode(sent)

	received := Q{}
	decoder.Decode(&received)

	println(received.X, received.Y)
}
