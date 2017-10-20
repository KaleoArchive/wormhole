package job

import (
	"runtime"

	"github.com/jeffail/tunny"
	"github.com/kaleocheng/wormhole/trans"
)

// Job struct
type Job struct {
	Number    int
	Pool      *tunny.WorkPool
	RateLimit *float64
}

var j *Job

func init() {
	numCPUs := jobNum()
	j = &Job{
		Number: numCPUs,
	}
	runtime.GOMAXPROCS(numCPUs + 1)
}

func jobNum() int {
	return runtime.NumCPU()
}

func (j *Job) rateLimit() *float64 {
	if j.RateLimit == nil {
		return nil
	}

	if j.Number == 0 {
		return nil
	}

	result := *j.RateLimit / float64(uint(j.Number))
	return &result
}

// SetRateLimit limit the rate of trans
func SetRateLimit(ratelimie float64) {
	j.setRateLimit(ratelimie)
}

func (j *Job) setRateLimit(ratelimit float64) {
	j.RateLimit = &ratelimit
}

// Start create a new job pool and open it
func Start(t *trans.Trans) {
	j.Start(t)
}

// Start create a new job pool and open it
func (j *Job) Start(t *trans.Trans) {
	j.Pool, _ = tunny.CreatePool(j.Number, func(object interface{}) interface{} {
		image, _ := object.(*trans.Image)
		return t.Migrate(image, j.rateLimit())
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
