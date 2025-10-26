package customscheduler

import "sync"

type Task func()

type Scheduler struct { tasks chan Task }

func New() *Scheduler { return &Scheduler{ tasks: make(chan Task, 1024) } }
func (s *Scheduler) Submit(t Task) { s.tasks <- t }
func (s *Scheduler) Run(workers int) { wg := sync.WaitGroup{}; for i:=0;i<workers;i++ { wg.Add(1); go func(){ for t := range s.tasks { t() }; wg.Done() }() }; wg.Wait() }
func (s *Scheduler) Close() { close(s.tasks) }
