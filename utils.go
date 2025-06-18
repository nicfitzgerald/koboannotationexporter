package main

import (
	"strings"
	"time"
)

func formatTime(t string) string {
	layout := "2006-01-02"
	rawTime, _ := time.Parse(layout, t)
	return rawTime.Format("January 02, 2006")
	// return formatTime
}

// func display(ann Annotations) {
// 	for _, author := range ann {
// 		fmt.Printf("=== %s ===\n", author.name)
// 		for _, book := range author.books {
// 			fmt.Printf("📖 %s\n", book.title)
// 			for _, excerpt := range book.excerpts {
// 				fmt.Printf("📆 Created at: %s\n%s\n\n", excerpt.createdAt, excerpt.text)
// 			}
// 		}
// 	}
// }

func excerptFormat(e string) string {
	lines := strings.Split(e, "\n")
	for i, line := range lines {
		lines[i] = strings.Trim(line, " \t")
	}

	return strings.Join(lines, "\n")
}
