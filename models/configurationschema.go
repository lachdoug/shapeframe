package models

// import (
// 	"fmt"
// 	"sf/app"
// )

// type ConfigurationSchema struct {
// 	Form       []*FormComponentSchema
// 	Properties []*FormPropertySchema
// }

// func (cs *ConfigurationSchema) validation(id string, vn *validations.Validation) (err error) {
// 	if cs.Properties != nil {
// 		for i, csp := range cs.Properties {
// 			csp.validation(fmt.Sprintf("%s properties[%d]", id, i), vn)
// 		}
// 	}
// 	if cs.Form != nil {
// 		for i, fmc := range cs.Form {
// 			fmc.validation(fmt.Sprintf("%s form[%d]", id, i), vn)
// 		}
// 	}
// 	return
// }
