package ddtags_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/webdestroya/ddtags"
)

func TestValidator(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		require.NoError(t, ddtags.Validate(&simpleTagStruct{}))
		require.NoError(t, ddtags.Validate(&fullTagStruct{}))

		require.NoError(t, ddtags.Validate(&struct {
			StrPtr *string `ddtag:"strptr"`
		}{}))

		require.NoError(t, ddtags.Validate(&struct {
			Thing map[string]string `ddtag:""`
		}{}))

		require.NoError(t, ddtags.Validate(&struct {
			Thing map[string]string `ddtag:"-"`
		}{}))
	})

	t.Run("invalid", func(t *testing.T) {
		require.ErrorContains(t, ddtags.Validate(&struct {
			StrPtr *string           `ddtag:"strptr"`
			Thing  map[string]string `ddtag:"someval"`
		}{}), "invalid field type: Thing")

		require.ErrorContains(t, ddtags.Validate(&struct {
			StrPtr *string       `ddtag:"strptr"`
			Thing  fullTagStruct `ddtag:"someval"`
		}{}), "invalid field type: Thing")

		require.ErrorContains(t, ddtags.Validate(&struct {
			StrPtr *string        `ddtag:"strptr"`
			Thing  *fullTagStruct `ddtag:"someval"`
		}{}), "invalid field type: Thing")
	})
}
