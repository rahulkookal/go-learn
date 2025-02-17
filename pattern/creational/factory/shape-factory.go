package factory

type shapeType string

const (
	CircleShape    shapeType = "Circle"
	RectangleShape shapeType = "Rectangle"
)

func GetShape(s shapeType) Shape {

	switch s {
	case CircleShape:
		return Circle{}
	case RectangleShape:
		return Rectangle{}
	default:
		return nil
	}

}
