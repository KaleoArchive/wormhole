package job

import (
	"runtime"

	"github.com/jeffail/tunny"
	"github.com/kaleocheng/wormhole/trans"
)

// Job struct
type Job struct {
	Number int
	Pool   *tunny.WorkPool
}

var j *Job

func init() {
	numCPUs := InitJobNum()
	j = &Job{
		Number: numCPUs,
	}
	runtime.GOMAXPROCS(numCPUs + 1)
}

// InitJobNum return the job num base on CPU num for now
func InitJobNum() int {
	return runtime.NumCPU()
}

// Start create a new job pool and open it
func Start(t *trans.Trans) {
	j.Start(t)
}

// Start create a new job pool and open it
func (j *Job) Start(t *trans.Trans) {
	j.Pool, _ = tunny.CreatePool(j.Number, func(object interface{}) interface{} {
		image, _ := object.(*trans.Image)
		return t.Migrate(image)
	}).Open()
}

// Add a new job
func Add(i *trans.Image) (interface{}, error) {
	return j.Add(i)
}

// Add a new job
func (j *Job) Add(i *trans.Image) (interface{}, error) {
	return j.Pool.SendWork(i)
}

// Close the job pool
func Close() {
	j.Close()
}

// Close the job pool
func (j *Job) Close() {
	j.Pool.Close()
}
