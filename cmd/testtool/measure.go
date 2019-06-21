package testtool

import (
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/process"
)

type Measure struct {
	mem     float64
	cputime float64
	time    float64
	pro     *process.Process
}

func NewMeasure() Measure {
	me := new(Measure)
	PID := os.Getpid()
	me.pro, _ = process.NewProcess(int32(PID))

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
	m.mem = float64(eMem.Alloc-sMem.Alloc) / float64(1024*1024)
}

func (m *Measure) StartCpu() {
	sCPUStat, _ := m.pro.Times()
	sCPU = sCPUStat.Total()
}

func (m *Measure) EndCpu() {
	eCPUStat, _ := m.pro.Times()
	eCPU = eCPUStat.Total()
}

func (m *Measure) CalcCpu() {
	m.cputime = eCPU - sCPU
}

func (m *Measure) StartTime() {
	sTime = time.Now()
}

func (m *Measure) EndTime() {
	eTime = time.Since(sTime)
}

func (m *Measure) CalcTime() {
	m.time = eTime.Seconds()
}

func (m *Measure) AllStart() {
	m.StartMem()
	m.StartCpu()
	m.StartTime()
}

func (m *Measure) AllEnd() {
	m.EndMem()
	m.EndCpu()
	m.EndTime()
}

func (m *Measure) AllCalc() {

}
