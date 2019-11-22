package basic

import "encoding/json"

type Student struct {
	StudentId      string `json:"sid"`
	StudentName    string `json:"sname"`
	StudentClass   string `json:"class"`
	StudentTeacher string `json:"teacher"`
}

func toJson(s *Student) string {
	res, _ := json.Marshal(*s)
	return string(res)
}

func toObj(str string) *Student {
	s := new(Student)
	json.Unmarshal([]byte(str), s)
	return s
}
