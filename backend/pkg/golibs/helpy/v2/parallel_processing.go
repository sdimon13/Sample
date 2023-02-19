package helpy

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

type processingFunc func(m *sync.Mutex) error

// Обертка для параллельного запуска загрузок
type Processing struct {
	// Список функций для запуска
	routines map[string]processingFunc
	// Обрабатывать задачи поочереди
	singleThread bool
}

func NewProcessing() *Processing {
	return &Processing{
		routines: make(map[string]processingFunc),
	}
}

// Добавить новую функцию для запуска
func (pr *Processing) Push(name string, fn processingFunc) {
	pr.routines[name] = fn
}

// Выставить флаг, чтобы задачи выполнялись последовательно
func (pr *Processing) SingleThread() {
	pr.singleThread = true
}

// Запустить все добавленные функции и проверить их на наличие ошибок
func (pr Processing) Run() []error {
	m := sync.Mutex{}
	var errs []error

	maxRoutines := runtime.NumCPU()
	var activeRoutines int64

	var wg sync.WaitGroup
	for name, fn := range pr.routines {
		if pr.singleThread {
			if err := fn(&m); err != nil {
				errs = append(errs, fmt.Errorf("%s %w", name, err))
			}
			continue
		}

		// Do not run new goroutine if all CPU is busy
		for {
			if atomic.LoadInt64(&activeRoutines) < int64(maxRoutines*2) {
				break
			}
			time.Sleep(10)
		}

		wg.Add(1)
		atomic.AddInt64(&activeRoutines, 1)
		go func(name string, fn processingFunc) {
			defer func() {
				wg.Done()
				atomic.AddInt64(&activeRoutines, -1)
			}()
			if err := fn(&m); err != nil {
				m.Lock()
				errs = append(errs, fmt.Errorf("%s %w", name, err))
				m.Unlock()
			}
		}(name, fn)
	}
	wg.Wait()

	return errs
}
