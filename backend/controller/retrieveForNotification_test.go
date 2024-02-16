package controller

import (
	"reflect"
	"testing"
)

func TestGetMentionedStudents(t *testing.T) {
	type args struct {
		notification string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test case 1",
			args: args{
				notification: "Hello students!",
			},
			want: []string{},
		},
		{
			name: "Test case 2",
			args: args{
				notification: "Hello students! tom@domain.com, james@example.com is absent today.",
			},
			want: []string{"tom@domain.com", "james@example.com"},
		},
		{
			name: "Test case 3",
			args: args{
				notification: "Hello students! joe@gmail.com, @invalid.com to meet principal Lim @3pm.",
			},
			want: []string{"joe@gmail.com"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMentionedStudents(tt.args.notification); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMentionedStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}

