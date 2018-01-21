package lowrank

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

type Recommender struct {
	UserLatent  *mat.Dense
	MovieLatent *mat.Dense
	Rating      *mat.Dense
}

func NewRecommender(R *mat.Dense, K int) *Recommender {
	I, J := R.Dims()
	return &Recommender{
		UserLatent:  RandMat(I, K),
		MovieLatent: RandMat(J, K),
		Rating:      R,
	}
}

func (r *Recommender) ModelPredict() (*mat.Dense, error) {
	I, KI := r.UserLatent.Dims()
	J, KJ := r.MovieLatent.Dims()

	if KI != KJ {
		return nil, mat.ErrShape
	}

	result := mat.NewDense(I, J, nil)
	result.Mul(r.UserLatent, r.MovieLatent.T())
	return result, nil
}

func (r *Recommender) Loss(reg float64) (float64, float64, error) {
	prediction, err := r.ModelPredict()
	if err != nil {
		return 0, 0, err
	}

	I, J := prediction.Dims()
	diff := mat.NewDense(I, J, nil)
	diff.Sub(prediction, r.Rating)
	accuracy := AbsAverage(diff)
	diff.MulElem(diff, diff)

	loss := 0.5 * mat.Sum(diff)

	USquared := mat.DenseCopyOf(r.UserLatent)
	USquared.MulElem(USquared, USquared)
	loss += reg * mat.Sum(USquared) / 2.0

	MSquared := mat.DenseCopyOf(r.MovieLatent)
	MSquared.MulElem(MSquared, MSquared)
	loss += reg * mat.Sum(MSquared) / 2.0

	return loss, accuracy, nil
}

func (r *Recommender) Gradients(reg float64) (*mat.Dense, *mat.Dense, error) {
	prediction, err := r.ModelPredict()
	if err != nil {
		return nil, nil, err
	}

	I, J := prediction.Dims()
	GradR := mat.NewDense(I, J, nil)
	GradR.Sub(prediction, r.Rating)

	_, K := r.UserLatent.Dims()

	GradU := mat.NewDense(I, K, nil)
	GradU.Mul(GradR, r.MovieLatent)
	RegU := mat.NewDense(I, K, nil)
	RegU.Scale(reg, r.UserLatent)
	GradU.Add(GradU, RegU)

	GradM := mat.NewDense(J, K, nil)
	GradM.Mul(GradR.T(), r.UserLatent)
	RegM := mat.NewDense(J, K, nil)
	RegM.Scale(reg, r.MovieLatent)
	GradM.Add(GradM, RegM)

	return GradU, GradM, nil
}

func (r *Recommender) Train(steps int, epochSize int, reg float64, learnRate float64) {
	I, _ := r.UserLatent.Dims()
	J, _ := r.MovieLatent.Dims()
	for step := 0; step < steps; step += 1 {
		if step%epochSize == 0 {
			loss, accuracy, _ := r.Loss(reg)
			fmt.Printf("%d: net loss %v, avg loss %v, and accuracy %v \n", step, loss, loss/float64(I*J), accuracy)
		}

		if GradU, GradM, err := r.Gradients(reg); err == nil {
			GradU.Scale(learnRate, GradU)
			r.UserLatent.Sub(r.UserLatent, GradU)

			GradM.Scale(learnRate, GradM)
			r.MovieLatent.Sub(r.MovieLatent, GradM)
		}
	}
}
