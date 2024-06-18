## Ini File Parser

* Parser(): Ini - parse ini file into a object containing it's properties
  * has two states
    * ReadSection
    * ReadKeyValue
* GetProperty(sectionName , key , value ): string - get a property of a scection eg. [ general ]
* SetProperty(sectionName , key , value ) - set a property of a section
* DelProperty(sectionName,key) - delete a property of a section
* HasProprty(sectionName,key): bool - check if property exists or not

## Ini File

Ini file is a simple configuration file that contains sections and properties , same as `.yaml` file

* property is represented in key: value format
* the properties can be grouped into a section
