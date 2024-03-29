// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// BenchesColumns holds the columns for the "benches" table.
	BenchesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "os", Type: field.TypeString},
		{Name: "arch", Type: field.TypeString},
		{Name: "cpu", Type: field.TypeString},
		{Name: "package", Type: field.TypeString},
		{Name: "pass", Type: field.TypeBool},
	}
	// BenchesTable holds the schema information for the "benches" table.
	BenchesTable = &schema.Table{
		Name:       "benches",
		Columns:    BenchesColumns,
		PrimaryKey: []*schema.Column{BenchesColumns[0]},
	}
	// BenchResultsColumns holds the columns for the "bench_results" table.
	BenchResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "n", Type: field.TypeInt64},
		{Name: "ns_per_op", Type: field.TypeFloat64},
		{Name: "alloced_bytes_per_op", Type: field.TypeInt64},
		{Name: "allocs_per_op", Type: field.TypeInt64},
		{Name: "mb_per_s", Type: field.TypeFloat64},
		{Name: "measured", Type: field.TypeInt64},
		{Name: "ord", Type: field.TypeInt64},
		{Name: "bench_bench_result", Type: field.TypeInt, Nullable: true},
	}
	// BenchResultsTable holds the schema information for the "bench_results" table.
	BenchResultsTable = &schema.Table{
		Name:       "bench_results",
		Columns:    BenchResultsColumns,
		PrimaryKey: []*schema.Column{BenchResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "bench_results_benches_bench_result",
				Columns:    []*schema.Column{BenchResultsColumns[9]},
				RefColumns: []*schema.Column{BenchesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		BenchesTable,
		BenchResultsTable,
	}
)

func init() {
	BenchResultsTable.ForeignKeys[0].RefTable = BenchesTable
}
