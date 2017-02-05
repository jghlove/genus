package genus

import (
	"reflect"
	"testing"
)

func TestTemplate_loadFile(t *testing.T) {
	type fields struct {
		Name        string
		Source      string
		TargetDir   string
		Filename    string
		SkipExists  bool
		SkipFormat  bool
		rawTemplate []byte
	}
	tests := []struct {
		name     string
		fields   fields
		wantData []byte
		wantErr  bool
	}{
		{"OK", fields{
			Name:      "t1",
			Source:    "./testdata/template/t1.tpl",
			TargetDir: "_test",
			Filename:  "t1.go",
		}, []byte("func T1(){}\n"), false},
		{"KO - missing file", fields{
			Name:      "t1",
			Source:    "./testdata/template/not-exist-t1.tpl",
			TargetDir: "_test",
			Filename:  "t1.go",
		}, nil, true},
		{"KO - blank template data", fields{
			Name:      "t1",
			TargetDir: "_test",
			Filename:  "t1.go",
		}, nil, true},
	}
	for _, tt := range tests {
		tmpl := &Template{
			Name:        tt.fields.Name,
			Source:      tt.fields.Source,
			TargetDir:   tt.fields.TargetDir,
			Filename:    tt.fields.Filename,
			SkipExists:  tt.fields.SkipExists,
			SkipFormat:  tt.fields.SkipFormat,
			rawTemplate: tt.fields.rawTemplate,
		}
		gotData, err := tmpl.loadFile()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Template.loadFile() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(gotData, tt.wantData) {
			t.Errorf("%q. Template.loadFile() = %v, want %v", tt.name, string(gotData), string(tt.wantData))
		}
	}
}

func TestTemplate_render(t *testing.T) {
	type fields struct {
		Name        string
		Source      string
		TargetDir   string
		Filename    string
		SkipExists  bool
		SkipFormat  bool
		rawTemplate []byte
	}
	type args struct {
		context interface{}
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData []byte
		wantErr  bool
	}{
		{"OK", fields{
			rawTemplate: []byte("type {{ .Name }} struct{}"),
		}, args{map[string]string{"Name": "A", "Package": "p1"}}, []byte("package p1\n\n\ntype A struct{}"), false},
		{"KO", fields{
			rawTemplate: []byte("type {{ .Name"),
		}, args{map[string]string{"Name": "A"}}, nil, true},
	}
	for _, tt := range tests {
		tmpl := &Template{
			Name:        tt.fields.Name,
			Source:      tt.fields.Source,
			TargetDir:   tt.fields.TargetDir,
			Filename:    tt.fields.Filename,
			SkipExists:  tt.fields.SkipExists,
			SkipFormat:  tt.fields.SkipFormat,
			rawTemplate: tt.fields.rawTemplate,
		}
		gotData, err := tmpl.render(tt.args.context)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Template.render() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(gotData, tt.wantData) {
			t.Errorf("%q. Template.render() = %v, want %v", tt.name, string(gotData), string(tt.wantData))
		}
	}
}

func TestTemplate_SetRawTemplate(t *testing.T) {
	type fields struct {
		Name        string
		Source      string
		TargetDir   string
		Filename    string
		SkipExists  bool
		SkipFormat  bool
		rawTemplate []byte
		rawResult   []byte
	}
	type args struct {
		raw []byte
	}
	tests := []struct {
		name     string
		fields   fields
		args     args
		wantData []byte
	}{
		{"OK", fields{}, args{[]byte("abc")}, []byte("abc")},
	}
	for _, tt := range tests {
		tmpl := &Template{
			Name:        tt.fields.Name,
			Source:      tt.fields.Source,
			TargetDir:   tt.fields.TargetDir,
			Filename:    tt.fields.Filename,
			SkipExists:  tt.fields.SkipExists,
			SkipFormat:  tt.fields.SkipFormat,
			rawTemplate: tt.fields.rawTemplate,
			rawResult:   tt.fields.rawResult,
		}
		if gotData := tmpl.SetRawTemplate(tt.args.raw); !reflect.DeepEqual(gotData, tt.wantData) {
			t.Errorf("%q. Template.SetRawTemplate() = %v, want %v", tt.name, gotData, tt.wantData)
		}
	}
}

func TestTemplate_load(t *testing.T) {
	type fields struct {
		Name        string
		Source      string
		TargetDir   string
		Filename    string
		SkipExists  bool
		SkipFormat  bool
		rawTemplate []byte
		rawResult   []byte
	}
	tests := []struct {
		name     string
		fields   fields
		wantData []byte
		wantErr  bool
	}{
		{"OK", fields{
			Name:      "t1",
			Source:    "./testdata/template/t1.tpl",
			TargetDir: "_test",
			Filename:  "t1.go",
		}, []byte("func T1(){}\n"), false},
		{"OK", fields{
			Name:        "t1",
			Source:      "./testdata/template/t1.tpl",
			TargetDir:   "_test",
			Filename:    "t1.go",
			rawTemplate: []byte("abc"),
		}, []byte("abc"), false},
	}
	for _, tt := range tests {
		tmpl := &Template{
			Name:        tt.fields.Name,
			Source:      tt.fields.Source,
			TargetDir:   tt.fields.TargetDir,
			Filename:    tt.fields.Filename,
			SkipExists:  tt.fields.SkipExists,
			SkipFormat:  tt.fields.SkipFormat,
			rawTemplate: tt.fields.rawTemplate,
			rawResult:   tt.fields.rawResult,
		}
		gotData, err := tmpl.load()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Template.load() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(gotData, tt.wantData) {
			t.Errorf("%q. Template.load() = %v, want %v", tt.name, gotData, tt.wantData)
		}
	}
}

