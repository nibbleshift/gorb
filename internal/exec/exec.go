package exec

import "os/exec"

type Exec struct {
	padding int
}

func (e *Exec) Run() {
	if shell {
		out, err := exec.Command("ls").Output()
		if err != nil {
			panic("some error found")
		}
		return out
	} else {
		out, err := exec.Command(cmd).Output()
		if err != nil {
			panic("some error found")
		}
		return out
	}
}
