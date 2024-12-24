package main

import (
	"fmt"
	"os"
)

func Main() error {
	args, err := parseArgs()
	if err != nil {
		return err
	}

	scheduler := NewScheduler(args.minCPU, args.maxCPU)

	var idx int64 = 0
	for {
		cpu := scheduler.WaitForFreeCPU()
		go func(idx int64, cpu int64) {
			defer scheduler.ReleaseCPU(cpu)

			topology := idx % args.nTopologies
			seed := args.startSeed + idx/args.nTopologies

			fmt.Printf("running seed=%d, topology=%d on cpu=%d\n", seed, topology, cpu)

			err := runAndSaveStdout(args.logsPath, args.execPath, cpu, seed, topology)
			if err != nil {
				fmt.Println(err)
			}
		}(idx, cpu)

		idx += 1
	}
}

func main() {
	err := Main()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
