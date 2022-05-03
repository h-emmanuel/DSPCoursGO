package cours

import (
	"fmt"
	"time"
)

func Master() {
	// creer un iPC (Internal Process Channel)
	c := make(chan string)

	// 4 actions longue a effectuer
	go worker(c, func() string { return tache(10) })
	go worker(c, func() string { return tache(5) })
	go worker(c, func() string { return tache(1) })
	go worker(c, func() string { return tache(3) })

	// recupère les resultat
	i := 0
	for result := range c {
		fmt.Println(result)

		if i >= 3 {
			break
		}
		i++
	}
}

func tache(delta time.Duration) string {
	time.Sleep(delta * time.Second)
	return fmt.Sprintf("%v", delta*time.Second)
}

func worker(c chan string, callback func() string) {
	c <- callback()
}
