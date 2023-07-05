package organization_test

import (
	"errors"
	"strings"
	"testing"

	"github.com/ispec-inc/starry/api-go-ddd-graphql/service/clinic/internal/domain/organization"
	"github.com/stretchr/testify/assert"
)

func TestName_Validate(t *testing.T) {
	type (
		give struct {
			name string
		}
		want struct {
			err error
		}
	)

	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "no_error",
			give: give{
				name: "鈴木歯科医院",
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "error_name_is_empty",
			give: give{
				name: "",
			},
			want: want{
				err: errors.New("organization: name is empty"),
			},
		},
		{
			name: "name_is_just_50",
			give: give{
				name: strings.Repeat("あ", 50),
			},
			want: want{
				err: nil,
			},
		},
		{
			name: "error_name_is_over_50",
			give: give{
				name: strings.Repeat("あ", 51),
			},
			want: want{
				err: errors.New("organization: name exceeds 50 characters"),
			},
		},
	}

	for i := range tests {
		test := tests[i]

		t.Run(test.name, func(t *testing.T) {
			name := organization.Name(test.give.name)
			err := name.Validate()
			if test.want.err != nil {
				assert.EqualError(t, test.want.err, err.Error())
			} else {
				assert.NoError(t, err)
			}
		})
	}
}
