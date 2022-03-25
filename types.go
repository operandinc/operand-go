package operand

// CollectionSource is an enumeration over various supported collection sources.
type CollectionSource string

// CollectionSource enumerations.
const (
	CollectionSourceNone CollectionSource = "none"
)

// Various metadata types for collections.
type (
	NoneCollectionMetadata struct{}
)

// A Collection is a top-level grouping (of groups).
type Collection struct {
	ID        string           `json:"id"`
	CreatedAt int64            `json:"createdAt"`
	UpdatedAt int64            `json:"updatedAt"`
	Name      string           `json:"name"`
	Source    CollectionSource `json:"source"`
	Metadata  any              `json:"metadata"`
}

// GroupKind is an enumeration over various supported group kinds.
type GroupKind string

// GroupKind enumerations.
const (
	GroupKindDefault GroupKind = "default"
)

// Various metadata types for groups.
type (
	DefaultGroupKindMetadata struct{}
)

// A Group is a logical grouping of atoms.
type Group struct {
	ID           string    `json:"id"`
	CollectionID string    `json:"collectionId"`
	CreatedAt    int64     `json:"createdAt"`
	UpdatedAt    int64     `json:"updatedAt"`
	Name         string    `json:"name"`
	Kind         GroupKind `json:"kind"`
	Metadata     any       `json:"metadata"`
}

// Properties is a set of properties adding additional context to an atom.
type Properties map[string]any

// An Atom is a single unit of information, maps to an embedding.
type Atom struct {
	ID         string     `json:"id"`
	GroupID    string     `json:"groupId"`
	CreatedAt  int64      `json:"createdAt"`
	UpdatedAt  int64      `json:"updatedAt"`
	Content    string     `json:"content"`
	Properties Properties `json:"properties"`
}

// ListResponse is the response from a list operation.
type ListResponse[T any] struct {
	Items []T  `json:"items"`
	More  bool `json:"more"`
}

// DeleteResponse is the response from a delete operation.
type DeleteResponse struct {
	Deleted bool `json:"deleted"`
}
