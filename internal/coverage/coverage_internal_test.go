package coverage

import (
	"testing"

	"golang.org/x/tools/cover"

	"github.com/go-gremlins/gremlins/internal/gomodule"
)

func TestRemoveModuleFromPathNormalizesSeparators(t *testing.T) {
	t.Parallel()

	cov := &Coverage{
		mod: gomodule.GoModule{
			Name:       "example.com/project",
			CallingDir: ".",
		},
	}

	got := cov.removeModuleFromPath(&cover.Profile{
		FileName: "example.com/project/internal/coverage/coverage.go",
	})

	want := "internal/coverage/coverage.go"
	if got != want {
		t.Fatalf("expected normalized path %q, got %q", want, got)
	}
}
