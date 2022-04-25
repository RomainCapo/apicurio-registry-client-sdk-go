/*
Apicurio Registry API [v2]

Apicurio Registry is a datastore for standard event schemas and API designs. Apicurio Registry enables developers to manage and share the structure of their data using a REST interface. For example, client applications can dynamically push or pull the latest updates to or from the registry without needing to redeploy. Apicurio Registry also enables developers to create rules that govern how registry content can evolve over time. For example, this includes rules for content validation and version compatibility.  The Apicurio Registry REST API enables client applications to manage the artifacts in the registry. This API provides create, read, update, and delete operations for schema and API artifacts, rules, versions, and metadata.   The supported artifact types include: - Apache Avro schema - AsyncAPI specification - Google protocol buffers - GraphQL schema - JSON Schema - Kafka Connect schema - OpenAPI specification - Web Services Description Language - XML Schema Definition   **Important**: The Apicurio Registry REST API is available from `https://MY-REGISTRY-URL/apis/registry/v2` by default. Therefore you must prefix all API operation paths with `../apis/registry/v2` in this case. For example: `../apis/registry/v2/ids/globalIds/{globalId}`. 

API version: 2.2.4-SNAPSHOT
Contact: apicurio@lists.jboss.org
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package registryclient

import (
	"encoding/json"
	"fmt"
)

// RoleType the model 'RoleType'
type RoleType string

// List of RoleType
const (
	READ_ONLY RoleType = "READ_ONLY"
	DEVELOPER RoleType = "DEVELOPER"
	ADMIN RoleType = "ADMIN"
)

// All allowed values of RoleType enum
var AllowedRoleTypeEnumValues = []RoleType{
	"READ_ONLY",
	"DEVELOPER",
	"ADMIN",
}

func (v *RoleType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RoleType(value)
	for _, existing := range AllowedRoleTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RoleType", value)
}

// NewRoleTypeFromValue returns a pointer to a valid RoleType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRoleTypeFromValue(v string) (*RoleType, error) {
	ev := RoleType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RoleType: valid values are %v", v, AllowedRoleTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RoleType) IsValid() bool {
	for _, existing := range AllowedRoleTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to RoleType value
func (v RoleType) Ptr() *RoleType {
	return &v
}

type NullableRoleType struct {
	value *RoleType
	isSet bool
}

func (v NullableRoleType) Get() *RoleType {
	return v.value
}

func (v *NullableRoleType) Set(val *RoleType) {
	v.value = val
	v.isSet = true
}

func (v NullableRoleType) IsSet() bool {
	return v.isSet
}

func (v *NullableRoleType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoleType(val *RoleType) *NullableRoleType {
	return &NullableRoleType{value: val, isSet: true}
}

func (v NullableRoleType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoleType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

