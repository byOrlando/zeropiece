package push

type LevelOptionParams struct {
	active        string `default:"active"`
	timeSensitive string `default:"timeSensitive"`
	passive       string `default:"passive"`
}

func (n *LevelOptionParams) init() {
	n.active = "active"
	n.timeSensitive = "timeSensitive"
	n.passive = "passive"
}

func (n *LevelOptionParams) Get(num int) string {
	switch num {
	case 0:
		return n.active
	case 1:
		return n.timeSensitive
	case 2:
		return n.passive
	default:
		return n.active
	}
}

func InitLevelOptionParams() *LevelOptionParams {
	opt := LevelOptionParams{}
	opt.init()
	return &opt
}
