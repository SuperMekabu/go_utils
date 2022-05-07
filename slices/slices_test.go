package slices

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestChunk(t *testing.T) {
	type args[T any] struct {
		org       []T
		chunkSize int
	}

	type test[M any] struct {
		name string
		args args[M]
		want [][]M
	}

	tests := []test[int]{
		{
			name: "want OK",
			args: args[int]{
				org:       []int{1, 2, 3, 4, 5},
				chunkSize: 2,
			},
			want: [][]int{{1, 2}, {3, 4}, {5}},
		},
		{
			name: "want OK2",
			args: args[int]{
				org:       []int{1, 2, 3, 4, 5},
				chunkSize: 1,
			},
			want: [][]int{{1}, {2}, {3}, {4}, {5}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.org, tt.args.chunkSize); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args[T any] struct {
		src []T
		fn  func(T) bool
	}

	type test[M any] struct {
		name string
		args args[M]
		want []M
	}

	stringTests := []test[string]{
		{
			name: "string OK",
			args: args[string]{
				[]string{"a", "b", "aa"},
				func(t string) bool {
					return strings.Contains(t, "a")
				},
			},
			want: []string{"a", "aa"},
		}, {
			name: "string OK2",
			args: args[string]{
				[]string{"a", "b", "aa"},
				func(t string) bool {
					return strings.HasPrefix(t, "b")
				},
			},
			want: []string{"b"},
		},
	}

	intTests := []test[int]{
		{
			name: "int OK",
			args: args[int]{
				[]int{1, 2, 3, 4, 5},
				func(t int) bool {
					return t%2 == 0
				},
			},
			want: []int{2, 4},
		},
		{
			name: "int OK2",
			args: args[int]{
				[]int{1, 2, 3, 4, 5},
				func(t int) bool {
					return math.Pow(float64(t), 2) > 10
				},
			},
			want: []int{4, 5},
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.src, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.src, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[T any, M any] struct {
		src []T
		fn  func(T) (M, bool)
	}

	type test[K any, L any] struct {
		name string
		args args[K, L]
		want []L
	}

	stringTests := []test[string, string]{
		{
			name: "string OK",
			args: args[string, string]{
				[]string{"1A", "2A", "c", "4A"},
				func(t string) (string, bool) {
					if strings.HasSuffix(t, "A") {
						return t, true
					}
					return "", false
				},
			},
			want: []string{"1A", "2A", "4A"},
		},
	}

	stringTests2 := []test[string, string]{
		{
			name: "string OK2",
			args: args[string, string]{
				[]string{"1", "2", "3"},
				func(t string) (string, bool) {
					return fmt.Sprintf("%s:%s", t, t), true
				},
			},
			want: []string{"1:1", "2:2", "3:3"},
		},
	}

	intTests := []test[int, float64]{
		{
			name: "int OK",
			args: args[int, float64]{
				[]int{1, 2, 3, 4, 5},
				func(t int) (float64, bool) {
					return float64(t) / 2, true
				},
			},
			want: []float64{0.5, 1.0, 1.5, 2.0, 2.5},
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.src, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range stringTests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.src, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.src, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIncludes(t *testing.T) {
	type args[T comparable] struct {
		src []T
		tgt T
	}

	type test[K comparable] struct {
		name string
		args args[K]
		want bool
	}

	stringTests := []test[string]{
		{
			name: "string OK",
			args: args[string]{
				[]string{"1", "2", "3"},
				"2",
			},
			want: true,
		}, {
			name: "string OK2",
			args: args[string]{
				[]string{"1", "2", "3"},
				"4",
			},
			want: false,
		},
	}

	intTests := []test[int]{
		{
			name: "int OK",
			args: args[int]{
				[]int{1, 2, 3},
				2,
			},
			want: true,
		},
		{
			name: "int OK2",
			args: args[int]{
				[]int{1, 2, 3},
				4,
			},
			want: false,
		},
	}

	type original struct {
		id   int
		name string
	}

	orgTests := []test[original]{
		{
			name: "original OK",
			args: args[original]{
				[]original{{1, "john"}, {2, "jack"}, {3, "jade"}},
				original{2, "jack"},
			},
			want: true,
		},
		{
			name: "original OK",
			args: args[original]{
				[]original{{1, "john"}, {2, "jack"}, {3, "jade"}},
				original{4, "jason"},
			},
			want: false,
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Includes(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Includes() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Includes(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Includes() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range orgTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Includes(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Includes() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestRemoveFirst(t *testing.T) {
	type args[T comparable] struct {
		src []T
		tgt T
	}

	type test[K comparable] struct {
		name string
		args args[K]
		want []K
	}

	stringTests := []test[string]{
		{
			name: "string OK",
			args: args[string]{
				[]string{"1", "2", "3"},
				"2",
			},
			want: []string{"1", "3"},
		},
		{
			name: "string OK2",
			args: args[string]{
				[]string{"2", "2", "3"},
				"2",
			},
			want: []string{"2", "3"},
		},
	}

	intTests := []test[int]{
		{
			name: "int OK",
			args: args[int]{
				[]int{1, 2, 3},
				2,
			},
			want: []int{1, 3},
		},
	}

	type original struct {
		id   int
		name string
	}

	orgTests := []test[original]{
		{
			name: "original OK",
			args: args[original]{
				[]original{{1, "john"}, {2, "jack"}, {3, "jade"}},
				original{2, "jack"},
			},
			want: []original{{1, "john"}, {3, "jade"}},
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFirst(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFirst() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFirst(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFirst() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range orgTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveFirst(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveFirst() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestRemoveAll(t *testing.T) {
	type args[T comparable] struct {
		src []T
		tgt T
	}

	type test[K comparable] struct {
		name string
		args args[K]
		want []K
	}

	stringTests := []test[string]{
		{
			name: "string OK",
			args: args[string]{
				[]string{"1", "2", "3"},
				"2",
			},
			want: []string{"1", "3"},
		},
		{
			name: "string OK2",
			args: args[string]{
				[]string{"2", "2", "3"},
				"2",
			},
			want: []string{"3"},
		},
	}

	intTests := []test[int]{
		{
			name: "int OK",
			args: args[int]{
				[]int{1, 2, 3},
				2,
			},
			want: []int{1, 3},
		},
		{
			name: "int OK2",
			args: args[int]{
				[]int{2, 2, 3},
				2,
			},
			want: []int{3},
		},
	}

	type original struct {
		id   int
		name string
	}

	orgTests := []test[original]{
		{
			name: "original OK",
			args: args[original]{
				[]original{{1, "john"}, {2, "jack"}, {3, "jade"}},
				original{2, "jack"},
			},
			want: []original{{1, "john"}, {3, "jade"}},
		},
		{
			name: "original OK2",
			args: args[original]{
				[]original{{2, "jack"}, {2, "jack"}, {3, "jade"}},
				original{2, "jack"},
			},
			want: []original{{3, "jade"}},
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAll(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAll() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAll(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAll() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range orgTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveAll(tt.args.src, tt.args.tgt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveAll() = %v, want %v", got, tt.want)
			}
		})
	}

}

func TestEvery(t *testing.T) {
	type args[T any] struct {
		src []T
		fn  func(T) bool
	}

	type test[K any] struct {
		name string
		args args[K]
		want bool
	}

	stringTests := []test[string]{
		{
			name: "string OK",
			args: args[string]{
				[]string{"1", "2", "3"},
				func(t string) bool {
					parse, err := strconv.Atoi(t)
					if err != nil {
						return false
					}
					return parse < 4
				},
			},
			want: true,
		}, {
			name: "string OK2",
			args: args[string]{
				[]string{"1", "2", "3"},
				func(t string) bool {
					return t == "2"
				},
			},
			want: false,
		},
	}

	intTests := []test[int]{
		{
			name: "int OK",
			args: args[int]{
				[]int{2, 4, 6},
				func(t int) bool {
					return t%2 == 0
				},
			},
			want: true,
		}, {
			name: "int OK2",
			args: args[int]{
				[]int{1, 2, 3, 4, 5},
				func(t int) bool {
					return t%2 == 0
				},
			},
			want: false,
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Every(tt.args.src, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Every() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Every(tt.args.src, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Every() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSome(t *testing.T) {
	type args[T any] struct {
		src []T
		fn  func(T) bool
	}

	type test[K any] struct {
		name string
		args args[K]
		want bool
	}

	stringTests := []test[string]{
		{
			name: "string OK",
			args: args[string]{
				[]string{"1", "2", "3"},
				func(t string) bool {
					parse, err := strconv.Atoi(t)
					if err != nil {
						return false
					}
					return parse > 3
				},
			},
			want: false,
		}, {
			name: "string OK2",
			args: args[string]{
				[]string{"1", "2", "3"},
				func(t string) bool {
					return t == "2"
				},
			},
			want: true,
		},
	}

	intTests := []test[int]{
		{
			name: "int OK",
			args: args[int]{
				[]int{1, 3, 5},
				func(t int) bool {
					return t%2 == 0
				},
			},
			want: false,
		}, {
			name: "int OK2",
			args: args[int]{
				[]int{1, 2, 3, 4, 5},
				func(t int) bool {
					return t%2 == 0
				},
			},
			want: true,
		},
	}

	for _, tt := range stringTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Some(tt.args.src, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Some() = %v, want %v", got, tt.want)
			}
		})
	}

	for _, tt := range intTests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Some(tt.args.src, tt.args.fn); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Some() = %v, want %v", got, tt.want)
			}
		})
	}
}
