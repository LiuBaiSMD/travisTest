package main
import (
	"fmt"
	"time"
	"testing"
	)
func TestInit(t *testing.T) {
	Init()
	time.Sleep(time.Second * 5)
	fmt.Println("success over")
}