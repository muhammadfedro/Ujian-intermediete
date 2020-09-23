package models

import	"database/sql"
import	"ujian/entities"

type MhsModel struct{
	Db *sql.DB
}

func (m MhsModel) Deletes(id int64) (int64,error){
	result, err := m.Db.Exec("delete from mahasiswa where id = ?",id)
if err != nil{
	return 0, err
}else{
	return result.RowsAffected()
}
}


func (m MhsModel) Updates(maha entities.Mahasiswa) (int64,error){
	result, err := m.Db.Exec("update mahasiswa set nama=?, alamat=?, status=? where id = ?",
	maha.Name, maha.Alamat, maha.Status, maha.Id)
if err != nil{
	return 0, err
}else{
	return result.RowsAffected()
}
}



func (m MhsModel) Creates(maha *entities.Mahasiswa) error{
	result, err := m.Db.Exec("insert into mahasiswa(id,Nama,alamat,status) values (?,?,?,?)", maha.Id,maha.Name,maha.Alamat,maha.Status)
if err != nil{
	return err
}else{
	maha.Id,_=result.LastInsertId()
	return nil
}
}


func (m MhsModel) Finds(Id int64) (entities.Mahasiswa, error){
	rows, err := m.Db.Query("select * from mahasiswa where id = ?", Id)
	if err != nil{
		return entities.Mahasiswa{}, err
	}else{
		matkul := entities.Mahasiswa{}
		for rows.Next(){
			var id int64
			var name string
			var alamat string
			var status string
			err2 := rows.Scan(&id , &name, &alamat, &status)
			if err2 != nil{
				return entities.Mahasiswa{}, err2
			}else{
				matkul = entities.Mahasiswa{id,name, alamat, status}
			}
		}
		return matkul,nil
	}
}

func (m MhsModel) Findalls() ([]entities.Mahasiswa, error){
	rows, err := m.Db.Query("select * from mahasiswa")
	if err != nil{
		return nil, err
	}else{
		mahasiswas := []entities.Mahasiswa{}
		for rows.Next(){
			var id int64
			var name string
			var alamat string
			var status string
			err2 := rows.Scan(&id , &name, &alamat, &status)
			if err2 != nil{
				return nil, err2
			}else{
				maha := entities.Mahasiswa{id,name, alamat, status}
				mahasiswas = append(mahasiswas, maha)
			}
		}
		return mahasiswas,nil
	}
}