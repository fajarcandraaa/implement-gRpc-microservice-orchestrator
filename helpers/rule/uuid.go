package rule

import (
	"regexp"

	"github.com/fajarcandraaa/implement-gRpc-microservice-orchestrator/helpers/unique"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type Rule func(value interface{}) error

// IsUUID to validate payload is UUID
func IsUUID() Rule {
	fn := func(value interface{}) error {
		switch v := value.(type) {
		case string:
			return unique.ValidateUUID(v)
		case *string:
			return unique.ValidateUUID(*v)
		default:
			return unique.ErrInvalidUUID
		}
	}

	return fn
}

func (U Rule) Validate(value interface{}) error {
	return U(value)
}

const regexURNString = "^urn:[a-z0-9][a-z0-9-]{0,31}:[a-z0-9()+,\\-.:=@;$_!*'%/?#]+$"

var regexURN = regexp.MustCompile(regexURNString)

// IsURI to validate payload is Uri
func IsURI() validation.Rule {
	return Rule(func(value interface{}) error {
		var s string
		switch v := value.(type) {
		default:
			return nil
		case string:
			s = v
		case *string:
			s = *v
		}

		if s == "" {
			return nil
		}

		return nil
	})
}
