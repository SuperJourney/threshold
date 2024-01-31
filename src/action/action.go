package action

type IdentIFace interface {
}

type ByteIdent struct {
}

type PareIdent struct {
	IdentIFace
}

type RootIdent interface {
	IdentIFace
}

type ActionPare struct {
	Op    string
	Left  IdentIFace
	Right IdentIFace
}

func Parse(path string) int {
	var root Ident
	var pareI = 0  // 括号计数
	var identI = 0 // 标识符计数

	l := len(path)
	for i := 0; i < l; {
		switch path[i] {
		case '(':
			pareI++
		case ')':
			pareI--
		}
	}
	return 0
}

func main() {
	path := "a ,(a,b),c"
	ParsePare([]byte(path))
}

// 解析括号里面的内容
func ParsePare(path []byte) (IdentIFace, int) {
	var left []byte
	var l = len(path)
	for i := 0; i < l; i++ {
		switch path[i] {
		case '(':
		case ')':
		case ',':
		case '|':
		default:
			left = append(left, path[i])
		}
	}
	return nil, -1
}

/*

(a | b )

pare(
	[
		op = |
		left = (a)
		right = (b)
	]
)

(a | (b ,a ))


pare(
	[
		op = |
		left = (a)
		right = pare(
			[
				op = ,
				left = (b)
				right = (a)
			]
		)
	]
)

}
*/
