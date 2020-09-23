package main
import "strings"
import "fmt"
import "ujian/config"
import "ujian/models"
import "ujian/entities"
import "os"
import "os/exec"
import "runtime"
func main()  {
	fmt.Println("=========SELAMAT DATANG========")
Tampilmenu()
}


func Tampilmenu(){
	var userChoice string
	
	fmt.Println("       Menu Matakuliah         ")
	fmt.Println("01. Tambah MataKuliah")
	fmt.Println("02. Tampilkan MataKuliah")
	fmt.Println("03. Cari MataKuliah")
	fmt.Println("04. Update MataKuliah")
	fmt.Println("05. Hapus MataKuliah")
	fmt.Println("       Menu Mahasiswa         ")
	fmt.Println("06. Tambah Mahasiswa")
	fmt.Println("07. Tampilkan Mahasiswa")
	fmt.Println("08. Cari Mahasiswa")
	fmt.Println("09. Update Mahasiswa")
	fmt.Println("10. Hapus Mahasiswa")
	fmt.Print("Pilihan Anda:")
	fmt.Scan(&userChoice)
		switch {
		case userChoice == "01":
			Menucreate()
		case userChoice == "02":
			Menufindall()
		case userChoice == "03":
			Menufind()
		case userChoice == "04":
			Menuupdate()
		case userChoice == "05":
			Menudelete()
		case userChoice == "06":
			Menucreates()
		case userChoice == "07":
			Menufindalls()
		case userChoice == "08":
			Menufinds()
		case userChoice == "09":
			Menuupdates()
		case userChoice == "10":
			Menudeletes()
		default:
			fmt.Println("Unknown Menu Code")
		}
}



func Menudelete()  {
	db, err := config.GetMySQLDB()
	if err != nil{
		fmt.Println(err)
	}else{
		matkulModel := models.MatakuliahModel{
			Db: db,
		}
		var hapus int64
		ConsoleClear()
		fmt.Println("Masukan Id:")
		fmt.Scan(&hapus)
		rows, err := matkulModel.Delete(hapus)
		if err != nil{
			fmt.Println(err)
		}else{
			if rows > 0{
				ConsoleClear()
				fmt.Println("Data dengan Id-",hapus," Telah Terhapus")
				Tampilmenu()
			}
		}
		}
	}


func Menuupdate()  {
	db, err := config.GetMySQLDB()
	if err != nil{
		fmt.Println(err)
	}else{
		matkulModel := models.MatakuliahModel{
			Db: db,
		}
		var id int64
		var name string
		ConsoleClear()
		fmt.Print("Masukan Id:")
		fmt.Scan(&id)
		fmt.Print("Masukan Nama:")
		fmt.Scan(&name)
		matkul := entities.Matkul{
			Id: id,
			Name: name,
		}
		rows, err := matkulModel.Update(matkul)
		if err != nil{
			fmt.Println(err)
		}else{
			if rows > 0{
				ConsoleClear()
				fmt.Println("Data Id-",id,"Telah Terupdate")
				Tampilmenu()
			}
		}
		}
	}

func Menucreate()  {
	db, err := config.GetMySQLDB()
	if err != nil{
		fmt.Println(err)
	}else{
		matkulModel := models.MatakuliahModel{
			Db: db,
		}
		var name string
		ConsoleClear()
		fmt.Print("Masukan Mata Kuliah:")
		fmt.Scan(&name)
		matkul := entities.Matkul{
			Name: name,
		}
		err := matkulModel.Create(&matkul)
		if err != nil{
			fmt.Println(err)
		}else{
			ConsoleClear()
			fmt.Println("Lastest Id",matkul.Id)
			fmt.Println("Sukses menyimpan data")
			Tampilmenu()
		}
		}
	}


func Menufind()  {
	db, err := config.GetMySQLDB()
	if err != nil{
		fmt.Println(err)
	}else{
		matkulModel := models.MatakuliahModel{
			Db: db,
		}
		var cari int64
		ConsoleClear()
		fmt.Print("Masukan Id:")
		fmt.Scan(&cari)
		matkul, err := matkulModel.Find(cari)
		if err != nil{
			fmt.Println(err)
		}else{
			ConsoleClear()
			fmt.Printf("%s\n", strings.Repeat("*", 30))
			fmt.Printf("%23s\n", "Hasil Pencarian")
			fmt.Printf("%s\n", strings.Repeat("*", 30))
			fmt.Printf("%-5d%-10s\n",matkul.Id,matkul.Name)
			Tampilmenu()
			
		}
	}
}

func Menufindall()  {
	db, err := config.GetMySQLDB()
	if err != nil{
		fmt.Println(err)
	}else{
		matkulModel := models.MatakuliahModel{
			Db: db,
		}
		matkuls, err := matkulModel.Findall()
		if err != nil{
			fmt.Println(err)
		}else{
			ConsoleClear()
			fmt.Printf("%s\n", strings.Repeat("=", 30))
			fmt.Printf("%26s\n", "Daftar Mata Kuliah")
			fmt.Printf("%s\n", strings.Repeat("=", 30))
			fmt.Printf("%-5s%-10s\n", "ID", "Mata Kuliah")
			fmt.Printf("%s\n", strings.Repeat("=", 30))
			for _, matkul := range matkuls{
				fmt.Printf("%-5d%-10s\n",matkul.Id,matkul.Name)
			}
			fmt.Printf("%s\n", strings.Repeat("=", 30))
			Tampilmenu()
		}
	}
}


