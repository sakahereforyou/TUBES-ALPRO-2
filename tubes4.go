package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Flight struct {
	Airline    string
	Price      int
	Departure  string
	Arrival    string
	FlightTime string
}

type Booking struct {
	Name        string
	Flight      Flight
	Destination string
}

var bookings []Booking
var flights []Flight

func main() {
	// Initialize with some dummy flights
	flights = []Flight{
		{"Garuda Indonesia", 1500000, "08:00", "10:00", "2h"},
		{"Lion Air", 900000, "09:00", "11:00", "2h"},
		{"AirAsia", 1200000, "10:00", "12:00", "2h"},
		{"Singapore Airlines", 2500000, "11:00", "13:00", "2h"},
	}

	for {
		choice := displayMenu()
		if choice == 4 {
			addFlight()
			continue
		} else if choice == 3 {
			fmt.Println("Terima kasih telah menggunakan layanan kami.")
			break
		} else if choice != 1 && choice != 2 {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			time.Sleep(2 * time.Second)
			continue
		}

		var destination string
		if choice == 1 {
			fmt.Print("Masukkan tujuan penerbangan domestik: ")
		} else if choice == 2 {
			fmt.Print("Masukkan tujuan penerbangan internasional: ")
		}
		fmt.Scan(&destination)

		destinationFlights := searchFlights(destination)
		if len(destinationFlights) == 0 {
			fmt.Println("Tidak ada penerbangan yang tersedia untuk tujuan ini.")
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Println("Daftar penerbangan:")
		printFlights(destinationFlights)

		sortChoice := displaySortMenu()
		switch sortChoice {
		case 1:
			insertionSort(destinationFlights, func(i, j int) bool {
				return destinationFlights[i].Price < destinationFlights[j].Price
			})
		case 2:
			insertionSort(destinationFlights, func(i, j int) bool {
				return destinationFlights[i].Price > destinationFlights[j].Price
			})
		case 3:
			insertionSort(destinationFlights, func(i, j int) bool {
				return destinationFlights[i].Departure < destinationFlights[j].Departure
			})
		default:
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			time.Sleep(2 * time.Second)
			continue
		}

		fmt.Println("Daftar penerbangan setelah sorting:")
		printFlights(destinationFlights)

		var flightChoice int
		fmt.Print("Pilih penerbangan (nomor): ")
		fmt.Scan(&flightChoice)
		if flightChoice < 1 || flightChoice > len(destinationFlights) {
			fmt.Println("Pilihan tidak valid. Silakan coba lagi.")
			time.Sleep(2 * time.Second)
			continue
		}

		flight := destinationFlights[flightChoice-1]
		fmt.Printf("Detail penerbangan yang dipilih:\nMaskapai: %s\nHarga: Rp%d\nKeberangkatan: %s\nKedatangan: %s\nDurasi: %s\n",
			flight.Airline, flight.Price, flight.Departure, flight.Arrival, flight.FlightTime)

		var confirm string
		fmt.Print("Apakah Anda ingin melanjutkan dengan penerbangan ini? (yes/no): ")
		fmt.Scan(&confirm)
		if strings.ToLower(confirm) != "yes" {
			fmt.Println("Pemesanan dibatalkan.")
			time.Sleep(2 * time.Second)
			continue
		}

		var name string
		fmt.Print("Masukkan nama Anda: ")
		fmt.Scan(&name)

		booking := Booking{
			Name:        name,
			Flight:      flight,
			Destination: destination,
		}
		bookings = append(bookings, booking)

		fmt.Printf("Anda telah sukses terdaftar pada penerbangan %s dengan tujuan %s.\n", flight.Airline, destination)
		time.Sleep(3 * time.Second)
	}
}

func displayMenu() int {
	var choice int
	fmt.Println("Menu:")
	fmt.Println("1. Penerbangan Domestik")
	fmt.Println("2. Penerbangan Internasional")
	fmt.Println("3. Keluar")
	fmt.Println("4. Tambahkan Penerbangan")
	fmt.Print("Pilih menu: ")
	fmt.Scan(&choice)
	return choice
}

func displaySortMenu() int {
	var choice int
	fmt.Println("Menu Sorting:")
	fmt.Println("1. Termurah ke Termahal")
	fmt.Println("2. Termahal ke Termurah")
	fmt.Println("3. Berdasarkan Waktu Keberangkatan")
	fmt.Print("Pilih menu sorting: ")
	fmt.Scan(&choice)
	return choice
}

func addFlight() {
	scanner := bufio.NewScanner(os.Stdin)

	var airline, departure, arrival, flightTime string
	var price int

	fmt.Print("Masukkan nama maskapai: ")
	scanner.Scan()
	airline = scanner.Text()

	fmt.Print("Masukkan harga: ")
	fmt.Scan(&price)

	fmt.Print("Masukkan waktu keberangkatan (HH:MM): ")
	scanner.Scan()
	departure = scanner.Text()

	fmt.Print("Masukkan waktu kedatangan (HH:MM): ")
	scanner.Scan()
	arrival = scanner.Text()

	fmt.Print("Masukkan durasi penerbangan (xh): ")
	scanner.Scan()
	flightTime = scanner.Text()

	newFlight := Flight{
		Airline:    airline,
		Price:      price,
		Departure:  departure,
		Arrival:    arrival,
		FlightTime: flightTime,
	}
	flights = append(flights, newFlight)

	fmt.Println("Penerbangan berhasil ditambahkan.")
	time.Sleep(2 * time.Second)
}

func getFlights(destination string) []Flight {
	return flights
}

func searchFlights(destination string) []Flight {
	allFlights := getFlights(destination)
	var destinationFlights []Flight
	for _, flight := range allFlights {
		if strings.Contains(strings.ToLower(flight.Airline), strings.ToLower(destination)) {
			destinationFlights = append(destinationFlights, flight)
		}
	}
	return destinationFlights
}

func printFlights(flights []Flight) {
	for i, flight := range flights {
		fmt.Printf("%d. %s - Rp%d (Keberangkatan: %s, Kedatangan: %s)\n", i+1, flight.Airline, flight.Price, flight.Departure, flight.Arrival)
	}
}

func insertionSort(flights []Flight, less func(i, j int) bool) {
	for i := 1; i < len(flights); i++ {
		j := i
		for j > 0 && less(j, j-1) {
			flights[j], flights[j-1] = flights[j-1], flights[j]
			j--
		}
	}
}
