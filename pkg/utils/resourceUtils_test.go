package utils

import (
	"testing"
)

func TestResourceHandler_CreateProjectResources(t *testing.T) {
	type fields struct {
		BaseURL string
	}
	type args struct {
		project   string
		resources []*Resource
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "create resource",
			fields: fields{
				BaseURL: "localhost:8080",
			},
			args: args{
				project: "rockshop",
				resources: []*Resource{
					&Resource{
						ResourceURI:     "shipyard-hello2.yaml",
						ResourceContent: "foo",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ResourceHandler{
				BaseURL: tt.fields.BaseURL,
			}
			got, err := r.CreateProjectResources(tt.args.project, tt.args.resources)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceHandler.CreateProjectResources() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) <= 0 {
				t.Errorf("Got empty response")
			}
		})
	}
}

func TestResourceHandler_UpdateProjectResource(t *testing.T) {
	type fields struct {
		BaseURL string
	}
	type args struct {
		project  string
		resource *Resource
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "update resource",
			fields: fields{
				BaseURL: "localhost:8080",
			},
			wantErr: false,
			args: args{
				project: "rockshop",
				resource: &Resource{
					ResourceURI:     "shipyard-tests.yaml",
					ResourceContent: "this is a test!",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ResourceHandler{
				BaseURL: tt.fields.BaseURL,
			}
			got, err := r.UpdateProjectResource(tt.args.project, tt.args.resource)
			if (err != nil) != tt.wantErr {
				t.Errorf("ResourceHandler.UpdateProjectResource() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ResourceHandler.UpdateProjectResource() = %v, want %v", got, tt.want)
			}
		})
	}
}