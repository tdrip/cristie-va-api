package client

func TestError(t *Testing) {

	bytes := []byte{}
	err := GetError(bytes)
	if err != nil {
		t.Fatalf("%v", err.Error())
	}
}
