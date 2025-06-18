package main

import (
	"context"
	"database/sql"
	"log"
	"regexp"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	DB_NEW_LOC    = "./reader_new.sqlite"
	DB_OLD_LOC    = "./reader.sqlite"
	DB_NEWEST_LOC = "./reader_newest.sqlite"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

type Author struct {
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type Book struct {
	Title    string    `json:"title"`
	Excerpts []Excerpt `json:"excerpts"`
}

type Excerpt struct {
	Text      string `json:"text"`
	CreatedAt string `json:"createdAt"`
}

type Annotations struct {
	Annotations []Author `json:"annotations"`
}

func (a *App) GetInfoFromDB() Annotations {
	db, err := sql.Open("sqlite3", DB_OLD_LOC)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query("SELECT DISTINCT t1.Attribution, t1.Title, t2.Text, t2.DateCreated FROM content t1 CROSS JOIN Bookmark t2 WHERE t1.BookID IS NULL AND t1.ContentId LIKE t2.VolumeID")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	re := regexp.MustCompile(`T.*$`)
	authorMap := make(map[string]*Author)
	bookMap := make(map[string]*Book) // key: "authorName|bookTitle"

	for rows.Next() {
		var attribution, title, text, dateStr string

		err = rows.Scan(&attribution, &title, &text, &dateStr)
		if err != nil {
			log.Fatal(err)
		}

		createdAt := formatTime(re.ReplaceAllString(dateStr, ""))
		text = strings.Trim(text, "\n")
		text = strings.Trim(text, " ")

		// Get or create author
		author, exists := authorMap[attribution]
		if !exists {
			author = &Author{Name: attribution, Books: []Book{}}
			authorMap[attribution] = author
		}

		// Get or create book
		bookKey := attribution + "|" + title
		book, exists := bookMap[bookKey]
		if !exists {
			book = &Book{Title: title, Excerpts: []Excerpt{}}
			bookMap[bookKey] = book
			author.Books = append(author.Books, *book)
		}

		// Add excerpt to book
		newExcerpt := Excerpt{
			Text:      excerptFormat(text),
			CreatedAt: createdAt,
		}
		book.Excerpts = append(book.Excerpts, newExcerpt)

		// Update the book in the author's slice
		for i := range author.Books {
			if author.Books[i].Title == title {
				author.Books[i] = *book
				break
			}
		}
	}

	annotations := Annotations{}
	for _, author := range authorMap {
		annotations.Annotations = append(annotations.Annotations, *author)
	}

	return annotations
}
