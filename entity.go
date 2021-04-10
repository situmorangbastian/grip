package grip

// ExampleEntity holds information of example entity.
type ExampleEntity struct {
	ID string `json:"id"`
}

// ExampleEdge holds information of example edge.
type ExampleEdge struct {
	Cursor string
	Node   ExampleEntity
}

// ExampleResult holds information of example result.
type ExampleResult struct {
	PageInfo
	TotalCount int
	Edges      []ExampleEdge
}

// PageInfo holds information of page info.
type PageInfo struct {
	EndCursor   string
	HasNextPage bool
}
