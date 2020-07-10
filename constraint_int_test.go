package govalidator

import (
	"fmt"
	"reflect"
	"testing"
)

func TestFiledTag_describe(t *testing.T) {
	type sa struct {
		Age   int8   `min:"1000" max:"20" default:"10"`
		Score *int16 `min:"1" max:"200000" default:"99"`
	}

	a := &sa{Age: 50, Score: new(int16)}
	v := reflect.ValueOf(a).Elem()
	for i := 0; i < v.NumField(); i++ {
		fi := v.Type().Field(i)
		ds, er := describeInt(&fi)
		if er != nil {
			fmt.Println(er)
			fmt.Println(ds)
		} else {
			t.Fatal("fail to detect describe errors")
		}

		if ds.isSet() {
			val := v.Field(i)
			if es := ds.validate(&val, true); es != nil {
				fmt.Println(es)
			}

		}
	}
	fmt.Printf("Age: %d, Score: %d\n", a.Age, *a.Score)
	if a.Age != 20 {
		t.Errorf("fail to auto fix filed Age")
	}
	if *a.Score != 99 {
		t.Errorf("fail to auto fix filed Score")
	}
}
