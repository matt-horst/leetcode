package p3578

const MOD int = 1e9 + 7

type Deque struct {
	data []int
}

func NewDeque() Deque {
	return Deque{data: make([]int, 0)}
}

func (d *Deque) push(v int) {
	d.data = append(d.data, v)
}

func (d *Deque) pop() int {
	v := d.data[0]
	d.data = d.data[1:]

	return v
}

func (d *Deque) popEnd() int {
	v := d.data[len(d.data)-1]
	d.data = d.data[0 : len(d.data)-1]

	return v
}

func (d *Deque) isEmpty() bool {
	return len(d.data) == 0
}

func (d *Deque) peek() int {
	return d.data[0]
}

func (d *Deque) peekEnd() int {
	return d.data[len(d.data)-1]
}

func countPartitions(nums []int, k int) int {
	n := len(nums)
	dp := make([]int, n+1)
	prefix := make([]int, n+1)

	dp[0] = 1
	prefix[0] = 1

	maxQueue := NewDeque()
	minQueue := NewDeque()

	for i, j := 0, 0; i < n; i++ {
		v := nums[i]

		for !maxQueue.isEmpty() && nums[maxQueue.peekEnd()] <= v {
			maxQueue.popEnd()
		}
		maxQueue.push(i)

		for !minQueue.isEmpty() && nums[minQueue.peekEnd()] >= v {
			minQueue.popEnd()
		}
		minQueue.push(i)

		for !maxQueue.isEmpty() && !minQueue.isEmpty() &&
			nums[maxQueue.peek()]-nums[minQueue.peek()] > k {
			if maxQueue.peek() == j {
				maxQueue.pop()
			}

			if minQueue.peek() == j {
				minQueue.pop()
			}

			j++
		}

		if j > 0 {
			dp[i+1] = (prefix[i] - prefix[j-1] + MOD) % MOD
		} else {
			dp[i+1] = prefix[i] % MOD
		}

		prefix[i+1] = (prefix[i] + dp[i+1]) % MOD
	}

	return dp[n]
}
