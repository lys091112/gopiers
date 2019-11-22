package basic

import "testing"

func TestToJson(t *testing.T) {

	s := Student{StudentId: "1", StudentName: "name", StudentTeacher: "lily", StudentClass: "china"}
	t.Log(toJson(&s))
}
