package main

// Generator is responsible for running the generation of new timetables
type Generator struct {
	Counters        *ThreadSafeCounters
	NumberOfWorkers int
	ProcessingQueue *ProcessingQueue
	ShouldFinish    chan bool
}

// Start starts the generation of new timetables based on the NumberOfWorkers
func (g *Generator) Start() {
	for i := 0; i < g.NumberOfWorkers; i++ {
		go g.GenerationWorkerStart()
	}
}

// GenerationWorkerStart is the worker that generates new timetables
func (g *Generator) GenerationWorkerStart() {
	defaultTimeTable := Table{}
	defaultTimeTable.CreateDefault()
	for {
		select {
		case <-g.ShouldFinish:
			println(g.ProcessingQueue.Hashes[defaultTimeTable.Hash()])
			return
		default:
			defaultTimeTable.Shuffle()
			g.ProcessingQueue.AddHash(defaultTimeTable.Hash())
			g.Counters.IncrementGenerated()
		}
	}
}
