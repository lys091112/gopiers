package jvm

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
	"os"
)


// https://zserge.com/posts/jvm/

type loader struct {
	r   io.Reader
	err error
}

func (l *loader) bytes(n int) []byte {
	b := make([]byte, n, n)

	if l.err == nil {
		_, l.err = io.ReadFull(l.r, b)
	}
	return b
}

func (l *loader) u1() uint8 {
	return l.bytes(1)[0]
}
func (l *loader) u2() uint16 {
	return binary.BigEndian.Uint16(l.bytes(2))
}
func (l *loader) u4() uint32 {
	return binary.BigEndian.Uint32(l.bytes(4))
}
func (l *loader) u8() uint64 {
	return binary.BigEndian.Uint64(l.bytes(8))
}

// Const 常量标识
type Const struct {
	Tag              byte
	NameIndex        uint16
	ClassIndex       uint16
	NameAndTypeIndex uint16
	StringIndex      uint16
	DescIndex        uint16
	String           string
}

// ConstPool 常量池
type ConstPool []Const

func (l *loader) cpinfo() (constPool ConstPool) {
	constPoolCount := l.u2()
	fmt.Println("constPoolCount=", constPoolCount)
	// Valid constant pool indices start from 1
	for i := uint16(1); i < constPoolCount; i++ {
		c := Const{Tag: l.u1()}
		switch c.Tag {
		case 0x01: // UTF8 string literal, 2 bytes length + data
			c.String = string(l.bytes(int(l.u2())))
		case 0x07: // Class index
			c.NameIndex = l.u2()
		case 0x08: // String reference index
			c.StringIndex = l.u2()
		case 0x09, 0x0a: // Field and method: class index + NaT index
			c.ClassIndex = l.u2()
			c.NameAndTypeIndex = l.u2()
		case 0x0c: // Name-and-type
			c.NameIndex, c.DescIndex = l.u2(), l.u2()
		default:
			l.err = fmt.Errorf("unsupported tag: %d", c.Tag)
		}
		constPool = append(constPool, c)
	}
	return constPool
}

// Resolve 获取制定下标的自变量
func (cp ConstPool) Resolve(index uint16) string {
	if cp[index-1].Tag == 0x01 {
		return cp[index-1].String
	}
	return ""
}

func (l *loader) interfaces(cp ConstPool) (interfaces []string) {
	interfacesCount := l.u2()

	for i := uint16(0); i < interfacesCount; i++ {
		interfaces = append(interfaces, cp.Resolve(l.u2()))
	}
	return interfaces
}

// Field type is used for both, fields and methods
type Field struct {
	Flags      uint16
	Name       string
	Descriptor string
	Attributes []Attribute
}

// Attribute contain addition information about fields and classes
// The most useful is "Code" attribute, which contains actual byte code
type Attribute struct {
	Name string
	Data []byte
}

//
func (l *loader) fields(cp ConstPool) (fields []Field) {
	fieldCount := l.u2()
	for i := uint16(0); i < fieldCount; i++ {
		fields = append(fields, Field{
			Flags:      l.u2(),
			Name:       cp.Resolve(l.u2()),
			Descriptor: cp.Resolve(l.u2()),
			Attributes: l.attrs(cp),
		})
	}
	return fields
}

func (l *loader) attrs(cp ConstPool) (attr []Attribute) {
	attributeCount := l.u2()
	for i := uint16(0); i < attributeCount; i++ {
		attr = append(attr, Attribute{
			Name: cp.Resolve(l.u2()),
			Data: l.bytes(int(l.u4())),
		})
	}
	return attr
}

// Class java类文件包含内容
type Class struct {
	ConstPool  ConstPool
	Name       string
	Super      string
	Flags      uint16
	Interfaces []string
	Fields     []Field
	Methods    []Field
	Attributes []Attribute
}

// Frame 执行框架
type Frame struct {
	Class  Class
	IP     uint32
	Code   []byte
	Locals []interface{}
	Stack  []interface{}
}

// Frame 生成执行框架
func (c Class) Frame(method string, args ...interface{}) Frame {
	for _, m := range c.Methods {
		if m.Name == method {
			for _, a := range m.Attributes {
				if a.Name == "Code" && len(a.Data) > 8 {
					maxLocals := binary.BigEndian.Uint16(a.Data[2:4])
					frame := Frame{
						Class:  c,
						Code:   a.Data[8:],
						Locals: make([]interface{}, maxLocals, maxLocals),
					}
					for i := 0; i < len(args); i++ {
						frame.Locals[i] = args[i]
					}
					return frame
				}
			}
		}
	}
	panic("method not found")
}

// Exec 模拟执行
func Exec(f Frame) interface{} {
	for {
		op := f.Code[f.IP]
		fmt.Printf("OP:%02x STACK:%v\n", op, f.Stack)
		n := len(f.Stack)
		switch op {
		case 26: // iload_0
			f.Stack = append(f.Stack, f.Locals[0])
		case 27: // iload_1
			f.Stack = append(f.Stack, f.Locals[1])
		case 96:
			a := f.Stack[n-1].(int32)
			b := f.Stack[n-2].(int32)
			f.Stack[n-2] = a + b
			f.Stack = f.Stack[:n-1]
		case 172: // ireturn
			v := f.Stack[n-1]
			f.Stack = f.Stack[:n-1]
			return v
		}
		f.IP++
	}
}

func parse(path string) {
	f, _ := os.Open(path)

	c := Class{}

	loader := &loader{r: f}

	cafebabe := loader.u4()
	major := loader.u2()
	minor := loader.u2()

	fmt.Printf("Flag:%x,major:%d,minor:%d\n", cafebabe, major, minor)

	cp := loader.cpinfo()

	c.ConstPool = cp

	c.Flags = loader.u2()
	c.Name = cp.Resolve(loader.u2())
	c.Super = cp.Resolve(loader.u2())

	c.Interfaces = loader.interfaces(cp)
	c.Fields = loader.fields(cp)
	c.Methods = loader.fields(cp)
	c.Attributes = loader.attrs(cp)

	if loader.err == nil {
		fmt.Printf("%+v\n", c)
		b, _ := json.Marshal(c)
		fmt.Println(string(b))
	} else {
		fmt.Println("error:", loader.err)
	}

	frame := c.Frame("add", int32(2), int32(3))

	r := Exec(frame)
	fmt.Printf("result:%d\n", r.(int32))

}
