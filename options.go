package zipkinsql

// TraceOption allows for managing zipkinsql configuration using functional options.
type TraceOption func(o *TraceOptions)

// TraceOptions holds configuration of our zipkinsql tracing middleware.
// By default all boolean options are set to false intentionally when creating
// a wrapped driver and provide the most sensible default with both performance
// and security in mind.
type TraceOptions struct {
	// LastInsertIDSpan, if set to true, will enable the creation of spans on
	// LastInsertId calls.
	LastInsertIDSpan bool

	// RowsAffectedSpan, if set to true, will enable the creation of spans on
	// RowsAffectedSpan calls.
	RowsAffectedSpan bool

	// TagQuery, if set to true, will enable recording of sql queries in spans.
	// Only allow this if it is safe to have queries recorded with respect to
	// security.
	TagQuery bool

	TagAffectedRows bool

	// DefaultTags will be set to each span as default.
	DefaultTags map[string]string
}

// WithAllTraceOptions enables all available trace options.
func WithAllTraceOptions() TraceOption {
	return func(o *TraceOptions) {
		*o = AllTraceOptions
	}
}

// AllTraceOptions has all tracing options enabled.
var AllTraceOptions = TraceOptions{
	RowsAffectedSpan: true,
	LastInsertIDSpan: true,
	TagQuery:         true,
	TagAffectedRows:  true,
}

// WithOptions sets the zipkinsql tracing middleware options through a single
// TraceOptions object.
func WithOptions(options TraceOptions) TraceOption {
	return func(o *TraceOptions) {
		*o = options
	}
}

// WithRowsAffectedSpan if set to true, will enable the creation of spans on
// RowsAffected calls.
func WithRowsAffectedSpan(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.RowsAffectedSpan = b
	}
}

// WithLastInsertIDSpan if set to true, will enable the creation of spans on
// LastInsertId calls.
func WithLastInsertIDSpan(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.LastInsertIDSpan = b
	}
}

// WithTagQuery if set to true, will enable recording of SQL queries in spans.
// Only allow this if it is safe to have queries recorded with respect to
// security.
func WithTagQuery(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.TagQuery = b
	}
}

// WithTagAffectedRows if set to true, will enable recording of the affected rows
// number in spans.
func WithTagAffectedRows(b bool) TraceOption {
	return func(o *TraceOptions) {
		o.TagAffectedRows = b
	}
}

// WithDefaultTags will be set to each span as default.
func WithDefaultTags(tags map[string]string) TraceOption {
	return func(o *TraceOptions) {
		o.DefaultTags = tags
	}
}
