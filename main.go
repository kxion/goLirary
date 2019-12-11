package main

import "goLirary/messageTool/kafka"

func main() {
	produce := kafka.NewProduce()
	produce.SendSync([]string{"127.0.0.1:9092"}, "xiaoheiwa", "ceshi", []byte("ceshi123"))
}
