package addrlists

import "testing"

func Test_parse(t *testing.T) {
	type args struct {
		cidr string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Case 1",
			args: args{
				cidr: "1.1.1.1/32",
			},
			want:    "1.1.1.1",
			wantErr: false,
		},
		{
			name: "Case 2",
			args: args{
				cidr: "1.1.1.0/31",
			},
			want:    "1.1.1.0/31",
			wantErr: false,
		},
		{
			name: "Case 3",
			args: args{
				cidr: "1.1.1.1",
			},
			want:    "1.1.1.1",
			wantErr: false,
		},
		{
			name: "Case 4",
			args: args{
				cidr: "1.1.1.256",
			},
			want:    "",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parse(tt.args.cidr)
			if (err != nil) != tt.wantErr {
				t.Errorf("parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parse() got = %v, want %v", got, tt.want)
			}
		})
	}
}
