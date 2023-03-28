package entities

type (
	// Role type for application role
	Role string
	// Roles array of Role
	Roles []Role
)

const (
	// ReadRole role for reader
	ReadRole Role = "read"
	// WriteRole role for write
	WriteRole Role = "write"
)

// PossibleRoles application possible roles
var PossibleRoles = Roles{ReadRole, WriteRole}

// Contains returns a true bool if a Role exists in Roles
func (roles Roles) Contains(r Role) bool {
	for _, role := range roles {
		if role == r {
			return true
		}
	}

	return false
}
