package password

type Hasher struct{}

func NewHasher() *Hasher {
	return &Hasher{}
}

func (hasher *Hasher) Hash(
	plainPassword string,
) (string, error) {
	return Hash(plainPassword)
}

func (hasher *Hasher) Verify(
	plainPassword string,
	encodedPassword string,
) (bool, error) {
	return Verify(plainPassword, encodedPassword)
}
