package testtool

import (
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/process"
)

type Measure struct {
	Mem     float64
	Cputime float64
	Time    float64
	Pro     *process.Process
}

func NewMeasure() Measure {
	me := new(Measure)
	PID := os.Getpid()
	me.Pro, _ = process.NewProcess(int32(PID))

	return *me
}

var (
	sMem  runtime.MemStats
	eMem  runtime.MemStats
	sCPU  float64
	eCPU  float64
	sTime time.Time
	eTime time.Duration
)

func (m *Measure) StartMem() {
	runtime.ReadMemStats(&sMem)
}

func (m *Measure) EndMem() {
	runtime.ReadMemStats(&eMem)
}

func (m *Measure) CalcMem() {
	m.Mem = float64(eMem.Alloc-sMem.Alloc) / float64(1024*1024)
}

func (m *Measure) StartCpu() {
	sCPUStat, _ := m.Pro.Times()
	sCPU = sCPUStat.Total()
}

func (m *Measure) EndCpu() {
	eCPUStat, _ := m.Pro.Times()
	eCPU = eCPUStat.Total()
}

func (m *Measure) CalcCpu() {
	m.Cputime = eCPU - sCPU
}

func (m *Measure) StartTime() {
	sTime = time.Now()
}

func (m *Measure) EndTime() {
	eTime = time.Since(sTime)
}

func (m *Measure) CalcTime() {
	m.Time = eTime.Seconds()
}

func (m *Measure) StartAll() {
	m.StartMem()
	m.StartCpu()
	m.StartTime()
}

func (m *Measure) EndAll() {
	m.EndMem()
	m.EndCpu()
	m.EndTime()
}

func (m *Measure) CalcAll() {
	m.CalcMem()
	m.CalcCpu()
	m.CalcTime()
}
