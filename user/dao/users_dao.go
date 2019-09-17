package dao

import(
	"database/sql"
  _ "github.com/go-sql-driver/mysql"
	"fmt"
	"../model"
	"../dto"
)
//	"github.com/mparaiso/simple-row-mapper-go"


func main(){
	db := openDB()
	defer db.Close()

	var version string
	db.QueryRow("1.0").Scan(&version)
	fmt.Println("Connected to:", version)
}

type userDao struct{
	sql *sql.DB
}

//"fluttershy:podipa@tcp(127.0.0.1:3306)/hospital"
func openDB() *sql.DB{
	db, err := sql.Open("mysql", "fluttershy:podipa@tcp(127.0.0.1:3306)/hospital") 
	if err != nil {
		panic(err)
	  }
	return db
}

func FindByID(id int) *model.User{
	db := openDB()
	user := &model.User{} // user
	id = user.ID //defining id param
	query := "SELECT idUser as ID, username as Username, email as Email FROM User WHERE id = ?;" //query
	ret := handleQueryRow(db, query, user, id) //return 
	defer db.Close()
	return ret
}

func FindByUsername(username string) *model.User{
	db := openDB()
	user := &model.User{} // user
	query := "SELECT idUser as ID, username as Username, email as Email FROM User WHERE username = ?;" //query
	username = user.Username //defining id param
	ret := handleQueryRow(db, query, user, username)//return
	defer db.Close()
	return ret
}

func FindAll() []*model.User{
	db := openDB()
	sql := "SELECT idUser as ID, username as Username, email as Email FROM User;"
	query, err := db.Query(sql)
	user := &model.User{} //defining user
	users := []*model.User{}
	ret := handleQuery(query, user.ID , err, users) // defining user list 
	defer db.Close()
	return ret
}

func Update(userDto *dto.UserDto){
	
}

func Save(createDto *dto.UserCreateDto) int{
	db := openDB()
	user := &model.User{} // user
	id := createDto.ID //defining id param
	createDto = &dto.UserCreateDto{}
	query := "INSERT INTO User(username, password) VALUES(?,?);" //query
	ret := handleQueryRow(db, query, user, createDto.Username, createDto.Password) //return 
	fmt.Println(ret)
	defer db.Close()
	return id
}

func Delete(id int){
	db := openDB()
	user := &model.User{} // user
	id = user.ID //defining id param
	query := "DELETE FROM User WHERE id = ?;" //query
	ret := handleQueryRow(db, query, user, id) //return 
	fmt.Println(ret)
	defer db.Close()
}


func handleQuery(query *sql.Rows, i interface{}, err error, iList []*model.User) []*model.User {
	for query.Next() {  // scaning query
		err = query.Scan(i) // accessing error from query, scaning query params
		if err == nil {
			panic(err) //error 
		}
		fmt.Println(iList)
	}
	err = query.Err() //accessing and defining err
	if err != nil {
    	panic(err)
	}
	return iList
}

func handleQueryRow(db *sql.DB, query string, user *model.User, i... interface{}) *model.User {
	db = openDB()
	err := db.QueryRow(query, i)
	scanErr := err.Scan(i)
	if err != nil {
		if scanErr == sql.ErrNoRows {
			fmt.Println("Zero rows found")
		} else {
			panic(err)
		}
	}
	return user
}
 

//db *sql.DB

/*for query.Next() {  // scaning query
		err = query.Scan(user.ID, user.Username) // accessing error from query, scaning query params
		if err == nil {
			panic(err) //error 
		}
		fmt.Println(user)
	}
	err = query.Err() //accessing and defining err
	if err != nil {
    	panic(err)
	}

	err := db.QueryRow(query, id).Scan(user.ID)// err scaning query, defining ? = id
	if err != nil {
			if err == sql.ErrNoRows {
				fmt.Println("Zero rows found")
			} else {
				panic(err)
			}
		}
	*/