func ConsoleClear() {
	currOS := runtime.GOOS
	switch currOS {
	case "windows":
		cmd := exec.Command("cmd", "/c", "cls")
		exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	default:
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}


func Menudeletes()  {
	db, err := config.GetMySQLDB()
	if err != nil{
		fmt.Println(err)
	}else{
		mhsModel := models.MhsModel{
			Db: db,
		}
		var hapus int64
		ConsoleClear()
		fmt.Println("Masukan Id:")
		fmt.Scan(&hapus)
		rows, err := mhsModel.Deletes(hapus)
		if err != nil{
			fmt.Println(err)
		}else{
			if rows > 0{
				ConsoleClear()
				fmt.Println("Data denga Id-",hapus," Telah Terhapus")
				Tampilmenu()
			}
		}
		}
	}


func Menuupdates()  {
	db, err := config.GetMySQLDB()
	if err != nil{
		fmt.Println(err)
	}else{
		mhsModel := models.MhsModel{
			Db: db,
		}
		var id int64
		var name string
		var alamat string
		var status string
		ConsoleClear()
		fmt.Print("Masukan Id:")
		fmt.Scan(&id)
		fmt.Print("Masukan Nama:")
		fmt.Scan(&name)
		fmt.Print("Masukan Alamat:")
		fmt.Scan(&alamat)
		fmt.Print("Masukan Status:")
		fmt.Scan(&status)


		mhs := entities.Mahasiswa{
			Id: id,
			Name: name,
			Alamat: alamat,
			Status: status,
		}
		rows, err := mhsModel.Updates(mhs)
		if err != nil{
			fmt.Println(err)
		}else{
			if rows > 0{
				ConsoleClear()
				fmt.Println("Data Id-",id,"Telah Terupdate")
				Tampilmenu()
			}
		}
		}
	}

func Menucreates()  {
	db, err := config.GetMySQLDB()
	if err != nil{
		fmt.Println(err)
	}else{
		mhsModel := models.MhsModel{
			Db: db,
		}
		var name string
		var alamat string
		var status string
		ConsoleClear()
		fmt.Print("Masukan Nama:")
		fmt.Scan(&name)
		fmt.Print("Masukan Alamat:")
		fmt.Scan(&alamat)
		fmt.Print("Masukan Status:")
		fmt.Scan(&status)

		mhs := entities.Mahasiswa{
			Name: name,
			Alamat: alamat,
			Status: status,
		}
		err := mhsModel.Creates(&mhs)
		if err != nil{
			fmt.Println(err)
		}else{
			ConsoleClear()
			fmt.Println("Lastest Id",mhs.Id)
			fmt.Println("Sukses menyimpan data")
			Tampilmenu()
		}
		}
	}


func Menufinds()  {
	db, err := config.GetMySQLDB()
	if err != nil{
		fmt.Println(err)
	}else{
		mhsModel := models.MhsModel{
			Db: db,
		}
		var cari int64
		ConsoleClear()
		fmt.Print("Masukan Id:")
		fmt.Scan(&cari)
		maha, err := mhsModel.Finds(cari)
		if err != nil{
			fmt.Println(err)
		}else{
			ConsoleClear()
			fmt.Printf("%s\n", strings.Repeat("*", 30))
			fmt.Printf("%23s\n", "Hasil Pencarian")
			fmt.Printf("%s\n", strings.Repeat("*", 30))
			fmt.Printf("%-5d%-10s%-10s%-10s\n",maha.Id,maha.Name,maha.Alamat,maha.Status)
			Tampilmenu()
			
		}
	}
}

func Menufindalls()  {
	db, err := config.GetMySQLDB()
	if err != nil{
		fmt.Println(err)
	}else{
		mhsModel := models.MhsModel{
			Db: db,
		}
		mahas, err := mhsModel.Findalls()
		if err != nil{
			fmt.Println(err)
		}else{
			ConsoleClear()
			fmt.Printf("%s\n", strings.Repeat("=", 30))
			fmt.Printf("%26s\n", "Daftar Mata Kuliah")
			fmt.Printf("%s\n", strings.Repeat("=", 30))
			fmt.Printf("%-5s%-10s%-10s%-10s\n", "ID", "Mahasiswa","Alamat","Status")
			fmt.Printf("%s\n", strings.Repeat("=", 30))
			for _, maha := range mahas{
				fmt.Printf("%-5d%-10s%-10s%-10s\n",maha.Id,maha.Name,maha.Alamat,maha.Status)
			}
			fmt.Printf("%s\n", strings.Repeat("=", 30))
			Tampilmenu()
		}
	}
}