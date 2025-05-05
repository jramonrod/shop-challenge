package mappers

import "shop-backend/internal/infra/models"

// ConvertAll will call ToDomainObject on every element returning a slice of the resulting
// domain objects.
//
// It defines 2 generic types one of type M that fulfills the models.DomainModel interface
// and another of type T which is our desired output Domain type.
func ConvertAll[M models.DomainModel[T], T any](models []M) []T {
	result := make([]T, len(models))
	for i, m := range models {
		result[i] = m.ToDomainObject()
	}

	return result
}
