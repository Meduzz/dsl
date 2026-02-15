package schema

type (
	Kind    string
	SubType string // TODO should this be a struct? <- 1.

	Field struct {
		Name    string  `json:"name"`
		Kind    Kind    `json:"kind"`
		SubType SubType `json:"subtype,omitempty"`
		Default any     `json:"default,omitempty"`
		Format  string  `json:"format,omitempty"`
		Id      bool    `json:"id,omitempty"`
		// TODO form hint field multiselt, select, radio, checkbox
		// TODO options field for enums
		// TODO number or rows in the textarea for longtext?
		// TODO field to indicate that this field can be label when entity is used in dropdown (or value on reference subtype)
		// TODO field to indicate that the value is generated elsewhere (like uuid, hashes etc.)
		// TODO field to indicate type of array?
		// TODO field to indicate number of decimals in decimal subtype?
		// TODO field to indicate if the field is required or not
		// TODO numbers min, max & interval?
	}
)

const (
	Text    = Kind("text")
	Number  = Kind("number")
	Boolean = Kind("boolean")
	Array   = Kind("array") // TODO borderline
	// If I add object here, I'll have to support like oneOfs and all that jazz.

	Longtext = SubType("longtext")
	Date     = SubType("date")
	Datetime = SubType("datetime")
	Enum     = SubType("enum")
	Tag      = SubType("tag")
	Colour   = SubType("colour")
	Email    = SubType("email")
	Url      = SubType("url")
	UUID     = SubType("uuid")
	Ref      = SubType("reference")
	Decimal  = SubType("decimal")
	Integer  = SubType("integer")
)
