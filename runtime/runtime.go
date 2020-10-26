package runtime

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
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

	f, err := os.Open(fp)
	if err != nil {
		return errors.New("Reading Error")
	}
	defer f.Close()

	ri := bufio.NewReader(f)

	line := ""

	cols := []string{}

	for ln := 0; ; ln++ {
		l, cont, err := ri.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		line += string(l)

		if !cont {
			fmt.Printf("%s\n", line)

			if ln == 0 {
				if strings.HasPrefix(line, "#") {
					line = line[1:]
				}
				nameCols := strings.Split(line, ",")
				for _, name := range nameCols {
					cols = append(cols, name)
				}
			} else {
				if !strings.HasPrefix(line, "#") {
					colsValue := strings.Split(line, ",")
					lineValue := []table.ColumnValue{}
					for n, c := range colsValue {
						iv, err := strconv.ParseInt(c, 10, 64)
						if err != nil {
							return err
						}
						t := table.ColumnValue{
							Name: cols[n],
							Value: table.Value{
								Type:     table.Integer,
								Integral: int(iv),
							},
						}
						lineValue = append(lineValue, t)
					}
					fn(lineValue)
				}
			}

			line = ""
		}
	}
	return nil
}
