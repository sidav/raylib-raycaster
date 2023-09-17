package raycaster

import "time"

type timer struct {
	beginning       time.Time
	currentlyPassed time.Duration

	totalMeasures         int
	meanPassedCalculated  time.Duration
	meanPassedAccumulated time.Duration
}

func (b *timer) newMeasure() {
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

func (b *timer) measure(f func()) {
	start := time.Now()
	f()
	b.currentlyPassed += time.Since(start)
	b.meanPassedAccumulated += time.Since(start)
}

func (b *timer) getMeasuredPassedTime() time.Duration {
	return b.currentlyPassed
}

func (b *timer) getMeanPassedTime() time.Duration {
	return b.meanPassedCalculated
}
