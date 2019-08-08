package wrr

import (
	"fmt"
	"testing"
	"time"
)

//var i int = -1  //表示上一次选择的服务器
//var cw int = 0  //表示当前调度的权值
//var gcd int = 2 //当前所有权重的最大公约数 比如 2，4，8 的最大公约数为：2

type MyServiceTest struct {
	w    int
	name string
}

func (this *MyServiceTest) GetWeight() int {
	return this.w
}

var services = []Service{
	&MyServiceTest{2, "2"},
	&MyServiceTest{4, "4"},
	&MyServiceTest{8, "8"},
	&MyServiceTest{9, "9"},
}

func Test_GetService(t *testing.T) {
	record := make(map[string]int)

	wrb := NewWeightedRoundRobin(services)

	startTime := time.Now().Unix()

	for i := 0; i < 1000000; i++ {
		s := wrb.getService().(*MyServiceTest)
		if record[s.name] != 0 {
			record[s.name]++
		} else {
			record[s.name] = 1
		}
	}

	endTime := time.Now().Unix()

	fmt.Println("consume time: ", endTime-startTime)
	fmt.Println("====================================")
	for k, v := range record {
		fmt.Println(k, " : ", v)
	}
}
