# JSON Struct

Tasks:
1. Define Person struct with fields: ID (int), Name (string), Email (string, omit if empty), CreatedAt (time.Time) with custom RFC3339 format using MarshalJSON / UnmarshalJSON or custom type.
2. Demonstrate marshaling to JSON and back with error handling.
3. Show use of omitempty and `-` tag.
4. Add test covering round trip and empty Email omission.

Stretch:
- Add custom type Email that validates format.
- Stream encode slice of Person to file with Encoder.
