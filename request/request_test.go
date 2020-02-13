package request

import (
	"reflect"
	"testing"
)

func TestArticleRequest_Validate(t *testing.T) {
	type fields struct {
		Title string
		Body  string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{"required title", fields{"", "body"}, []string{"Titleは必須項目です"}},
		{"required body", fields{"title", ""}, []string{"Bodyは必須項目です"}},
		{"gte title", fields{"ti", "body"}, []string{"Titleは5文字以上で入力してください"}},
		{"lt title", fields{"titletitletitletitlet", "body"}, []string{"Titleは20文字未満で入力してください"}},
		{"passed validation", fields{"title", "b"}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &ArticleRequest{
				Title: tt.fields.Title,
				Body:  tt.fields.Body,
			}
			if got := req.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArticleRequest.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCreateArticleRequest_Validate(t *testing.T) {
	type fields struct {
		ArticleRequest *ArticleRequest
		AuthorRequest  *AuthorRequest
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &CreateArticleRequest{
				ArticleRequest: tt.fields.ArticleRequest,
				AuthorRequest:  tt.fields.AuthorRequest,
			}
			if got := req.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateArticleRequest.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthorRequest_Validate(t *testing.T) {
	type fields struct {
		AuthorName string
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := &AuthorRequest{
				AuthorName: tt.fields.AuthorName,
			}
			if got := req.Validate(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AuthorRequest.Validate() = %v, want %v", got, tt.want)
			}
		})
	}
}
