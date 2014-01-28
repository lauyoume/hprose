/**********************************************************\
|                                                          |
|                          hprose                          |
|                                                          |
| Official WebSite: http://www.hprose.com/                 |
|                   http://www.hprose.net/                 |
|                   http://www.hprose.org/                 |
|                                                          |
\**********************************************************/
/**********************************************************\
 *                                                        *
 * hprose/client_test.go                                  *
 *                                                        *
 * hprose Client Test for Go.                             *
 *                                                        *
 * LastModified: Jan 28, 2014                             *
 * Author: Ma Bingyao <andot@hprfc.com>                   *
 *                                                        *
\**********************************************************/

package hprose

import (
	"fmt"
	"reflect"
	"testing"
	"time"
)

type testUser struct {
	Name     string
	Sex      int
	Birthday time.Time
	Age      int
	Married  bool
}

func TestClient(t *testing.T) {
	client := NewClient("http://www.hprose.com/example/")
	var r1 chan string
	if err := client.Invoke("hello", []interface{}{"world"}, nil, &r1); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(<-r1)

	var r2 chan int
	if err := client.Invoke("sum", []interface{}{1, 2, 3, 4, 5, 6, 7}, nil, &r2); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(<-r2)

	var r3 chan map[string]string
	m := make(map[string]string)
	m["Jan"] = "January"
	m["Feb"] = "February"
	m["Mar"] = "March"
	m["Apr"] = "April"
	m["May"] = "May"
	m["Jun"] = "June"
	m["Jul"] = "July"
	m["Aug"] = "August"
	m["Sep"] = "September"
	m["Oct"] = "October"
	m["Nov"] = "November"
	m["Dec"] = "December"
	if err := client.Invoke("swapKeyAndValue", []interface{}{&m}, &InvokeOptions{ByRef: true}, &r3); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(m)
	fmt.Println(<-r3)
	fmt.Println(m)

	ClassManager.Register(reflect.TypeOf(testUser{}), "User")
	var r4 chan []testUser
	if err := client.Invoke("getUserList", []interface{}{}, nil, &r4); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(<-r4)

	var r5 chan []byte
	if err := client.Invoke("hello", []interface{}{"马秉尧"}, &InvokeOptions{ResultMode: Serialized}, &r5); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(string(<-r5))
	if err := client.Invoke("hello", []interface{}{"马秉尧"}, &InvokeOptions{ResultMode: Raw}, &r5); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(string(<-r5))
	if err := client.Invoke("hello", []interface{}{"马秉尧"}, &InvokeOptions{ResultMode: RawWithEndTag}, &r5); err != nil {
		t.Error(err.Error())
	}
	fmt.Println(string(<-r5))

}