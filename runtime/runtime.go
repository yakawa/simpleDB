package runtime

import (
	"errors"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"sync"

	"github.com/yakawa/simpleDB/runtime/storage/table"
)

var instance *Runtime
var once sync.Once

func GetInstance() *Runtime {
	once.Do(func() {
		instance = &Runtime{
			localTables: make(map[string]map[string]string),
		}
	})
	return instance
}

type Runtime struct {
	localTableDir string
	localTables   map[string]map[string]string
}

func (r *Runtime) Set(t string) *Runtime {
	r.localTableDir = t
	r.readLocalTables("_")
	return r
}

func (r *Runtime) readLocalTables(db string) {
	r.localTables[db] = make(map[string]string)
	dbPath := ""
	if db != "_" {
		dbPath = filepath.Join(r.localTableDir, db)
	} else {
		dbPath = r.localTableDir
	}
	files, _ := ioutil.ReadDir(dbPath)
	for _, f := range files {
		if f.IsDir() {
			if db == "_" {
				r.readLocalTables(f.Name())
			}
			continue
		}
		fn := f.Name()
		fp := filepath.Join(r.localTableDir, fn)
		w := strings.Split(fn, ".")
		r.localTables["_"][w[0]] = fp
	}
}

func (r *Runtime) ReadLineFromLocalTable(db string, tbl string, fn func([]table.ColumnValue)) error {
	fp, exists := r.localTables[db][tbl]
	if !exists {
		return errors.New(fmt.Sprintf("Table (%s) Not Found", tbl))
	}

	return nil
}
