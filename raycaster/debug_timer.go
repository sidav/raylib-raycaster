package raycaster

import "time"

type Timer struct {
	beginning       time.Time
	currentlyPassed time.Duration

	totalMeasures         int
	meanPassedCalculated  time.Duration
	meanPassedAccumulated time.Duration
}

func (b *Timer) NewMeasure() {
	const calculateMeanIn = 100
	if b.totalMeasures == calculateMeanIn {
		b.meanPassedCalculated = b.meanPassedAccumulated / time.Duration(calculateMeanIn)
		b.meanPassedAccumulated = 0
		b.totalMeasures = 0
	}

	b.totalMeasures++
	b.currentlyPassed = 0
	b.beginning = time.Now()
}

func (b *Timer) Measure(f func()) {
	start := time.Now()
	f()
	b.currentlyPassed += time.Since(start)
	b.meanPassedAccumulated += time.Since(start)
}

func (b *Timer) GetMeasuredPassedTime() time.Duration {
	return b.currentlyPassed
}

func (b *Timer) GetMeanPassedTime() time.Duration {
	return b.meanPassedCalculated
}
