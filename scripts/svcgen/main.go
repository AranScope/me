package main

import (
	. "github.com/dave/jennifer/jen"
	"github.com/iancoleman/strcase"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type TemplateFiller struct {
	Path             string
	CamelCasePath    string
	LetterCasePath   string
	LetterCaseEntity string
	LowerCaseEntity  string
	RequestBody      string
	ResponseBody     string
}

type Param struct {
	Name  string  `yaml:"name"`
	Type  string  `yaml:"type"`
	Items []Param `yaml:"items"`
}

type Endpoint struct {
	Path       string  `yaml:"path"`
	Idempotent bool    `yaml:"idempotent"`
	Request    []Param `yaml:"request"`
	Response   []Param `yaml:"response"`
}

type ServiceDefinition struct {
	Endpoints []Endpoint `yaml:"endpoints"`
}

func main() {

	svcDefBytes, err := ioutil.ReadFile("/Users/aran/go/src/github.com/AranScope/me/scripts/svcgen/service-definition.yml")
	if err != nil {
		panic(err)
	}

	svcDef := ServiceDefinition{}

	err = yaml.Unmarshal(svcDefBytes, &svcDef)
	if err != nil {
		panic(err)
	}

	out, err := os.Create("/Users/aran/go/src/github.com/AranScope/me/scripts/svcgen/generated.go")
	if err != nil {
		panic(err)
	}

	headerTmpl, err := ioutil.ReadFile("/Users/aran/go/src/github.com/AranScope/me/scripts/svcgen/header.go.tmpl")
	if err != nil {
		panic(err)
	}

	_, err = out.Write(headerTmpl)

	bodyTmpl, err := template.ParseFiles("/Users/aran/go/src/github.com/AranScope/me/scripts/svcgen/body.go.tmpl")
	if err != nil {
		panic(err)
	}

	for _, endpoint := range svcDef.Endpoints {
		path := endpoint.Path
		entity := strings.Split(endpoint.Path, "_")[1]

		//paramsToStructs(endpoint.Request, strcase.ToCamel(entity)+"Request").GoString()

		responseBody := ""
		for _, s := range paramsToStructs(endpoint.Response, strcase.ToCamel(entity)) {
			responseBody += s.GoString()
		}

		err = bodyTmpl.Execute(out, TemplateFiller{
			Path:             path,
			CamelCasePath:    strcase.ToLowerCamel(path),
			LetterCasePath:   strcase.ToCamel(path),
			LetterCaseEntity: strcase.ToCamel(entity),
			LowerCaseEntity:  strings.ToLower(entity),
			//RequestBody:      requestBody,
			ResponseBody:     responseBody,
		})
		if err != nil {
			panic(err)
		}
	}
}

func paramsToStructs(params []Param, name string) []*Statement {

	var structs []*Statement
	var structFields []Code

	for _, param := range params {
		structField := &Statement{}

		switch param.Type {
		// todo: fill in the rest of these
		case "string":
			structField = Id(strcase.ToCamel(param.Name))
			structField = structField.String()
			structField = structField.Tag(map[string]string{"json": strcase.ToSnake(param.Name)})
			break
		case "integer":
			structField = Id(strcase.ToCamel(param.Name))
			structField = structField.Int()
			structField = structField.Tag(map[string]string{"json": strcase.ToSnake(param.Name)})

			break
		case "array":
			nestedStructs := paramsToStructs(param.Items, param.Name)
			structField = Id(param.Name).List(nestedStructs[len(nestedStructs)-1])
			structs = append(structs, nestedStructs...)
		}
		structFields = append(structFields, structField)
	}

	requestBody := Type().Id(name).Struct(
		structFields...,
	)

	structs = append(structs, requestBody)

	return structs
}
