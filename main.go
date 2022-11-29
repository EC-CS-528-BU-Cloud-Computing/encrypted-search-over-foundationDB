package main

import (
    "github.com/apple/foundationdb/bindings/go/src/fdb"
    "log"
    "fmt"
)
/*	hello world func for FDB/go API

func main() {
    // Different API versions may expose different runtime behaviors.
    fdb.MustAPIVersion(720)

    // Open the default database from the system cluster
    db := fdb.MustOpenDefault()

    // Database reads and writes happen inside transactions
    ret, e := db.Transact(func(tr fdb.Transaction) (interface{}, error) {
        tr.Set(fdb.Key("hello"), []byte("world"))
        return tr.Get(fdb.Key("foo")).MustGet(), nil
        // db.Transact automatically commits (and if necessary,
        // retries) the transaction
    })
    if e != nil {
        log.Fatalf("Unable to perform FDB transaction (%v)", e)
    }

    fmt.Printf("hello is now world, foo was: %s\n", string(ret.([]byte)))
}
*/

// will organize project in a neater fashion with further understanding of go project structure
// incomplete atm, add FDB go API code 

type FDBService struct {
     FILE_MAX_SIZE  int = 10 * 1024 // this value will be different/unnessessary... needed in java impl
	 STR_LENGTH  string = "LENGTH"
     DOWNLOAD_PATH  string = ".\\src\\test\\downloads\\"
}

func FDBService_1() * FDBService {
    var this * FDBService =  &FDBService{};
    return this;
}

func (this *FDBService) listFiles( dirPath*  File)* java.util.List {
    var  Path*  java.util.List = Arrays.asList(dirPath.listFiles());
    var  files*  java.util.List = []();
    for  index_, file :=  range Path {
        if (file.isFile()) {
            files= append(files,file);
        } else {
            files.addAll(this.listFiles(file));
        }
    }
    return files;
}

func (this *FDBService) DownloadAFile( fileName  string) {
    var buffer byte
    var  index  int = 0;
    
	// FDB/go API code here
	// open data base, transact
	// error handling 
}

func (this *FDBService) UploadAll( path  string)bool {
    var  root*  File =  java.io.File(path);
    var  files*  java.util.List = this.listFiles(root);
    var  flag  bool = true;
    for  index_, file :=  range files {
        flag &= this.UploadAFile(file);
    }
    return flag;
}

func (this *FDBService) ClearAll( path  string)bool {
    var  root*  File =  java.io.File(path);
    var  files*  java.util.List = this.listFiles(root);
    var  flag  bool = true;
    for  index_, file :=  range files {
        flag &= this.ClearAFile(file);
    }
    return flag;

}
func (this *FDBService) ClearAFile( file*  File)bool {
    var  fileName  string = file.getName();
    var  index  int = 0;

	// using FDB Go API, open data base, transact
	// error handling 
    

}
func (this *FDBService) UploadAFile( file*  File)bool {
    var buffer [this.FILE_MAX_SIZE]byte
    var  fileName  string = file.getName();
    var  index  int = 0;

	// FDB/go API code here
	// open data base, transact
	// error handling 
}