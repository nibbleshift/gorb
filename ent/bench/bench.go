// Code generated by ent, DO NOT EDIT.

package bench

const (
	// Label holds the string label denoting the bench type in the database.
	Label = "bench"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// Table holds the table name of the bench in the database.
	Table = "benches"
)

// Columns holds all SQL columns for bench fields.
var Columns = []string{
	FieldID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
