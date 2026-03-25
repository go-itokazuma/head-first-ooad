package main

type InstrumentSpec struct {
	properties map[string]interface{}
}

func NewInstrumentSpec(properties map[string]interface{}) *InstrumentSpec {
	if properties == nil {
		return &InstrumentSpec{
			properties: make(map[string]interface{}),
		}
	}
	copiedProps := make(map[string]interface{})
	for k, v := range properties {
		copiedProps[k] = v
	}
	return &InstrumentSpec{
		properties: copiedProps,
	}
}

func (is *InstrumentSpec) GetProperty(propertyName string) interface{} {
	return is.properties[propertyName]
}

func (is *InstrumentSpec) GetProperties() map[string]interface{} {
	return is.properties
}

func (is *InstrumentSpec) Matches(otherSpec *InstrumentSpec) bool {
	for propertyName, otherValue := range otherSpec.GetProperties() {
		if is.properties[propertyName] != otherValue {
			return false
		}
	}
	return true
}
