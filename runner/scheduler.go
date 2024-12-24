package main

import "fmt"

type Scheduler struct {
	releaseChanByIdx map[int64]chan struct{}
	lockChan         chan int64
}

func NewScheduler(minCPU, maxCPU int64) *Scheduler {
	lockChan := make(chan int64)
	releaseChanByIdx := make(map[int64]chan struct{})

	for cpuIdx := minCPU; cpuIdx <= maxCPU; cpuIdx++ {
		releaseChan := make(chan struct{})
		releaseChanByIdx[cpuIdx] = releaseChan

		go func(cpuIdx int64, releaseChan chan struct{}) {
			fmt.Printf("[scheduler] start cpu %d\n", cpuIdx)
			for {
				lockChan <- cpuIdx
				<-releaseChan
			}
		}(cpuIdx, releaseChan)
	}

	return &Scheduler{
		releaseChanByIdx: releaseChanByIdx,
		lockChan:         lockChan,
	}
}

func (s *Scheduler) WaitForFreeCPU() int64 {
	return <-s.lockChan
}

func (s *Scheduler) ReleaseCPU(cpu int64) {
	s.releaseChanByIdx[cpu] <- struct{}{}
}
