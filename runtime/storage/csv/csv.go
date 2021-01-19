package csv

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"

	"github.com/yakawa/simpleDB/runtime/storage/table"
)

func Read(fn string) (*table.TableValue, error) {
	tbl := &table.TableValue{}

	f, err := os.Open(fn)
	if err != nil {
		return tbl, errors.New("Reading Error")
	}
	defer f.Close()

	ri := bufio.NewReader(f)

	line := ""

	for ln := 0; ; ln++ {
		l, cont, err := ri.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return tbl, err
		}
		line += string(l)

		if !cont {
			if ln == 0 {
				cols := []string{}
				if strings.HasPrefix(line, "#") {
					line = line[1:]
				}
				cols, err = splitColumn(line)
				if err != nil {
					return tbl, err
				}
				tbl.Header = cols
			} else {
				if !strings.HasPrefix(line, "#") {
					colsValue, _ := splitColumn(line)
					lineValue := map[string]table.ColumnValue{}
					for n, c := range colsValue {
						c = strings.Trim(c, " ")
						iv, err := strconv.ParseInt(c, 10, 64)
						if err != nil {
							return tbl, err
						}
						t := table.ColumnValue{
							Name: tbl.Header[n],
							Value: table.Value{
								Type:     table.Integer,
								Integral: int(iv),
							},
						}
						lineValue[tbl.Header[n]] = t
					}
					tbl.Values = append(tbl.Values, lineValue)
				}
			}

			line = ""
		}
	}
	return tbl, nil
}

func splitColumn(line string) ([]string, error) {
	cols := []string{}

	header := []rune(line)

	col := []rune("")
	inQuote := false
	for n, ch := range header {
		if ch == ',' && !inQuote {
			cols = append(cols, string(col))
			col = []rune("")
			continue
		}
		if ch == '"' && !inQuote {
			inQuote = true
			continue
		} else if ch == '"' && inQuote {
			if n != len(header)-1 {
				if header[n+1] != ',' {
					return cols, errors.New("Header is broken")
				}
			}
			inQuote = false
			continue
		}
		if ch == '\\' {
			if n == len(header)-1 {
				return cols, errors.New("Unexpected Terminated")
			}
			if header[n+1] != '"' && header[n+1] != '\\' {
				return cols, errors.New("Unknown Escaped Character")
			}
			continue
		}
		col = append(col, ch)
	}
	cols = append(cols, strings.Trim(string(col), " "))
	return cols, nil
}
