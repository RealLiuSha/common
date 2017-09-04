package maps

import (
	"encoding/json"
	"github.com/itchenyi/common/containers"
	"github.com/itchenyi/common/utils"
)

func assertSerializationImplementation() {
	var _ containers.JSONSerializer = (*HashMap)(nil)
	var _ containers.JSONDeserializer = (*HashMap)(nil)
}

// ToJSON outputs the JSON representation of list's elements.
func (m *HashMap) ToJSON() ([]byte, error) {
	elements := make(map[string]interface{})
	for key, value := range m.m {
		elements[utils.ToString(key)] = value
	}
	return json.Marshal(&elements)
}

// FromJSON populates list's elements from the input JSON representation.
func (m *HashMap) FromJSON(data []byte) error {
	elements := make(map[string]interface{})
	err := json.Unmarshal(data, &elements)
	if err == nil {
		m.Clear()
		for key, value := range elements {
			m.m[key] = value
		}
	}
	return err
}