func TestTemplate_Render(t *testing.T) {
	type fields struct {
		Name        string
		Source      string
		TargetDir   string
		Filename    string
		SkipExists  bool
		SkipFormat  bool
		rawTemplate []byte
		rawResult   []byte
	}
	type args struct {
		data interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantResult []byte
		wantErr    bool
	}{
		{"OK - load template from file", fields{
			Name:      "t3",
			Source:    "./testdata/template/t3.tpl",
			TargetDir: "_test",
			Filename:  "t3.go",
		}, args{map[string]string{"Package": "p1", "Name": "A"}}, []byte("package p1\n\ntype A struct{}\n"), false},
		{"OK - set raw template", fields{
			rawTemplate: []byte("type {{ .Name }} struct{}"),
		}, args{map[string]string{"Package": "p1", "Name": "A"}}, []byte("package p1\n\ntype A struct{}\n"), false},
		{"KO", fields{
			Name:      "t3",
			TargetDir: "_test",
			Filename:  "t3.go",
		}, args{map[string]string{"Package": "main"}}, nil, true},
		{"KO", fields{
			rawTemplate: []byte("package {{ .Package"),
		}, args{map[string]string{"Package": "main"}}, nil, true},
		{"KO", fields{
			Name:      "t4",
			Source:    "./testdata/template/t4.tpl",
			TargetDir: "_test",
			Filename:  "t4.go",
		}, args{map[string]string{"Package": "main"}}, nil, true},
	}
	for _, tt := range tests {
		tmpl := &Template{
			Name:        tt.fields.Name,
			Source:      tt.fields.Source,
			TargetDir:   tt.fields.TargetDir,
			Filename:    tt.fields.Filename,
			SkipExists:  tt.fields.SkipExists,
			SkipFormat:  tt.fields.SkipFormat,
			rawTemplate: tt.fields.rawTemplate,
			rawResult:   tt.fields.rawResult,
		}
		gotResult, err := tmpl.Render(tt.args.data)
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Template.Render() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(gotResult, tt.wantResult) {
			t.Errorf("%q. Template.Render() = %v, want %v", tt.name, string(gotResult), string(tt.wantResult))
		}
	}
}

func TestTemplate_format(t *testing.T) {
	type fields struct {
		Name        string
		Source      string
		TargetDir   string
		Filename    string
		SkipExists  bool
		SkipFormat  bool
		rawTemplate []byte
		rawResult   []byte
	}
	tests := []struct {
		name     string
		fields   fields
		wantData []byte
		wantErr  bool
	}{
		{"OK", fields{
			rawResult: []byte("package main\nfunc T1(){}\n"),
		}, []byte("package main\n\nfunc T1() {}\n"), false},
		{"KO - bad syntax", fields{
			rawResult: []byte("xxx main\nfunc T1(){}\n"),
		}, nil, true},
		{"KO - bad syntax but skip format", fields{
			SkipFormat: true,
			rawResult:  []byte("xxx main\nfunc T1(){}\n"),
		}, []byte("xxx main\nfunc T1(){}\n"), false},
	}
	for _, tt := range tests {
		tmpl := &Template{
			Name:        tt.fields.Name,
			Source:      tt.fields.Source,
			TargetDir:   tt.fields.TargetDir,
			Filename:    tt.fields.Filename,
			SkipExists:  tt.fields.SkipExists,
			SkipFormat:  tt.fields.SkipFormat,
			rawTemplate: tt.fields.rawTemplate,
			rawResult:   tt.fields.rawResult,
		}
		gotData, err := tmpl.format()
		if (err != nil) != tt.wantErr {
			t.Errorf("%q. Template.format() error = %v, wantErr %v", tt.name, err, tt.wantErr)
			continue
		}
		if !reflect.DeepEqual(gotData, tt.wantData) {
			t.Errorf("%q. Template.format() = %v, want %v", tt.name, gotData, tt.wantData)
		}
	}
}

func TestTemplate_write(t *testing.T) {
	type fields struct {
		Name        string
		Source      string
		TargetDir   string
		Filename    string
		SkipExists  bool
		SkipFormat  bool
		rawTemplate []byte
		rawResult   []byte
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{"OK", fields{
			TargetDir: "./_test",
			Filename:  "t1.go",
			rawResult: []byte("abc"),
		}, false},
		{"OK", fields{
			TargetDir:  "./_test",
			Filename:   "t1.go",
			SkipExists: true,
			rawResult:  []byte("abcd"),
		}, false},
	}
	for _, tt := range tests {
		tmpl := &Template{
			Name:        tt.fields.Name,
			Source:      tt.fields.Source,
			TargetDir:   tt.fields.TargetDir,
			Filename:    tt.fields.Filename,
			SkipExists:  tt.fields.SkipExists,
			SkipFormat:  tt.fields.SkipFormat,
			rawTemplate: tt.fields.rawTemplate,
			rawResult:   tt.fields.rawResult,
		}
		if err := tmpl.write(); (err != nil) != tt.wantErr {
			t.Errorf("%q. Template.write() error = %v, wantErr %v", tt.name, err, tt.wantErr)
		}
	}
}
