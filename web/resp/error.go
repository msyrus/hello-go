package resp

// Error represents a response object of api error
type Error struct {
	ID         string                 `json:"id,omitempty"`
	Message    string                 `json:"message,omitempty"`
	Details    map[string]interface{} `json:"details,omitempty"`
	StackTrace string                 `json:"stackTrace,omitempty"`
}
