// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package client

import (
	"fmt"
	"io"
	"strconv"
)

type Bench struct {
	ID      string         `json:"id"`
	Os      string         `json:"os"`
	Arch    string         `json:"arch"`
	CPU     string         `json:"cpu"`
	Package string         `json:"package"`
	Pass    bool           `json:"pass"`
	Results []*BenchResult `json:"results"`
}

func (Bench) IsNode() {}

// A connection to a list of items.
type BenchConnection struct {
	// A list of edges.
	Edges []*BenchEdge `json:"edges"`
	// Information to aid in pagination.
	PageInfo PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// An edge in a connection.
type BenchEdge struct {
	// The item at the end of the edge.
	Node *Bench `json:"node"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

type BenchResult struct {
	ID                string  `json:"id"`
	Name              string  `json:"name"`
	N                 int64   `json:"n"`
	NsPerOp           float64 `json:"nsperop"`
	AllocedBytesPerOp int64   `json:"allocedbytesperop"`
	AllocsPerOp       int64   `json:"allocsperop"`
	MBPerS            float64 `json:"mbpers"`
	Measured          int64   `json:"measured"`
	Ord               int64   `json:"ord"`
}

func (BenchResult) IsNode() {}

// A connection to a list of items.
type BenchResultConnection struct {
	// A list of edges.
	Edges []*BenchResultEdge `json:"edges"`
	// Information to aid in pagination.
	PageInfo PageInfo `json:"pageInfo"`
	// Identifies the total count of items in the connection.
	TotalCount int64 `json:"totalCount"`
}

// An edge in a connection.
type BenchResultEdge struct {
	// The item at the end of the edge.
	Node *BenchResult `json:"node"`
	// A cursor for use in pagination.
	Cursor string `json:"cursor"`
}

// CreateBenchInput is used for create Bench object.
// Input was generated by ent.
type CreateBenchInput struct {
	Os        string   `json:"os"`
	Arch      string   `json:"arch"`
	CPU       string   `json:"cpu"`
	Package   string   `json:"package"`
	Pass      bool     `json:"pass"`
	ResultIDs []string `json:"resultIDs,omitempty"`
}

// CreateBenchResultInput is used for create BenchResult object.
// Input was generated by ent.
type CreateBenchResultInput struct {
	Name              string  `json:"name"`
	N                 int64   `json:"n"`
	Nsperop           float64 `json:"nsperop"`
	Allocedbytesperop int64   `json:"allocedbytesperop"`
	Allocsperop       int64   `json:"allocsperop"`
	Mbpers            float64 `json:"mbpers"`
	Measured          int64   `json:"measured"`
	Ord               int64   `json:"ord"`
}

// Information about pagination in a connection.
// https://relay.dev/graphql/connections.htm#sec-undefined.PageInfo
type PageInfo struct {
	// When paginating forwards, are there more items?
	HasNextPage bool `json:"hasNextPage"`
	// When paginating backwards, are there more items?
	HasPreviousPage bool `json:"hasPreviousPage"`
	// When paginating backwards, the cursor to continue.
	StartCursor *string `json:"startCursor"`
	// When paginating forwards, the cursor to continue.
	EndCursor *string `json:"endCursor"`
}

// Possible directions in which to order a list of items when provided an `orderBy` argument.
type OrderDirection string

const (
	// Specifies an ascending order for a given `orderBy` argument.
	OrderDirectionAsc OrderDirection = "ASC"
	// Specifies a descending order for a given `orderBy` argument.
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
