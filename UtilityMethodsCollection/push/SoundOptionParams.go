package push

type SoundOptionParams struct {
	choo      string `default:"choo"`
	alarm     string `default:"alarm"`
	birdSong  string `default:"birdsong"`
	telegraph string `default:"telegraph"`
	newMail   string `default:"newmail"`
	horn      string `default:"horn"`
	ladder    string `default:"ladder"`
}

func (n *SoundOptionParams) init() {
	n.choo = "choo"
	n.alarm = "alarm"
	n.birdSong = "birdsong"
	n.telegraph = "telegraph"
	n.newMail = "newmail"
	n.horn = "horn"
	n.ladder = "ladder"
}

func (n *SoundOptionParams) Get(num int) string {
	switch num {
	case 0:
		return n.choo
	case 1:
		return n.alarm
	case 2:
		return n.birdSong
	case 3:
		return n.telegraph
	case 4:
		return n.newMail
	case 5:
		return n.horn
	case 6:
		return n.ladder
	default:
		return n.choo
	}
}

func InitSoundOptionParams() *SoundOptionParams {
	opt := SoundOptionParams{}
	opt.init()
	return &opt
}
