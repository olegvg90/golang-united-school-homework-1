package solution

import (
	"testing"
	emoji "github.com/kyokomi/emoji"
)

func TestGetMessage(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "simple test",
			want: emoji.Sprint("Hello :world_map:!"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMessage(); got != tt.want {
				t.Errorf("GetMessage() = %v, want %v", got, tt.want)
			}
		})
	}
}
