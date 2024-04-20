package main

import (
	"log"
	"time"
)

func main() {

	/*
		Ini adalah titik masuk untuk proses latar belakang atau worker.
		Dalam contoh ini, worker hanya mencetak pesan ke log setiap 10 detik,
		tapi Anda bisa memperluasnya untuk menjalankan tugas-tugas seperti pengolahan antrian,
		scheduled tasks, atau operasi background lainnya.
	*/
	for {
		log.Println("Worker is processing tasks...")
		// Tempatkan logika worker Anda di sini
		time.Sleep(10 * time.Second) // Contoh: tunggu 10 detik
	}
}
