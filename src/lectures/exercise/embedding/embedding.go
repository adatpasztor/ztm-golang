//--Summary:
//  Create a system monitoring dashboard using the existing dashboard
//  component structures. Each array element in the components represent
//  a 1-second sampling.
//
//--Requirements:
//* Create functions to calculate averages for each dashboard component
//* Using struct embedding, create a Dashboard structure that contains
//  each dashboard component
//* Print out a 5-second average from each component using promoted
//  methods on the Dashboard

package main

import "fmt"

type Bytes int
type Celcius float32

type BandwidthUsage struct {
	amount []Bytes
}

type MemoryUsage struct {
	amount []Bytes
}

type CpuTemp struct {
	temp []Celcius
}

func (b BandwidthUsage) AvgBandwithUsage() int {
	var sum int
	for i := 0; i < len(b.amount); i++ {
		sum += int(b.amount[i])
	}
	return sum / len(b.amount)
}

func (m MemoryUsage) AvgMemUsage() int {
	var sum int
	for i := 0; i < len(m.amount); i++ {
		sum += int(m.amount[i])
	}
	return sum / len(m.amount)
}

func (c CpuTemp) AvgCpuTemp() int {
	var sum int
	for i := 0; i < len(c.temp); i++ {
		sum += int(c.temp[i])
	}
	return sum / len(c.temp)
}

/// type embedding: a tipusokban levo mezok es a tipusok receiver fuggvenyei
/// mind hasznalhatoak a Dashboard tipusban
type Dashboard struct {
	BandwidthUsage
	MemoryUsage
	CpuTemp
}

func main() {
	bandwidth := BandwidthUsage{[]Bytes{50000, 100000, 130000, 80000, 90000}}
	// b := bandwidth.AvgBandwithUsage()
	// fmt.Println("b->", b)

	temp := CpuTemp{[]Celcius{50, 51, 53, 51, 52}}
	// c := temp.AvgCpuTemp()
	// fmt.Println("c->", c)

	memory := MemoryUsage{[]Bytes{800000, 800000, 810000, 820000, 800000}}
	// m := memory.AvgMemUsage()
	// fmt.Println("m->", m)
	d := Dashboard{
		BandwidthUsage: bandwidth,
		MemoryUsage:    memory,
		CpuTemp:        temp,
	}
	fmt.Println("Bandwith->", d.AvgBandwithUsage())
	fmt.Println("CpuTemp->", d.AvgCpuTemp())
	fmt.Println("MemUsage->", d.AvgMemUsage())

}
