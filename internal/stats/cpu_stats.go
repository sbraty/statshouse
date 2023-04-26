package stats

import (
	"fmt"
	"time"

	"github.com/prometheus/procfs"
	"github.com/vkcom/statshouse/internal/format"
)

type CPUStats struct {
	fs procfs.FS

	stat   procfs.Stat
	stats  map[int64]procfs.CPUStat
	writer MetricWriter
}

const (
	cpu  = format.BuiltinMetricNameCpuUsage
	irq  = ""
	sirq = ""
	cs   = ""
)

func (*CPUStats) Name() string {
	return "cpu_stats"
}

func (c *CPUStats) PushDuration(now int64, d time.Duration) {
	c.writer.WriteSystemMetricValueWithoutHost(now, format.BuiltinMetricNameSystemMetricScrapeDuration, d.Seconds(), format.TagValueIDSystemMetricCPU)
}

func NewCpuStats(writer MetricWriter) (*CPUStats, error) {
	fs, err := procfs.NewFS(procfs.DefaultMountPoint)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize procfs: %w", err)
	}
	return &CPUStats{
		fs:     fs,
		stats:  map[int64]procfs.CPUStat{},
		writer: writer,
	}, nil
}

func (c *CPUStats) WriteMetrics(nowUnix int64) error {
	stat, err := c.fs.Stat()
	if err != nil {
		return fmt.Errorf("failed to get cpu stats: %w", err)
	}
	writeMetric := true
	if c.stat.BootTime == 0 {
		writeMetric = false
	}
	if writeMetric {
		err = c.writeCPU(nowUnix, stat)
		if err != nil {
			return fmt.Errorf("failed to write cpu metrics: %w", err)
		}
		err = c.writeSystem(nowUnix, stat)
		if err != nil {
			return fmt.Errorf("failed to write system metrics: %w", err)
		}
	}
	c.stat = stat
	err = c.updateCPUStats(stat.CPU)
	if err != nil {
		return fmt.Errorf("failed to update cpu stats: %w", err)
	}
	return nil
}

func (c *CPUStats) writeCPU(nowUnix int64, stat procfs.Stat) error {
	t := stat.CPUTotal
	oldT := c.stat.CPUTotal
	c.writer.WriteSystemMetricValue(nowUnix, cpu, t.User-oldT.User, format.RawIDTagUser)
	c.writer.WriteSystemMetricValue(nowUnix, cpu, t.Nice-oldT.Nice, format.RawIDTagNice)
	c.writer.WriteSystemMetricValue(nowUnix, cpu, t.System-oldT.System, format.RawIDTagSystem)
	c.writer.WriteSystemMetricValue(nowUnix, cpu, t.Idle-oldT.Idle, format.RawIDTagIdle)
	c.writer.WriteSystemMetricValue(nowUnix, cpu, t.Iowait-oldT.Iowait, format.RawIDTagIOWait)
	c.writer.WriteSystemMetricValue(nowUnix, cpu, t.IRQ-oldT.IRQ, format.RawIDTagIRQ)
	c.writer.WriteSystemMetricValue(nowUnix, cpu, t.SoftIRQ-oldT.SoftIRQ, format.RawIDTagSoftIRQ)
	c.writer.WriteSystemMetricValue(nowUnix, cpu, t.Steal-oldT.Steal, format.RawIDTagSteal)
	irqTotal := stat.IRQTotal - c.stat.IRQTotal
	sirqTotal := stat.SoftIRQTotal - c.stat.SoftIRQTotal
	c.writer.WriteSystemMetricValue(nowUnix, irq, float64(irqTotal))
	c.writer.WriteSystemMetricValue(nowUnix, sirq, float64(sirqTotal))
	return nil
}

func (c *CPUStats) writeSystem(nowUnix int64, stat procfs.Stat) error {
	uptime := uint64(time.Now().Unix()) - stat.BootTime
	c.writer.WriteSystemMetricValue(nowUnix, format.BuiltinMetricNameSystemUptime, float64(uptime))
	c.writer.WriteSystemMetricValue(nowUnix, format.BuiltinMetricNameProcessStatus, float64(stat.ProcessesRunning), format.RawIDTagRunning)
	c.writer.WriteSystemMetricValue(nowUnix, format.BuiltinMetricNameProcessStatus, float64(stat.ProcessesBlocked), format.RawIDTagBlocked)
	c.writer.WriteSystemMetricCount(nowUnix, format.BuiltinMetricNameProcessCreated, float64(stat.ProcessCreated-c.stat.ProcessCreated))
	c.writer.WriteSystemMetricValue(nowUnix, cs, float64(stat.ContextSwitches))
	return nil
}

func (c *CPUStats) updateCPUStats(new map[int64]procfs.CPUStat) error {
	for cpu, stat := range new {
		c.stats[cpu] = stat
	}
	return nil
}