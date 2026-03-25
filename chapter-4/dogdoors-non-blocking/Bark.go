package main

type Bark struct {
	sound string
}

func NewBark(sound string) *Bark {
	return &Bark{
		sound: sound,
	}
}

func (b *Bark) GetSound() string {
	return b.sound
}

func (b *Bark) Equals(other *Bark) bool {
	if other == nil {
		return false
	}

	return b.sound == other.GetSound()
}
