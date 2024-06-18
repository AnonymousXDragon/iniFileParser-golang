package tools

type Ini struct {
	sections map[string]Section
}

func NewIni() *Ini {
	return &Ini{
		sections: make(map[string]Section),
	}
}

//[ set , get , del ]//

func (i *Ini) SetSection(name string, section Section) {
	i.sections[name] = section
}

func (i *Ini) GetSection(name string) Section {
	_, ok := i.sections[name]
	if !ok {
		panic("invalid section name found")
	}

	return i.sections[name]
}

func (i *Ini) DeleteSection(name string) {
	delete(i.sections, name)
}

//[ properties: set , get , delete , has ]//

func (i *Ini) SetProperty(sectionName string, key string, value string) {
	s := i.sections[sectionName]
	s.SetProperty(key, value)
}

func (i *Ini) GetProperty(sectionName string, key string) string {
	_, sec_ok := i.sections[sectionName]
	_, prop_ok := i.sections[sectionName].properties[key]
	if !sec_ok {
		panic("invalid section name")
	}

	if !prop_ok {
		panic("invalid property name ")
	}

	return i.sections[sectionName].properties[key]
}

func (i *Ini) DeleteProperty(sectionName string, key string) {
	delete(i.sections[sectionName].properties, key)
}

func (i *Ini) HasProperty(sectionName string, key string) bool {
	_, ok := i.sections[sectionName].properties[key]
	return ok
}

func (i *Ini) ToString() string {
	output := ""
	sep := "="
	for key, value := range i.sections {
		output += "[ " + key + " ]" + "\n"
		for key, value := range value.properties {
			output += key + " " + sep + " " + value + "\n"
		}
		output += "\n"
	}
	return output
}
