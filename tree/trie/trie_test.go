package trie

import (
	"fmt"
	"testing"
	"time"
)

func TestTrie(t *testing.T) {
	root := NewTrie()
	words := []string{"Trie", "insert", "search", "search", "startsWith", "insert", "see"}
	for _, word := range words {
		root.Insert(word)
	}

	isExist := root.Search("starts")
	t.Log("test Search() ", isExist)

	existWord := root.StartsWith("i")
	t.Log("test StartsWith() ", existWord)

	usZone := GetUsZoneLocation()
	datetime := time.Unix(time.Now().Unix(), 0).In(usZone).Format("2006-01-02")
	fmt.Println(datetime)

	formatTime, _ := time.ParseInLocation("2006-01-02", datetime, usZone)
	timestamp := formatTime.Unix()
	fmt.Println(timestamp)

	datetime = time.Unix(timestamp, 0).In(usZone).Format("2006-01-02 15:04:05")
	fmt.Println(datetime)

}

func GetUsZoneLocation() *time.Location {
	usZone, _ := time.LoadLocation("America/New_York")
	return usZone
}
