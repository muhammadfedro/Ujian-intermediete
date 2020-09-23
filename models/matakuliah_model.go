package models

import	"database/sql"
import	"ujian/entities"

type MatakuliahModel struct{
	Db *sql.DB
}

func (mm MatakuliahModel) Delete(id int64) (int64,error){
	result, err := mm.Db.Exec("delete from matakuliah where id = ?",id)
if err != nil{
	return 0, err
}else{
	return result.RowsAffected()
}
}


func (mm MatakuliahModel) Update(matkul entities.Matkul) (int64,error){
	result, err := mm.Db.Exec("update matakuliah set name=? where id = ?",matkul.Name,matkul.Id)
if err != nil{
	return 0, err
}else{
	return result.RowsAffected()
}
}



func (mm MatakuliahModel) Create(matkul *entities.Matkul) error{
	result, err := mm.Db.Exec("insert into matakuliah(id,Name) values (?,?)", matkul.Id,matkul.Name)
if err != nil{
	return err
}else{
	matkul.Id,_=result.LastInsertId()
	return nil
}
}


func (mm MatakuliahModel) Find(Id int64) (entities.Matkul, error){
	rows, err := mm.Db.Query("select * from matakuliah where id = ?", Id)
	if err != nil{
		return entities.Matkul{}, err
	}else{
		matkul := entities.Matkul{}
		for rows.Next(){
			var id int64
			var name string
			err2 := rows.Scan(&id , &name)
			if err2 != nil{
				return entities.Matkul{}, err2
			}else{
				matkul = entities.Matkul{id,name}
			}
		}
		return matkul,nil
	}
}

func (mm MatakuliahModel) Findall() ([]entities.Matkul, error){
	rows, err := mm.Db.Query("select * from matakuliah")
	if err != nil{
		return nil, err
	}else{
		matakuliahs := []entities.Matkul{}
		for rows.Next(){
			var id int64
			var name string
			err2 := rows.Scan(&id , &name)
			if err2 != nil{
				return nil, err2
			}else{
				matkul := entities.Matkul{id,name}
				matakuliahs = append(matakuliahs, matkul)
			}
		}
		return matakuliahs,nil
	}
}