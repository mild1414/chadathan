package integers

func TestAdder(.t *testing.T){
	sum := Add(2,2)
	expected := 4

	if sum != expected{
		t.Errorf()
	}
}
