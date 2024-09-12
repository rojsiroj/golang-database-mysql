package repository

import (
	"context"
	"fmt"
	golangdatabasemysql "golang-database-mysql"
	"golang-database-mysql/entity"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

func TestCommentInsert(t *testing.T) {
	db, err := golangdatabasemysql.GetConnection()
	if err != nil {
		panic(err)
	}

	commentRepository := NewCommentRepository(db)

	ctx := context.Background()
	comment := entity.Comment{Email: "repository@test.com", Comment: "Test repository"}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func TestFindById(t *testing.T) {

	db, err := golangdatabasemysql.GetConnection()
	if err != nil {
		panic(err)
	}

	commentRepository := NewCommentRepository(db)
	ctx := context.Background()

	comment, err := commentRepository.FindById(ctx, 23)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}
func TestFindAll(t *testing.T) {

	db, err := golangdatabasemysql.GetConnection()
	if err != nil {
		panic(err)
	}

	commentRepository := NewCommentRepository(db)
	ctx := context.Background()

	comments, err := commentRepository.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println(comments)
}
