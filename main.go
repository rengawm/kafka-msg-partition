package main

import (
	"flag"
	"fmt"

	"github.com/aviddiviner/go-murmur"
)

var partitions int
var msgKey string

func init() {
	flag.IntVar(&partitions, "partitions", 12, "how many partitions there are")
	flag.StringVar(&msgKey, "msg-key", "", "key to compute partition of")
}

func main() {
	flag.Parse()

	hash := murmur.MurmurHash2([]byte(msgKey), 0x9747b28c) & 0x7fffffff
	fmt.Printf("Partition: %v\n", hash%uint32(partitions))
}
