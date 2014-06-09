package mandel

type Data struct {
	MaxX, MaxY int

	CenterRe, CenterIm       float64
	Hd, Vd                   float64
	StepX, StepY             float64
	UpperLeftRe, UpperLeftIm float64
}

func NewData(maxX, maxY int, centerRe, centerIm, hd float64) *Data {
	d := new(Data)

	d.MaxX = maxX
	d.MaxY = maxY

	d.CenterRe = centerRe
	d.CenterIm = centerIm

	d.Hd = hd
	d.Vd = float64(maxY) / float64(maxX) * hd

	d.StepX = d.Hd / float64(maxX)
	d.StepY = d.Vd / float64(maxY)

	d.UpperLeftRe = centerRe - (float64(maxX/2) * d.StepX)
	d.UpperLeftIm = centerIm - (float64(maxY/2) * d.StepY)

	return d
}
