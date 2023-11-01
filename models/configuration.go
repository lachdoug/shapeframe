package models

import (
	"fmt"
	"sf/app"
	"sf/utils"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
)

type Configuration struct {
	Kind   string
	Name   string
	Schema map[any]any
	Values map[string]any
}

func NewConfiguration(
	kind string,
	name string,
	schema map[any]any,
	values map[string]any) (c *Configuration) {
	c = &Configuration{
		Kind:   kind,
		Name:   name,
		Schema: schema,
		Values: values,
	}
	return
}

func (c *Configuration) JsonSchema() (jschema *jsonschema.Schema, err error) {
	schema := utils.JsonMarshal(utils.MapKeyStrings(c.Schema))
	jschema, err = jsonschema.CompileString("schema.json", string(schema))
	return
}

func (c *Configuration) Validate() (err error) {
	var jschema *jsonschema.Schema
	if jschema, err = c.JsonSchema(); err != nil {
		return
	}
	if err = jschema.Validate(c.Values); err != nil {
		switch e := err.(type) {
		case *jsonschema.InfiniteLoopError:
			fmt.Println("configuration validate", "InfiniteLoopError", e)

		case *jsonschema.InvalidJSONTypeError:
			fmt.Println("configuration validate", "InvalidJSONTypeError", e)

		case *jsonschema.LoaderNotFoundError:
			fmt.Println("configuration validate", "LoaderNotFoundError", e)

		case *jsonschema.SchemaError:
			fmt.Println("configuration validate", "SchemaError", e)

		case *jsonschema.ValidationError:
			output := e.DetailedOutput()
			messages := make([]string, len(output.Errors))
			for i, e := range output.Errors {
				messages[i] = strings.Replace(e.Error, "additionalProperties", "additional properties", 1)
			}
			err = app.NewConfigValidationError(c.Kind, c.Name, messages)
		default:
			fmt.Println("configuration validate", "default", e)

		}
	}
	return
}
