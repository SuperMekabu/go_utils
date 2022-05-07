package maps

import (
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestFilter(t *testing.T) {
	type args[K comparable, V any] struct {
		src map[K]V
		fn  func(K, V) bool
	}

	type test[A comparable, B any] struct {
		name string
		args args[A, B]
		want map[A]B
	}

	stringTests := []test[int, int]{
		{
			name: "wants filtered",
			args: args[int, int]{
				src: map[int]int{1: 1, 2: 3, 3: 3},
				fn: func(k, v int) bool {
					return k == v
				},
			},
			want: map[int]int{1: 1, 3: 3},
		},
		{
			name: "wants not filtered",
			args: args[int, int]{
				src: map[int]int{1: 1, 2: 2, 3: 3},
				fn: func(k, v int) bool {
					return k == v
				},
			},
			want: map[int]int{1: 1, 2: 2, 3: 3},
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.src, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[K comparable, V any, R any] struct {
		src map[K]V
		fn  func(K, V) (R, bool)
	}

	type test[A comparable, B any, C any] struct {
		name string
		args args[A, B, C]
		want []C
	}

	stringTests := []test[int, int, int]{
		{
			name: "wants filter mapped",
			args: args[int, int, int]{
				src: map[int]int{1: 1, 2: 3, 3: 3},
				fn: func(k, v int) (int, bool) {
					if k == v {
						return v, true
					}
					return -1, false
				},
			},
			want: []int{1, 3},
		},
		{
			name: "wants no filter mapped",
			args: args[int, int, int]{
				src: map[int]int{1: 1, 2: 2, 3: 3},
				fn: func(k, v int) (int, bool) {
					if k == v {
						return v, true
					}
					return -1, false
				},
			},
			want: []int{1, 2, 3},
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			got := Map(tt.args.src, tt.args.fn)
			sort.Ints(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasKey(t *testing.T) {
	type args[K comparable, V any] struct {
		src map[K]V
		tgt K
	}

	type test[A comparable, B any] struct {
		name string
		args args[A, B]
		want bool
	}

	stringTests := []test[int, int]{
		{
			name: "wants true",
			args: args[int, int]{
				src: map[int]int{1: 1, 2: 3, 3: 3},
				tgt: 3,
			},
			want: true,
		},
		{
			name: "wants false",
			args: args[int, int]{
				src: map[int]int{1: 1, 2: 2, 3: 3},
				tgt: 4,
			},
			want: false,
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasKey(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasKey() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasValue(t *testing.T) {
	type args[K comparable, V any] struct {
		src map[K]V
		tgt K
	}

	type test[A comparable, B any] struct {
		name string
		args args[A, B]
		want bool
	}

	stringTests := []test[int, int]{
		{
			name: "wants true",
			args: args[int, int]{
				src: map[int]int{1: 1, 2: 3, 3: 3},
				tgt: 3,
			},
			want: true,
		},
		{
			name: "wants false",
			args: args[int, int]{
				src: map[int]int{1: 1, 2: 2, 3: 3},
				tgt: 4,
			},
			want: false,
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasValue(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HasValue() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemove(t *testing.T) {
	type args[K comparable, V any] struct {
		src map[K]V
		tgt K
	}

	type test[A comparable, B any] struct {
		name string
		args args[A, B]
		want map[A]B
	}

	stringTests := []test[int, int]{
		{
			name: "wants removed",
			args: args[int, int]{
				src: map[int]int{1: 1, 2: 3, 3: 3},
				tgt: 3,
			},
			want: map[int]int{1: 1, 2: 3},
		},
		{
			name: "wants not removed",
			args: args[int, int]{
				src: map[int]int{1: 1, 2: 2, 3: 3},
				tgt: 4,
			},
			want: map[int]int{1: 1, 2: 2, 3: 3},
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Remove() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEvery(t *testing.T) {
	type args[K comparable, V any] struct {
		src map[K]V
		fn  func(K, V) bool
	}
	type test[A comparable, B any] struct {
		name string
		args args[A, B]
		want bool
	}

	tests := []test[string, int]{
		{
			name: "Wants OK",
			args: args[string, int]{
				src: map[string]int{"1": 1, "2": 2, "3": 3},
				fn: func(k string, v int) bool {
					p, err := strconv.ParseInt(k, 10, 64)
					if err != nil {
						t.Fatalf("failed parse int %v", err)
					}
					return int64(v) == p
				},
			},
			want: true,
		}, {
			name: "Wants NG",
			args: args[string, int]{
				src: map[string]int{"1": 1, "2": 1, "3": 3},
				fn: func(k string, v int) bool {
					p, err := strconv.ParseInt(k, 10, 64)
					if err != nil {
						t.Fatalf("failed parse int %v", err)
					}
					return int64(v) == p
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Every(tt.args.src, tt.args.fn); got != tt.want {
				t.Errorf("Every() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSome(t *testing.T) {
	type args[K comparable, V any] struct {
		src map[K]V
		fn  func(K, V) bool
	}
	type test[A comparable, B any] struct {
		name string
		args args[A, B]
		want bool
	}

	tests := []test[string, int]{
		{
			name: "Wants OK",
			args: args[string, int]{
				src: map[string]int{"1": 2, "2": 3, "3": 3},
				fn: func(k string, v int) bool {
					p, err := strconv.ParseInt(k, 10, 64)
					if err != nil {
						t.Fatalf("failed parse int %v", err)
					}
					return int64(v) == p
				},
			},
			want: true,
		}, {
			name: "Wants NG",
			args: args[string, int]{
				src: map[string]int{"1": 2, "2": 1, "3": 1},
				fn: func(k string, v int) bool {
					p, err := strconv.ParseInt(k, 10, 64)
					if err != nil {
						t.Fatalf("failed parse int %v", err)
					}
					return int64(v) == p
				},
			},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Some(tt.args.src, tt.args.fn); got != tt.want {
				t.Errorf("Some() = %v, want %v", got, tt.want)
			}
		})
	}
}
