package WeightedRoundRobin

import (
	"testing"
)

//var i int = -1  //表示上一次选择的服务器
//var cw int = 0  //表示当前调度的权值
//var gcd int = 2 //当前所有权重的最大公约数 比如 2，4，8 的最大公约数为：2

type MyServiceTest2 struct {
	w    int
	name string
}

func (this *MyServiceTest2) GetWeight() int {
	return this.w
}

var services2 = []Service{
	&MyServiceTest2{2, "2"},
	&MyServiceTest2{4, "4"},
	&MyServiceTest2{8, "8"},
	&MyServiceTest2{9, "9"},
}

func Benchmark_GetService(b *testing.B) {
	b.StopTimer()
	wrb := NewWeightedRoundRobin(services2)

	b.StartTimer()
	for i := 0; i < b.N; i++ { //use b.N for looping
		wrb.getService()
	}
}
