package WeightedRoundRobin

import "sync"

//lastId int = -1  //上一次选择的服务器
//cw int = 0  //当前调度的权值
//gcd int = 2 //所有权重的最大公约数

type Service interface {
	GetWeight() int
}

type weightedRoundRobin struct {
	lastId   int
	cw       int
	gcd      int
	max      int
	Services []Service
	mu       sync.Mutex
	gcdMu    sync.Mutex
}

func NewWeightedRoundRobin(services []Service) *weightedRoundRobin {
	return &weightedRoundRobin{
		lastId:   -1,
		cw:       0,
		gcd:      -1,
		max:      -1,
		Services: services,
	}
}

// 辗转相除法取得公约数
func (this *weightedRoundRobin) getGcd(x int, y int) int {
	if y == 0 {
		return x
	}

	return this.getGcd(y, x%y)
}

// 取得所有权重的最大公约数
func (this *weightedRoundRobin) getMaxGcdForServices() int {
	if this.gcd == -1 {
		this.gcdMu.Lock()

		if this.gcd == -1 {
			w := 0
			for i := 0; i < len(this.Services)-1; i++ {
				if w == 0 {
					w = this.getGcd(this.Services[i].GetWeight(), this.Services[i+1].GetWeight())
				} else {
					w = this.getGcd(w, this.Services[i+1].GetWeight())
				}
			}

			this.gcd = w
		}

		this.gcdMu.Unlock()
	}

	return this.gcd
}

func (this *weightedRoundRobin) getService() Service {
	for {
		this.lastId = (this.lastId + 1) % len(this.Services)
		if this.lastId == 0 {
			this.cw = this.cw - this.getMaxGcdForServices()
			if this.cw <= 0 {
				this.cw = this.getMaxWeight()
				if this.cw == 0 {
					return nil
				}
			}
		}

		if weight := this.Services[this.lastId].GetWeight(); weight >= this.cw {
			return this.Services[this.lastId]
		}
	}
}

func (this *weightedRoundRobin) getMaxWeight() int {
	if this.max == -1 {
		this.mu.Lock()

		if this.max == -1 {
			max := 0
			for _, v := range this.Services {
				if weight := v.GetWeight(); weight >= this.max {
					max = weight
				}
			}
			this.max = max
		}

		this.mu.Unlock()
	}

	return this.max
}
