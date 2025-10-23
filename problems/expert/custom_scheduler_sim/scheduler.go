package customscheduler

// TODO: implement scheduler with work stealing

type Task func()

type Scheduler struct{}

func New() *Scheduler { return &Scheduler{} }
func (s *Scheduler) Submit(t Task) { /* TODO */ }
func (s *Scheduler) Run(workers int) { /* TODO */ }
