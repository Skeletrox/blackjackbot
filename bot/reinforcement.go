package bot

import "math"

func (r *Reinforcement) Init(Alpha, Gamma, RandomProb, TempDelta float64) {
	r.Alpha = Alpha
	r.Gamma = Gamma
	r.RandomProb = RandomProb
	r.TempDelta = TempDelta
	r.Rewards = make([][]float64, 21)
	for i := range r.Rewards {
		r.Rewards[i] = make([]float64, 2)
	}
	// If your score is 21, your best option is to stand.
	r.Rewards[20][1] = float64(math.MaxUint32)
}