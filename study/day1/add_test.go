package day1

import "testing"

func TestAdd(t *testing.T) {

	var dataset = []struct {
		a int
		b int
		c int
	}{
		{1, 2, 3},
		{2, 3, 5},
		{3, 4, 7},
	}
	for _, d := range dataset {
		if add(d.a, d.b) != d.c {
			t.Errorf("%d + %d != %d", d.a, d.b, d.c)
		}
	}
	if add(1, 2) != 3 {
		t.Error("1 + 2 != 3")
	}
}
func TestAdd2(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode.")
	}
}
func BenchmarkAdd(b *testing.B) {
	// 数据集 进行多次测试
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}

}
