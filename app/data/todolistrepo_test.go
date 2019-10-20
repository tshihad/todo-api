package data

import (
	"reflect"
	"testing"
	"todo/app/core"
	"todo/app/models"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_repo_GetTodoList(t *testing.T) {
	type args struct {
		userID int
	}
	tests := []struct {
		name    string
		args    args
		given   func(sqlmock.Sqlmock)
		want    []models.TodoList
		wantErr bool
	}{
		{
			name: "Normal test case 1 - get all todo list of user",
			args: args{
				userID: 1,
			},
			given: func(mock sqlmock.Sqlmock) {
				qry := `SELECT * FROM "todo_list"  WHERE ("todo_list"."users_id" = $1)`
				mock.ExpectQuery(core.FixedFullRe(qry)).WithArgs(1).
					WillReturnRows(
						sqlmock.NewRows([]string{
							"id", "todo_name", "users_id",
						}).AddRow(1, "sample", 23),
					)
			},
			want: []models.TodoList{
				{
					ID:     1,
					Name:   "sample",
					UserID: 23,
				},
			},
		},
	}
	mock, db := core.NewSQLMock()
	r := &RepoImp{
		db: db,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.given(mock)
			got, err := r.GetTodoList(tt.args.userID)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.GetTodoList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.GetTodoList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_InsertTodoList(t *testing.T) {
	type args struct {
		todoList models.TodoList
	}
	tests := []struct {
		name    string
		args    args
		given   func(sqlmock.Sqlmock)
		want    models.TodoList
		wantErr bool
	}{
		{
			name: "Normal test case 1",
			args: args{
				todoList: models.TodoList{
					Name:   "test",
					UserID: 23,
				},
			},
			given: func(mock sqlmock.Sqlmock) {
				qry := `INSERT INTO "todo_list" ("todo_name","users_id") VALUES ($1,$2) RETURNING "todo_list"."id"`
				mock.ExpectBegin()
				mock.ExpectQuery(core.FixedFullRe(qry)).WillReturnRows(
					sqlmock.NewRows([]string{
						"id",
					}).AddRow(2),
				)
				mock.ExpectCommit()
			},
			want: models.TodoList{
				ID:     2,
				Name:   "test",
				UserID: 23,
			},
		},
	}
	mock, db := core.NewSQLMock()
	r := &RepoImp{
		db: db,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.given(mock)
			got, err := r.InsertTodoList(tt.args.todoList)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.InsertTodoList() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.InsertTodoList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_UpdateTodoList(t *testing.T) {
	type args struct {
		id   int
		name string
	}
	tests := []struct {
		name    string
		args    args
		given   func(sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "Normal test case 1",
			args: args{1, "changed name"},
			given: func(mock sqlmock.Sqlmock) {
				qry := `UPDATE "todo_list" SET "todo_name" = $1 WHERE ("todo_list"."id" = $2)`
				mock.ExpectBegin()
				mock.ExpectExec(core.FixedFullRe(qry)).WithArgs("changed name", 1).WillReturnResult(
					sqlmock.NewResult(5, 1),
				)
				mock.ExpectCommit()
			},
		},
	}
	mock, db := core.NewSQLMock()
	r := &RepoImp{
		db: db,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.given(mock)
			if err := r.UpdateTodoList(tt.args.id, tt.args.name); (err != nil) != tt.wantErr {
				t.Errorf("repo.UpdateTodoList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_repo_DeleteTodoList(t *testing.T) {
	type args struct {
		todoListID int
	}
	tests := []struct {
		name    string
		args    args
		given   func(sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "Normal test case 1",
			args: args{2},
			given: func(mock sqlmock.Sqlmock) {
				qry := `DELETE FROM "todo_list"  WHERE "todo_list"."id" = $1`
				mock.ExpectBegin()
				mock.ExpectExec(core.FixedFullRe(qry)).WillReturnResult(
					sqlmock.NewResult(2, 1),
				)
				mock.ExpectCommit()
			},
		},
	}
	mock, db := core.NewSQLMock()
	r := &RepoImp{
		db: db,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.given(mock)
			if err := r.DeleteTodoList(tt.args.todoListID); (err != nil) != tt.wantErr {
				t.Errorf("repo.DeleteTodoList() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
