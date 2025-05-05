package models

// DomainModel is a model that maps to a domain entity. This is useful to help us map
// from database models to domain level entities.
type DomainModel[T any] interface {
	ToDomainObject() T
}
