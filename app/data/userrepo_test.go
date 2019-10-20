package data

import (
	"reflect"
	"testing"
	"todo/app/core"
	"todo/app/models"

	"github.com/DATA-DOG/go-sqlmock"
)

func Test_repo_InsertUser(t *testing.T) {
	type args struct {
		user models.Users
	}
	tests := []struct {
		name    string
		args    args
		given   func(sqlmock.Sqlmock)
		want    int
		wantErr bool
	}{
		{
			name: "Normal test case 1 - succesfully inserted user",
			args: args{
				user: models.Users{
					Email: "test@test.co",
					Name:  "test1",
				},
			},
			given: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				qry := `INSERT INTO "users" ("users_name","email") VALUES ($1,$2) RETURNING "users"."id"`
				mock.ExpectQuery(core.FixedFullRe(qry)).
					WithArgs("test1", "test@test.co").
					WillReturnRows(
						sqlmock.NewRows([]string{"id"}).AddRow(0),
					)
				mock.ExpectCommit()
			},
			want: 0,
		},
	}
	mock, db := core.NewSQLMock()
	r := &RepoImp{
		db: db,
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.given(mock)
			got, err := r.InsertUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.InsertUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("repo.InsertUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_repo_GetUser(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		given   func(sqlmock.Sqlmock)
		want    models.Users
		wantErr bool
	}{
		{
			name: "Normal test case 1 - succesfully get a user",
			args: args{
				name: "test",
			},
			given: func(mock sqlmock.Sqlmock) {
				qry := `SELECT * FROM "users" WHERE ("users"."users_name" = $1) ORDER BY "users"."id" ASC LIMIT 1`
				mock.ExpectQuery(core.FixedFullRe(qry)).WithArgs("test").
					WillReturnRows(
						sqlmock.NewRows([]string{
							"id", "users_name", "email",
						}).AddRow(1, "test1", "test@test.co"),
					)
			},
			want: models.Users{
				Email: "test@test.co",
				ID:    1,
				Name:  "test1",
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
			got, err := r.GetUser(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("repo.GetUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("repo.GetUser() = %v, want %v", got, tt.want)
			}
		})
	}
}
