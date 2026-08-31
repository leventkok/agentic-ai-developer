package main

import (
	"fmt"
	"log"

	"google.golang.org/protobuf/proto"

	bookmarksv1 "learn/go/day-92/internal/gen/bookmarksv1"
)

func main() {
	req := &bookmarksv1.CreateBookmarkRequest{
		Title: "Go Blog",
		Url:   "https://go.dev/blog",
		Tags:  []string{"go", "protobuf"},
	}

	data, err := proto.Marshal(req)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("marshaled %d bytes\n", len(data))

	var out bookmarksv1.CreateBookmarkRequest
	if err := proto.Unmarshal(data, &out); err != nil {
		log.Fatal(err)
	}

	fmt.Println("title:", out.GetTitle())
	fmt.Println("url:", out.GetUrl())
	fmt.Println("tags:", out.GetTags())
}
