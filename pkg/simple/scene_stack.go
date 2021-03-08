package simple

type SceneStack []Scene

func (s *SceneStack) Push(scene Scene) {
	*s = append(*s, scene)
}

func (s SceneStack) Current() Scene {
	return s[len(s)-1]
}

func (s *SceneStack) Pop() {
	if len(*s) == 1 {
		return
	}
	*s = (*s)[:len(*s)-1]
}

func (s SceneStack) Replace(scene Scene) {
	s[len(s)-1] = scene
}

func (s SceneStack) Render() (Widget, error) {
	return s.Current().Render()
}
