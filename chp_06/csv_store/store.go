package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// Post is a forum post struct.
type Post struct {
	ID      int
	Content string
	Author  string
}

func main() {
	/* Create a CSV file and write some posts to it */
	csvFile, err := os.Create("posts.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	allPosts := []Post{
		Post{ID: 1, Content: "Hello world!", Author: "Sau Sheong"},
		Post{ID: 2, Content: "Bonjour, monde!", Author: "Pierre"},
		Post{ID: 3, Content: "Hola, Mundo!", Author: "Pedro"},
		Post{ID: 4, Content: "Greetings, earthlings!", Author: "The Martian"},
	}

	writer := csv.NewWriter(csvFile)

	for _, post := range allPosts {
		line := []string{strconv.Itoa(post.ID), post.Content, post.Author}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()

	/* Read the posts from the CSV file created earlier */
	file, err := os.Open("posts.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	var posts []Post
	for _, item := range record {
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{ID: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Printf("*****************************************\n")
	for _, post := range posts {
		fmt.Printf(" Post ID:\t%d\n Author:\t%s\n Content:\t%s\n", post.ID, post.Author, post.Content)
		fmt.Printf("*****************************************\n")
	}
}
