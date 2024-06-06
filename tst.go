package main // Menyatakan bahwa ini adalah paket utama yang akan dieksekusi

import (
	"fmt" // Mengimpor paket fmt untuk input/output
	"math/rand"
	"strings" // Mengimpor paket strings untuk manipulasi string
	"time"    // Mengimpor paket time untuk manajemen waktu
)

type Flight struct {
	Airline     string
	Price       int
	Departure   string
	Arrival     string
	FlightTime  string
	Destination string
}

type Booking struct {
	Name        string
	Flight      Flight
	Destination string
}

var bookings []Booking
var flights []Flight
var status, password string

func main() {
	fmt.Println("English or spanish?")
	fmt.Println("Masukan Jenis Akses(Admin/User)")
	fmt.Scan(&status)

	rand.Seed(time.Now().UnixNano())
	randomizeFlights(&flights, 500)

	for {
		choice := displayMenu(&status)
		if choice == 3 {
			fmt.Println("Terima kasih telah menggunakan layanan kami.")
			break
		} else if choice == 5 {
			addFlight(&flights)
			continue
		} else if choice == 4 {
			sortFlightsMenu(&flights)
			continue
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
			insertionSort(&destinationFlights, func(i, j int) bool {
				return destinationFlights[i].Price < destinationFlights[j].Price
			})
		case 2:
			insertionSort(&destinationFlights, func(i, j int) bool {
				return destinationFlights[i].Price > destinationFlights[j].Price
			})
		case 3:
			insertionSort(&destinationFlights, func(i, j int) bool {
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
		fmt.Printf("Detail penerbangan yang dipilih:\nMaskapai: %s\nHarga: Rp%d\nKeberangkatan: %s\nKedatangan: %s\nDurasi: %s\nTujuan: %s\n",
			flight.Airline, flight.Price, flight.Departure, flight.Arrival, flight.FlightTime, flight.Destination)

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

func displayMenu(status *string) int {
	var choice int
	fmt.Println("Menu:")
	fmt.Println("1. Penerbangan Domestik")
	fmt.Println("2. Penerbangan Internasional")
	fmt.Println("3. Keluar")
	fmt.Println("4. Sortir Penerbangan")
	if *status == "Admin" {
		fmt.Println("5. Tambahkan Penerbangan")
	}
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

func sortFlightsMenu(flights *[]Flight) {
	fmt.Println("Pilihan Sorting:")
	fmt.Println("1. Harga Termurah ke Termahal")
	fmt.Println("2. Harga Termahal ke Termurah")
	fmt.Println("3. Waktu Keberangkatan")
	fmt.Print("Pilih metode sorting: ")
	var sortChoice int
	fmt.Scan(&sortChoice)
	switch sortChoice {
	case 1:
		insertionSort(flights, func(i, j int) bool {
			return (*flights)[i].Price < (*flights)[j].Price
		})
	case 2:
		insertionSort(flights, func(i, j int) bool {
			return (*flights)[i].Price > (*flights)[j].Price
		})
	case 3:
		insertionSort(flights, func(i, j int) bool {
			return (*flights)[i].Departure < (*flights)[j].Departure
		})
	default:
		fmt.Println("Pilihan tidak valid.")
		return
	}
	fmt.Println("Daftar penerbangan setelah disortir:")
	printFlights(*flights)
	time.Sleep(3 * time.Second)
}

func addFlight(flights *[]Flight) {
	var airline, departure, arrival, flightTime, destination string
	var price int

	fmt.Print("Masukkan nama maskapai: ")
	fmt.Scan(&airline)

	fmt.Print("Masukkan harga: ")
	fmt.Scan(&price)

	fmt.Print("Masukkan waktu keberangkatan (HH:MM): ")
	fmt.Scan(&departure)

	fmt.Print("Masukkan waktu kedatangan (HH:MM): ")
	fmt.Scan(&arrival)

	fmt.Print("Masukkan durasi penerbangan (xh): ")
	fmt.Scan(&flightTime)

	fmt.Print("Masukkan tujuan: ")
	fmt.Scan(&destination)

	newFlight := Flight{
		Airline:     airline,
		Price:       price,
		Departure:   departure,
		Arrival:     arrival,
		FlightTime:  flightTime,
		Destination: destination,
	}
	*flights = append(*flights, newFlight)

	fmt.Println("Penerbangan berhasil ditambahkan.")
	time.Sleep(2 * time.Second)
}

func insertionSort(flights *[]Flight, less func(i, j int) bool) {
	for i := 1; i < len(*flights); i++ {
		for j := i; j > 0 && less(j, j-1); j-- {
			(*flights)[j], (*flights)[j-1] = (*flights)[j-1], (*flights)[j]
		}
	}
}

func getFlights(destination string) []Flight {
	return flights
}

func searchFlights(destination string) []Flight {
	allFlights := getFlights(destination)
	var destinationFlights []Flight
	for i := 0; i < len(allFlights); i++ {
		flight := allFlights[i]
		if strings.Contains(strings.ToLower(flight.Destination), strings.ToLower(destination)) {
			destinationFlights = append(destinationFlights, flight)
		}
	}
	return destinationFlights
}

func printFlights(flights []Flight) {
	for i, flight := range flights {
		fmt.Printf("%d. Maskapai: %s, Harga: Rp%d, Keberangkatan: %s, Kedatangan: %s, Durasi: %s, Tujuan: %s\n", i+1, flight.Airline, flight.Price, flight.Departure, flight.Arrival, flight.FlightTime, flight.Destination)
	}
}

func randomizeFlights(flights *[]Flight, n int) {
	airlines := []string{"Garuda Indonesia", "Lion Air", "AirAsia", "Singapore Airlines", "Qatar Airways", "Cathay Pacific", "citilink", "Batik Air", "Emirates Airways", "British Airways", "Wings Air"}
	destinations := []string{"Jakarta", "Bali", "Surabaya", "Singapore", "Kuala Lumpur", "Bangkok", "Tokyo", "Seoul", "Hong Kong", "Dubai", "London", "New York", "Manchester", "Munich", "Milan", "Manado", "Bandung", "Surabaya", "Timika"}
	times := []string{"06:00", "07:00", "08:00", "09:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22.00", "23.00", "00:00", "01.00", "02.00", "03.00", "04.00", "05.00"}
	durations := []string{"1h", "2h", "3h", "4h", "5h", "6h", "7h", "8h", "9h", "10h", "11h", "12h"}

	for i := 0; i < n; i++ {
		airline := airlines[rand.Intn(len(airlines))]
		price := rand.Intn(20000000) + 500000 // Harga antara 500.000 sampai 2.500.000
		departure := times[rand.Intn(len(times))]
		arrival := times[rand.Intn(len(times))]
		flightTime := durations[rand.Intn(len(durations))]
		destination := destinations[rand.Intn(len(destinations))]

		newFlight := Flight{
			Airline:     airline,
			Price:       price,
			Departure:   departure,
			Arrival:     arrival,
			FlightTime:  flightTime,
			Destination: destination,
		}
		*flights = append(*flights, newFlight)
	}
}
