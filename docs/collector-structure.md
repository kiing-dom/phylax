## what should it do?

- collect spans
- store them
  - rn in memory
  - in production, probably a cassandra/elasticsearch etc.
- later query spans

### service send
e.g. service send
```json
{
    "trace_id": "123",
    "span_id": "A123",
    "parent_span_id": "A123",
    "service": "service-a",
    "operation": "GET /users",
    "start_time": "22:00",
    "duration": "555555"
}
```

## in-memory storage model

```go
map[TraceId][]Span
`e.g. 123 -> [span1, span2, span3]`
```