package push

type IconOptionParams struct {
	zero  string `default:"https://shkqg.com/bark/0.png"`
	one   string `default:"https://shkqg.com/bark/1.png"`
	two   string `default:"https://shkqg.com/bark/2.png"`
	three string `default:"https://shkqg.com/bark/3.png"`
	five  string `default:"https://shkqg.com/bark/5.png"`
	six   string `default:"https://shkqg.com/bark/6.png"`
	seven string `default:"https://shkqg.com/bark/7.png"`
	eight string `default:"https://shkqg.com/bark/8.png"`
	nine  string `default:"https://shkqg.com/bark/9.png"`
}

func (n *IconOptionParams) init() {
	n.zero = "https://shkqg.com/bark/0.png"
	n.one = "https://shkqg.com/bark/1.png"
	n.two = "https://shkqg.com/bark/2.png"
	n.three = "https://shkqg.com/bark/3.png"
	n.five = "https://shkqg.com/bark/5.png"
	n.six = "https://shkqg.com/bark/6.png"
	n.seven = "https://shkqg.com/bark/7.png"
	n.eight = "https://shkqg.com/bark/8.png"
	n.nine = "https://shkqg.com/bark/9.png"
}
func (n *IconOptionParams) Get(num int) string {
	switch num {
	case 0:
		return n.zero
	case 1:
		return n.one
	case 2:
		return n.two
	case 3:
		return n.three
	case 5:
		return n.five
	case 6:
		return n.six
	case 7:
		return n.seven
	case 8:
		return n.eight
	case 9:
		return n.nine
	default:
		return n.zero
	}
}

func InitIconOptionParams() *IconOptionParams {
	opt := IconOptionParams{}
	opt.init()
	return &opt
}
