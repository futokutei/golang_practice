package atomic_variant

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Start() {
	// Канали для парних та непарних чисел
	evenChan := make(chan int)
	oddChan := make(chan int)

	// Глобальний лічильник (обов'язково фіксованого розміру для atomic)
	var counter int64
	var wg sync.WaitGroup

	// ==========================================
	// ГОРУТИНИ ДЛЯ ОБРОБКИ ДАНИХ
	// ==========================================

	// Горутина 1: Обробка парних чисел
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-evenChan:
				if !ok {
					return
				}
				if val%3 == 0 {
					// Атомарно збільшуємо counter на 1 (counter++)
					atomic.AddInt64(&counter, 1)
				}
			}
		}
	}()

	// Горутина 2: Обробка непарних чисел
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			select {
			case val, ok := <-oddChan:
				if !ok {
					return
				}
				if val%33 == 0 {
					// Атомарно зменшуємо counter на 1 (counter--)
					// В Go віднімання в atomic робиться через додавання від'ємного числа
					atomic.AddInt64(&counter, -1)
				}
			}
		}
	}()

	// ==========================================
	// ЄДИНИЙ ГЕНЕРАТОР ЧИСЕЛ
	// ==========================================
	for i := 1; i <= 1000; i++ {
		if i%2 == 0 {
			evenChan <- i
		} else {
			oddChan <- i
		}
	}

	// Закриваємо канали після відправки всіх чисел
	close(evenChan)
	close(oddChan)

	// Чекаємо завершення обробки в горутинах
	wg.Wait()

	// Атомарно зчитуємо фінальне значення для безпечного виведення
	finalCounter := atomic.LoadInt64(&counter)
	fmt.Printf("Final value of counter in atomic variant: %d\n", finalCounter)
}
