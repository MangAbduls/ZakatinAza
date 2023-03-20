package main

import (
	"fmt"
	"image/color"
	"strconv"
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/canvas"
	"fyne.io/fyne/container"
	"fyne.io/fyne/widget"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

// @name		ZakatinAza
// @description	Penghitung Zakat
// @author		Bagas Kalih Putra - Zhafif Naufal Setyawan
// @verssion	1.2

func main() {
	a := app.New()
	wStart := a.NewWindow("ZakatinAza")
	wStart.Resize(fyne.NewSize(400, 400))
	r, _ := fyne.LoadResourceFromPath("zakat.png")
	wStart.SetIcon(r)

	btnStart := widget.NewButton("START", func() {
		w := a.NewWindow("ZakatinAza")
		w.Resize(fyne.NewSize(400, 100))
		w.SetMaster()
		r, _ := fyne.LoadResourceFromPath("zakat.png")
		w.SetIcon(r)

		label := widget.NewLabel("Jenis-Jenis Zakat")
		label.Alignment = fyne.TextAlignCenter

		btn1 := widget.NewButton("Zakat Fitrah", func() {
			w2 := a.NewWindow("Zakat Fitrah")
			w2.Resize(fyne.NewSize(500, 10))
			w2.SetIcon(r)

			label := widget.NewLabel("Zakat yang harus dikeluarkan sebesar 1 Sha' atau 2,5 kg")

			btnExit := widget.NewButton("EXIT", func() {
				w2.Close()
			})

			w2.SetContent(
				container.NewVBox(
					label,
					btnExit,
				))

			w2.Show()
		})
		btn2 := widget.NewButton("Zakat Mal", func() {
			w2 := a.NewWindow("Zakat Mal")
			w2.Resize(fyne.NewSize(500, 300))
			w2.SetIcon(r)

			w2.SetContent(widget.NewLabel("Berikut Jenis-Jenis Zakat Mal"))

			labelMal := widget.NewLabel("Jenis-Jenis Zakat Mal")
			labelMal.Alignment = fyne.TextAlignCenter

			btnEmas := widget.NewButton("Emas", func() {
				w3 := a.NewWindow("Zakat Mal - Emas")
				w3.Resize(fyne.NewSize(500, 500))
				w3.SetIcon(r)

				w3.SetContent(widget.NewLabel("Zakat Mal - Emas"))

				labelEmas := widget.NewLabel("Emas(gram):")
				labelTahun := widget.NewLabel("Kepemilikan(bulan):")

				inputPerak := widget.NewEntry()
				inputPerak.SetPlaceHolder("Masukkan jumlah emas yang dimiliki dalam satuan gram...")

				labelMax := widget.NewLabel("max input : 9999999999999")

				inputTahun := widget.NewEntry()
				inputTahun.SetPlaceHolder("Masukkan durasi kepemilikan emas dalam satuan bulan...")

				labelKoma := widget.NewLabel("*gunakan '.' untuk desimal")

				result := canvas.NewText("", color.White)
				result.Alignment = fyne.TextAlignCenter
				result.TextSize = 20

				btn1 := widget.NewButton("Hitung Zakat", func() {
					perak, _ := strconv.ParseFloat(inputPerak.Text, 64)
					tahun, _ := strconv.ParseFloat(inputTahun.Text, 64)

					result.Text = malEmas(perak, tahun)
					result.Refresh()
				})

				btnExit := widget.NewButton("EXIT", func() {
					w3.Close()
				})

				w3.SetContent(
					container.NewVBox(
						labelEmas,
						inputPerak,
						labelMax,
						labelTahun,
						inputTahun,
						labelKoma,
						btn1,
						result,
						btnExit,
					))

				w3.Show()
			})

			btnPerak := widget.NewButton("Perak", func() {
				w3 := a.NewWindow("Zakat Mal - Perak")
				w3.Resize(fyne.NewSize(500, 500))
				w3.SetIcon(r)

				w3.SetContent(widget.NewLabel("Zakat Mal - Perak"))

				labelPerak := widget.NewLabel("Perak(gram):")
				labelTahun := widget.NewLabel("Kepemilikan(bulan):")

				inputPerak := widget.NewEntry()
				inputPerak.SetPlaceHolder("Masukkan jumlah perak yang dimiliki dalam satuan gram...")

				labelMax := widget.NewLabel("max input : 9999999999999")

				inputTahun := widget.NewEntry()
				inputTahun.SetPlaceHolder("Masukkan durasi kepemilikan perak dalam satuan bulan...")

				labelKoma := widget.NewLabel("*gunakan '.' untuk desimal")

				result := canvas.NewText("", color.White)
				result.Alignment = fyne.TextAlignCenter
				result.TextSize = 20

				btn1 := widget.NewButton("Hitung Zakat", func() {
					perak, _ := strconv.ParseFloat(inputPerak.Text, 64)
					tahun, _ := strconv.ParseFloat(inputTahun.Text, 64)

					result.Text = malPerak(perak, tahun)
					result.Refresh()
				})

				btnExit := widget.NewButton("EXIT", func() {
					w3.Close()
				})

				w3.SetContent(
					container.NewVBox(
						labelPerak,
						inputPerak,
						labelMax,
						labelTahun,
						inputTahun,
						labelKoma,
						btn1,
						result,
						btnExit,
					))

				w3.Show()
			})

			btnProfesi := widget.NewButton("Profesi", func() {
				w3 := a.NewWindow("Zakat Mal - Profesi")
				w3.Resize(fyne.NewSize(500, 500))
				w3.SetIcon(r)

				w3.SetContent(widget.NewLabel("Zakat Mal - Profesi"))

				labelPenghasilan := widget.NewLabel("Penghasilan(bulanan):")
				labelTahun := widget.NewLabel("Lama Bekerja(bulan):")

				inputPenghasilan := widget.NewEntry()
				inputPenghasilan.SetPlaceHolder("Masukkan penghasilan anda dalam 1 bulan...")

				labelBilBulat := widget.NewLabel("*hanya masukkan bilangan bulat untuk input di atas, max input : 99999999999999999")

				inputTahun := widget.NewEntry()
				inputTahun.SetPlaceHolder("Masukkan lama anda bekerja di profesi tersebut dalam satuan bulan...")

				labelKoma := widget.NewLabel("*gunakan '.' untuk desimal")

				result := canvas.NewText("", color.White)
				result.Alignment = fyne.TextAlignCenter
				result.TextSize = 20

				btn1 := widget.NewButton("Hitung Zakat", func() {
					penghasilan, _ := strconv.Atoi(inputPenghasilan.Text)
					tahun, _ := strconv.ParseFloat(inputTahun.Text, 64)

					result.Text = malPenghasilan(penghasilan, tahun)
					result.Refresh()
				})

				btnExit := widget.NewButton("EXIT", func() {
					w3.Close()
				})

				w3.SetContent(
					container.NewVBox(
						labelPenghasilan,
						inputPenghasilan,
						labelBilBulat,
						labelTahun,
						inputTahun,
						labelKoma,
						btn1,
						result,
						btnExit,
					))

				w3.Show()
			})

			btnPenjualan := widget.NewButton("Penjualan", func() {
				w3 := a.NewWindow("Zakat Mal - Penjualan")
				w3.Resize(fyne.NewSize(500, 500))
				w3.SetIcon(r)

				w3.SetContent(widget.NewLabel("Zakat Mal - Penjualan"))

				labelPenjualan := widget.NewLabel("Hasil Penjualan(bulanan):")
				labelTahun := widget.NewLabel("Lama Berjualan(bulan):")

				inputPenjualan := widget.NewEntry()
				inputPenjualan.SetPlaceHolder("Masukkan penjualan anda dalam 1 bulan...")

				labelBilBulat := widget.NewLabel("*hanya masukkan bilangan bulat untuk input di atas, max input : 99999999999999999")

				inputTahun := widget.NewEntry()
				inputTahun.SetPlaceHolder("Masukkan lama anda berjualan satuan bulan...")

				labelKoma := widget.NewLabel("*gunakan '.' untuk desimal")

				result := canvas.NewText("", color.White)
				result.Alignment = fyne.TextAlignCenter
				result.TextSize = 20

				btn1 := widget.NewButton("Hitung Zakat", func() {
					penjualan, _ := strconv.Atoi(inputPenjualan.Text)
					tahun, _ := strconv.ParseFloat(inputTahun.Text, 64)

					result.Text = malPenjualan(penjualan, tahun)
					result.Refresh()
				})

				btnExit := widget.NewButton("EXIT", func() {
					w3.Close()
				})

				w3.SetContent(
					container.NewVBox(
						labelPenjualan,
						inputPenjualan,
						labelBilBulat,
						labelTahun,
						inputTahun,
						labelKoma,
						btn1,
						result,
						btnExit,
					))

				w3.Show()
			})

			btnBertani := widget.NewButton("Pertanian", func() {
				w3 := a.NewWindow("Zakat Mal - Pertanian")
				w3.Resize(fyne.NewSize(500, 500))
				w3.SetIcon(r)

				w3.SetContent(widget.NewLabel("Zakat Mal - Pertanian"))

				label1 := widget.NewLabel("Jika hasil tani berkulit, maka zakat yang harus dikeluarkan sebesar 1.4 Ton")
				label2 := widget.NewLabel("Jika hasil tani tidak berkulit, maka zakat yang harus dikeluarkan sebesar 0.7 Ton")

				btnExit := widget.NewButton("EXIT", func() {
					w3.Close()
				})

				w3.SetContent(
					container.NewVBox(
						label1,
						label2,
						btnExit,
					))
				w3.Show()
			})

			btnPanen := widget.NewButton("Hasil Panen", func() {
				w3 := a.NewWindow("Zakat Mal - Hasil Panen")
				w3.Resize(fyne.NewSize(500, 500))
				w3.SetIcon(r)

				w3.SetContent(widget.NewLabel("Zakat Mal - Hasil Panen"))

				labelPanen := widget.NewLabel("Hasil Panen(kg):")

				inputPanen := widget.NewEntry()
				inputPanen.SetPlaceHolder("Masukkan hasil panen anda dalam satuan kilogram...")

				labelKoma := widget.NewLabel("*gunakan '.' untuk desimal, max input : 9999999999")

				radio := widget.NewRadioGroup([]string{"Ya", "Tidak"}, func(value string) {})

				labelDefault := widget.NewLabel("*default: Tidak")

				label1 := widget.NewLabel("Apakah anda menggunakan air irigasi untuk bertani?")

				result := canvas.NewText("", color.White)
				result.Alignment = fyne.TextAlignCenter
				result.TextSize = 20

				btn1 := widget.NewButton("Hitung Zakat", func() {
					penjualan, _ := strconv.ParseFloat(inputPanen.Text, 64)
					irigasi := radio.Selected
					result.Text = malPanen(penjualan, irigasi)
					result.Refresh()
				})

				btnExit := widget.NewButton("EXIT", func() {
					w3.Close()
				})

				w3.SetContent(
					container.NewVBox(
						labelPanen,
						inputPanen,
						labelKoma,
						label1,
						radio,
						labelDefault,
						btn1,
						result,
						btnExit,
					))

				w3.Show()
			})

			btnHartaKarun := widget.NewButton("Harta Karun/Rikaz", func() {
				w3 := a.NewWindow("Zakat Mal - Hasil Panen")
				w3.Resize(fyne.NewSize(500, 500))
				w3.SetIcon(r)

				w3.SetContent(widget.NewLabel("Zakat Mal - Hasil Panen"))

				labelRikaz := widget.NewLabel("Nilai Jual Rikaz:")

				inputRikaz := widget.NewEntry()
				inputRikaz.SetPlaceHolder("Masukkan nilai jual barang yang anda temukan...")

				labelBilBulat := widget.NewLabel("*hanya masukkan bilangan bulat untuk input di atas, max input : 99999999999999999")

				result := canvas.NewText("", color.White)
				result.Alignment = fyne.TextAlignCenter
				result.TextSize = 20

				btn1 := widget.NewButton("Hitung Zakat", func() {
					rikaz, _ := strconv.Atoi(inputRikaz.Text)

					result.Text = malHartaKarun(rikaz)
					result.Refresh()
				})

				btnExit := widget.NewButton("EXIT", func() {
					w3.Close()
				})

				w3.SetContent(
					container.NewVBox(
						labelRikaz,
						inputRikaz,
						labelBilBulat,
						btn1,
						result,
						btnExit,
					))

				w3.Show()
			})

			btnExit := widget.NewButton("EXIT", func() {
				w2.Close()
			})

			// SET CONTENTS WINDOW2 MAL
			w2.SetContent(
				container.NewVBox(
					labelMal,
					btnEmas,
					btnPerak,
					btnProfesi,
					btnPenjualan,
					btnBertani,
					btnPanen,
					btnHartaKarun,
					btnExit,
				))

			w2.Show()
		})
		btn3 := widget.NewButton("Zakat Ternak", func() {
			w2 := a.NewWindow("Zakat Ternak")
			w2.Resize(fyne.NewSize(500, 500))
			w2.SetIcon(r)

			labelTernak := widget.NewLabel("Jenis-Jenis Zakat Ternak")

			btnUnta := widget.NewButton("Unta", func() {
				w3 := a.NewWindow("Zakat Ternak - Unta")
				w3.Resize(fyne.NewSize(500, 500))
				w3.SetIcon(r)

				w3.SetContent(widget.NewLabel("Zakat Ternak - Unta"))

				labelUnta := widget.NewLabel("Jumlah Unta:")
				labelUmur := widget.NewLabel("Umur Unta(bulan):")

				inputUnta := widget.NewEntry()
				inputUnta.SetPlaceHolder("Masukkan jumlah unta yang anda miliki...")

				labelBilBulat := widget.NewLabel("*hanya masukkan bilangan bulat untuk input di atas")

				inputUmur := widget.NewEntry()
				inputUmur.SetPlaceHolder("Masukkan umur unta yang anda miliki...")

				labelKoma := widget.NewLabel("*gunakan '.' untuk desimal")

				result := canvas.NewText("", color.White)
				result.Alignment = fyne.TextAlignCenter
				result.TextSize = 20

				btnRes := widget.NewButton("Hitung Zakat", func() {
					unta, _ := strconv.Atoi(inputUnta.Text)
					umur, _ := strconv.ParseFloat(inputUmur.Text, 64)

					result.Text = ternakUnta(unta, umur)
					result.Refresh()
				})

				btnExit := widget.NewButton("EXIT", func() {
					w3.Close()
				})

				w3.SetContent(
					container.NewVBox(
						labelUnta,
						inputUnta,
						labelBilBulat,
						labelUmur,
						inputUmur,
						labelKoma,
						btnRes,
						result,
						btnExit,
					))

				w3.Show()

			})

			btnSapi := widget.NewButton("Sapi", func() {
				w3 := a.NewWindow("Zakat Ternak - Sapi")
				w3.Resize(fyne.NewSize(500, 500))
				w3.SetIcon(r)

				w3.SetContent(widget.NewLabel("Zakat Ternak - Sapi"))

				labelSapi := widget.NewLabel("Jumlah Sapi:")
				labelUmur := widget.NewLabel("Umur Sapi(bulan):")

				inputSapi := widget.NewEntry()
				inputSapi.SetPlaceHolder("Masukkan jumlah sapi yang anda miliki...")

				labelBilBulat := widget.NewLabel("*hanya masukkan bilangan bulat untuk input di atas")

				inputUmur := widget.NewEntry()
				inputUmur.SetPlaceHolder("Masukkan umur sapi yang anda miliki...")

				labelKoma := widget.NewLabel("*gunakan '.' untuk desimal")

				result := canvas.NewText("", color.White)
				result.Alignment = fyne.TextAlignCenter
				result.TextSize = 20

				btnRes := widget.NewButton("Hitung Zakat", func() {
					sapi, _ := strconv.Atoi(inputSapi.Text)
					umur, _ := strconv.ParseFloat(inputUmur.Text, 64)

					result.Text = ternakSapi(sapi, umur)
					result.Refresh()
				})

				btnExit := widget.NewButton("EXIT", func() {
					w3.Close()
				})

				w3.SetContent(
					container.NewVBox(
						labelSapi,
						inputSapi,
						labelBilBulat,
						labelUmur,
						inputUmur,
						labelKoma,
						btnRes,
						result,
						btnExit,
					))

				w3.Show()

			})

			btnKambing := widget.NewButton("Kambing", func() {
				w3 := a.NewWindow("Zakat Ternak - Kambing")
				w3.Resize(fyne.NewSize(500, 500))
				w3.SetIcon(r)

				w3.SetContent(widget.NewLabel("Zakat Ternak - Kambing"))

				labelKambing := widget.NewLabel("Jumlah Kambing:")
				labelUmur := widget.NewLabel("Umur Kambing(bulan):")

				inputKambing := widget.NewEntry()
				inputKambing.SetPlaceHolder("Masukkan jumlah kambing yang anda miliki...")

				labelBilBulat := widget.NewLabel("*hanya masukkan bilangan bulat untuk input di atas")

				inputUmur := widget.NewEntry()
				inputUmur.SetPlaceHolder("Masukkan umur kambing yang anda miliki...")

				labelKoma := widget.NewLabel("*gunakan '.' untuk desimal")

				result := canvas.NewText("", color.White)
				result.Alignment = fyne.TextAlignCenter
				result.TextSize = 20

				btnRes := widget.NewButton("Hitung Zakat", func() {
					kambing, _ := strconv.Atoi(inputKambing.Text)
					umur, _ := strconv.ParseFloat(inputUmur.Text, 64)

					result.Text = ternakKambing(kambing, umur)
					result.Refresh()
				})

				btnExit := widget.NewButton("EXIT", func() {
					w3.Close()
				})

				w3.SetContent(
					container.NewVBox(
						labelKambing,
						inputKambing,
						labelBilBulat,
						labelUmur,
						inputUmur,
						labelKoma,
						btnRes,
						result,
						btnExit,
					))

				w3.Show()

			})

			btnLainnya := widget.NewButton("Lainnya", func() {
				w3 := a.NewWindow("Zakat Ternak - Lainnya")
				w3.Resize(fyne.NewSize(500, 500))
				w3.SetIcon(r)

				w3.SetContent(widget.NewLabel("Zakat Ternak - Lainnya"))

				labelLainnya := widget.NewLabel("Jumlah Penghasilan Ternak:")
				labelBulan := widget.NewLabel("Durasi usaha ternak(bulan):")

				inputLainnya := widget.NewEntry()
				inputLainnya.SetPlaceHolder("Masukkan penghasilan bulanan anda ...")

				labelBilBulat := widget.NewLabel("*hanya masukkan bilangan bulat untuk input di atas")

				inputBulan := widget.NewEntry()
				inputBulan.SetPlaceHolder("Masukkan lama anda memulai usaha ternak...")

				labelKoma := widget.NewLabel("*gunakan '.' untuk desimal")

				result := canvas.NewText("", color.White)
				result.Alignment = fyne.TextAlignCenter
				result.TextSize = 20

				btnRes := widget.NewButton("Hitung Zakat", func() {
					penghasilan, _ := strconv.Atoi(inputLainnya.Text)
					bulan, _ := strconv.ParseFloat(inputBulan.Text, 64)

					result.Text = ternakLainnya(penghasilan, bulan)
					result.Refresh()
				})

				btnExit := widget.NewButton("EXIT", func() {
					w3.Close()
				})

				w3.SetContent(
					container.NewVBox(
						labelLainnya,
						inputLainnya,
						labelBilBulat,
						labelBulan,
						inputBulan,
						labelKoma,
						btnRes,
						result,
						btnExit,
					))

				w3.Show()

			})
			btnExit := widget.NewButton("EXIT", func() {
				w2.Close()
			})

			// SET CONTENTS WINDOW2 TERNAK
			w2.SetContent(
				container.NewVBox(
					labelTernak,
					btnUnta,
					btnSapi,
					btnKambing,
					btnLainnya,
					btnExit,
				))

			w2.Show()
		})

		btnExit := widget.NewButton("EXIT", func() {
			w.Close()
		})

		w.SetContent(
			container.NewVBox(
				label,
				btn1,
				btn2,
				btn3,
				btnExit,
			))
		w.Show()
	})

	imageJudul := canvas.NewImageFromFile("ZakatinAza.png")
	imageJudul.FillMode = canvas.ImageFillOriginal

	btnExit := widget.NewButton("EXIT", func() {
		wStart.Close()
	})

	wStart.SetContent(
		container.NewVBox(
			imageJudul,
			btnStart,
			btnExit,
		))
	wStart.Show()

	a.Run()
}

func malEmas(emas, tahun float64) string {
	if emas <= 0 || emas > 9999999999999 {
		return "invalid input"
	} else if emas < 85 {
		return "Anda belum wajib mengeluarkan zakat"
	} else {
		if tahun <= 0 {
			return "invalid input"
		} else if tahun < 12 {
			return "Anda belum wajib mengeluarkan zakat"
		} else {
			jumlahZakat := emas * 25 / 1000
			if jumlahZakat > 1000 {
				jumlahZakat /= 1000
				if jumlahZakat > 1000 {
					jumlahZakat /= 1000
					return fmt.Sprintf("Zakat yang harus dikeluarkan sebanyak %v ton", jumlahZakat)
				} else {
					return fmt.Sprintf("Zakat yang harus dikeluarkan sebanyak %v kg", jumlahZakat)
				}
			} else {
				return fmt.Sprintf("Zakat yang harus dikeluarkan sebanyak %v gram", jumlahZakat)
			}
		}
	}
}

func malPerak(perak, tahun float64) string {
	if perak <= 0 || perak > 9999999999999 {
		return "invalid input"
	} else if perak < 595 {
		return "Anda belum wajib mengeluarkan zakat"
	} else {
		if tahun <= 0 {
			return "invalid input"
		} else if tahun < 12 {
			return "Anda belum wajib mengeluarkan zakat"
		} else {
			jumlahZakat := perak * 25 / 1000
			if jumlahZakat > 1000 {
				jumlahZakat /= 1000
				if jumlahZakat > 1000 {
					jumlahZakat /= 1000
					return fmt.Sprintf("Zakat yang harus dikeluarkan sebanyak %v ton", jumlahZakat)
				} else {
					return fmt.Sprintf("Zakat yang harus dikeluarkan sebanyak %v kg", jumlahZakat)
				}
			} else {
				return fmt.Sprintf("Zakat yang harus dikeluarkan sebanyak %v gram", jumlahZakat)
			}
		}
	}
}

func malPenghasilan(penghasilan int, bulanKerja float64) string {
	if penghasilan <= 0 || penghasilan > 99999999999999999 {
		return "invalid input"
	} else if penghasilan < 2075000 {
		return "Anda belum wajib mengeluarkan zakat"
	} else {
		if bulanKerja <= 0 {
			return "invalid input"
		} else if bulanKerja < 12 {
			return "Anda belum wajib mengeluarkan zakat"
		} else {
			jumlahZakat := penghasilan * 25 / 1000
			gk := message.NewPrinter(language.Greek)
			var convNumber string = gk.Sprintf("%d", jumlahZakat)
			return fmt.Sprintf("Zakat yang harus dikeluarkan sebanyak Rp.%v", convNumber)
		}
	}
}

func malPenjualan(penjualan int, bulanJualan float64) string {
	if penjualan <= 0 || penjualan > 99999999999999999 {
		return "invalid input"
	} else if penjualan < 2075000 {
		return "Anda belum wajib mengeluarkan zakat"
	} else {
		if bulanJualan <= 0 {
			return "invalid input"
		} else if bulanJualan < 12 {
			return "Anda belum wajib mengeluarkan zakat"
		} else {
			jumlahZakat := penjualan * 25 / 1000
			gk := message.NewPrinter(language.Greek)
			var convNumber string = gk.Sprintf("%d", jumlahZakat)
			return fmt.Sprintf("Zakat yang harus dikeluarkan sebanyak Rp.%v", convNumber)
		}
	}
}

func malPanen(panen float64, irigasi string) string {
	if panen <= 0 || panen > 9999999999 {
		return "invalid input"
	} else if panen < 750 {
		return "Anda belum wajib  mengeluarkan zakat"
	} else {
		if irigasi == "Ya" {
			jumlahZakat := panen * 10 / 100
			if jumlahZakat > 1000 {
				jumlahZakat = jumlahZakat / 1000
				return fmt.Sprintf("Jumlah zakat yang harus dikeluarkan sebesar %v ton", jumlahZakat)
			} else {
				return fmt.Sprintf("Jumlah zakat yang harus dikeluarkan sebesar %v kg", jumlahZakat)
			}
		} else {
			jumlahZakat := panen * 5 / 100
			if jumlahZakat > 1000 {
				jumlahZakat = jumlahZakat / 1000
				return fmt.Sprintf("Jumlah zakat yang harus dikeluarkan sebesar %v ton", jumlahZakat)
			} else {
				return fmt.Sprintf("Jumlah zakat yang harus dikeluarkan sebesar %v kg", jumlahZakat)
			}
		}
	}
}

func malHartaKarun(jumlah int) string {
	sum := jumlah * 20 / 100
	gk := message.NewPrinter(language.Greek)
	var convNumber string = gk.Sprintf("%d", sum)
	if jumlah <= 0 || jumlah >= 99999999999999999 {
		return "invalid input"
	} else {
		return fmt.Sprintf("Zakat yang harus dikeluarkan sebesar Rp.%v", convNumber)
	}
}

func ternakUnta(unta int, umur float64) string {
	if umur <= 0 {
		return "invalid input"
	} else if umur < 12 {
		return ("Anda belum wajib mengeluarkan zakat")
	} else {
		if unta > 30 {
			return ("Zakat yang harus dikeluarkan sebanyak 1 ekor unta berumur 2 tahun")
		} else if unta > 24 {
			return ("Zakat yang harus dikeluarkan sebanyak 4 ekor kambing")
		} else if unta > 19 {
			return ("Zakat yang harus dikeluarkan sebanyak 3 ekor kambing")
		} else if unta > 14 {
			return ("Zakat yang harus dikeluarkan sebanyak 2 ekor kambing")
		} else if unta > 9 {
			return ("Zakat yang harus dikeluarkan sebanyak 1 ekor kambing")
		} else if unta > 1 {
			return ("Anda belum wajib mengeluarkan zakat")
		} else {
			return "invalid input"
		}
	}
}

func ternakSapi(sapi int, umur float64) string {
	if umur <= 0 {
		return "invalid input"
	} else if umur < 12 {
		return ("Anda belum wajib mengeluarkan zakat")
	} else {
		if sapi > 89 {
			return ("Zakat yang harus dikeluarkan sebanyak 3 ekor sapi berumur 1 tahun")
		} else if sapi > 79 {
			return ("Zakat yang harus dikeluarkan sebanyak 2 ekor sapi berumur 2 tahun")
		} else if sapi > 69 {
			return ("Zakat yang harus dikeluarkan sebanyak 2 ekor sapi berumur 1 tahun dan 2 tahun")
		} else if sapi > 59 {
			return ("Zakat yang harus dikeluarkan sebanyak 2 ekor sapi berumur 1 tahun")
		} else if sapi > 39 {
			return ("Zakat yang harus dikeluarkan sebanyak 1 ekor sapi berumur 2 tahun")
		} else if sapi > 29 {
			return ("Zakat yang harus dikeluarkan sebanyak 1 ekor sapi berumur 1 tahun")
		} else if sapi > 0 {
			return ("Anda belum wajib mengeluarkan zakat")
		} else {
			return "invalid input"
		}
	}
}

func ternakKambing(kambing int, umur float64) string {
	if umur <= 0 {
		return "invalid input"
	} else if umur < 12 {
		return ("Anda belum wajib mengeluarkan zakat")
	} else {
		if kambing > 299 {
			return ("Zakat yang harus dikeluarkan sebanyak 4 ekor kambing")
		} else if kambing > 200 {
			return ("Zakat yang harus dikeluarkan sebanyak 3 ekor kambing")
		} else if kambing > 120 {
			return ("Zakat yang harus dikeluarkan sebanyak 2 ekor kambing")
		} else if kambing > 39 {
			return ("Zakat yang harus dikeluarkan sebanyak 1 ekor kambing")
		} else if kambing > 0 {
			return ("Anda belum wajib mengeluarkan zakat")
		} else {
			return "invalid input"
		}
	}
}

func ternakLainnya(penghasilan int, bulan float64) string {
	if penghasilan <= 0 {
		return "invalid input"
	} else if penghasilan < 2075000 {
		return ("Anda belum wajib mengeluarkan zakat")
	} else {
		if bulan < 12 {
			return ("Anda belum wajib mengeluarkan zakat")
		} else {
			jumlahZakat := penghasilan * 25 / 1000
			gk := message.NewPrinter(language.Greek)
			var convNumber string = gk.Sprintf("%d", jumlahZakat)
			return fmt.Sprintf("Zakat yang harus dikeluarkan sebanyak Rp.%v", convNumber)
		}
	}
}
