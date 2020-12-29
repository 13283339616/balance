package balance

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var insts []*Instance
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		host := fmt.Sprintf("192.168.%d.%d", rand.Intn(255), i)
		wc := rand.Intn(10)
		one := NewInstance(host, 8080, int64(wc))
		insts = append(insts, one)
	}

	var balanceNames = []string{"hash", "random", "roundrobin", "weight_roundrobin", "shuffle", "shuffle2"}

	for _, name := range balanceNames {
		startTime := time.Now().UnixNano()
		for i := 0; i < 10000; i++ {
			_, err := DoBalance(name, insts)
			if err != nil {
				fmt.Println("Do balance err:", err)
				continue
			}
		}
		endTime := time.Now().UnixNano()
		fmt.Println("name: ", name, "cost time: ", (endTime-startTime)/1000)
		for _, inst := range insts {
			if name == "weight_roundrobin" {
				fmt.Println(inst.GetResult(), ";weight: ", inst.Weight)
			} else {
				fmt.Println(inst.GetResult())
			}
			inst.CallTimes = 0
		}
	}

}
