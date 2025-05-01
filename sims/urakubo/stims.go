// Copyright (c) 2021 The Emergent Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package urakubo

import (
	"math/rand"

	"cogentcore.org/core/math32"
)

// Stims are the stimulus functions.
type Stims int32 //enums:enum

const (
	Baseline Stims = iota

	CaTarg

	ClampCa1

	GClamp

	STDP

	STDPSweep

	STDPPacketSweep

	Poisson

	SPoissonRGClamp

	PoissonHzSweep

	PoissonDurSweep

	OpPhaseDurSweep

	ThetaErr

	ThetaErrComp

	ThetaErrSweep

	ThetaErrAllSweep
)

// RGeStimForHzMap is the strength of GeStim G clamp to obtain a given R firing rate
var RGeStimForHzMap = map[int]float32{
	25:  .09,
	50:  .12,
	100: .15,
}

func RGeStimForHz(hz float32) float32 {
	var gel, geh, hzl, hzh float32
	switch {
	case hz <= 25:
		gel = 0
		geh = RGeStimForHzMap[25]
		hzl = 0
		hzh = 25
	case hz <= 50:
		gel = RGeStimForHzMap[25]
		geh = RGeStimForHzMap[50]
		hzl = 25
		hzh = 50
	case hz <= 100:
		gel = RGeStimForHzMap[50]
		geh = RGeStimForHzMap[100]
		hzl = 50
		hzh = 100
	default:
		gel = RGeStimForHzMap[100]
		geh = 2 * gel
		hzl = 100
		hzh = 200
	}
	return gel + ((hz-hzl)/(hzh-hzl))*(geh-gel)
}

// ClampCa1Ca is direct copy of Ca values from test_stdp.g genesis func
var ClampCa1Ca = []float64{
	509.987, 0.05731354654,
	509.990, 1.800978422,
	509.994, 3.778658628,
	509.999, 3.385097265,
	510.000, 3.192493439,
	510.001, 3.484202623,
	510.002, 9.131223679,
	510.003, 7.309978008,
	510.006, 3.479086161,
	510.010, 2.207912683,
	510.015, 1.591691375,
	510.021, 1.139062405,
	510.029, 0.8029100895,
	510.040, 0.5597535968,
	510.052, 0.3869290054,
	510.066, 0.2666117251,
	510.083, 0.1854287386,
	510.103, 0.1331911981,
	510.126, 0.1012685224,
	510.151, 0.08303561062,
	510.180, 0.07339032739,
	510.213, 0.06863268465,
	510.249, 0.06637191772,
	510.289, 0.06523273885,
	510.333, 0.06451437622,
	510.382, 0.06390291452,
	510.435, 0.06327681988,
	510.492, 0.062598221,
	510.555, 0.06186179072,
	510.623, 0.06107418239,
	510.696, 0.06024680287,
	510.774, 0.05939326808,
}

var ClampVm = []float64{
	16.990, -64.62194824,
	16.994, -63.22241211,
	16.999, -62.93297958,
	17.001, -62.99502563,
	17.002, -58.25670624,
	17.003, -4.396664619,
	17.004, -40.63706589,
	17.007, -55.61388779,
	17.010, -60.13838196,
	17.015, -61.69696426,
	17.022, -62.90755844,
	17.030, -63.82915497,
	17.040, -64.40870667,
}

// PerMsec returns per msec for given input data
func PerMsec(orig []float64) []float64 {
	ost := orig[0]
	nca := len(orig) / 2
	oet := orig[(nca-1)*2]
	dur := oet - ost
	dms := int(dur / 0.001)
	rdt := make([]float64, dms)
	si := 0
	mxi := 0
	for i := 0; i < dms; i++ {
		ct := ost + float64(i)*0.001
		st := orig[si*2]
		et := orig[(si+1)*2]
		sca := orig[si*2+1]
		eca := orig[(si+1)*2+1]
		if ct > et {
			si++
			if si >= nca-1 {
				break
			}
			st = orig[si*2]
			et = orig[(si+1)*2]
			sca = orig[si*2+1]
			eca = orig[(si+1)*2+1]
		}
		mxi = i
		pt := (ct - st) / (et - st)
		ca := sca + pt*(eca-sca)
		rdt[i] = ca
		// fmt.Printf("%d \tct:  %g  \tca:  %g  \tst:  %g  \tet:  %g  \tsca:  %g \teca:  %g\n", i, ct, ca, st, et, sca, eca)
	}
	return rdt[:mxi+1]
}

func (uk *Urakubo) BaselineFun() {
	for msec := 0; msec < 500000; msec++ { // 500000 = 500 sec for full baseline
		uk.NeuronUpdt(msec, 0, 0)
		uk.StatsDefault(0)
		if uk.StopNow() {
			break
		}
	}
	uk.Spine.InitCode()
	uk.Stopped()
}

func (uk *Urakubo) CaTargFun() {
	uk.Spine.Ca.SetBuffTarg(uk.CaTarg.Cyt, uk.CaTarg.PSD)
	for msec := 0; msec < 20000; msec++ {
		uk.NeuronUpdt(msec, 0, 0)
		uk.StatsDefault(0)
		if uk.StopNow() {
			break
		}
	}
	uk.Stopped()
}

func (uk *Urakubo) ClampCa1Fun() {
	cas := PerMsec(ClampCa1Ca)
	nca := len(cas)
	bca := 0.05
	for msec := 0; msec < 20000; msec++ {
		tms := (msec + 500) % 1000
		ca := bca
		if tms < nca {
			ca = cas[tms]
		}
		cca := bca + ((ca - bca) / 3)
		uk.Spine.Ca.SetClamp(cca, ca)
		uk.NeuronUpdt(msec, 0, 0)
		uk.StatsDefault(0)
		if uk.StopNow() {
			break
		}
	}
	uk.GraphRun(uk.FinalSecs, 0)
	uk.Stopped()
}

func (uk *Urakubo) STDPFun() {
	toff := 500
	dur := 1
	psms := toff + 5 - uk.DeltaT // 5 is lag
	tott := uk.NReps * 1000

	for msec := 0; msec < tott; msec++ {
		ims := msec % 1000
		if ims == psms {
			uk.Spine.States.PreSpike = 1
		} else {
			uk.Spine.States.PreSpike = 0
		}
		ge := float32(0.0)
		if ims >= toff && ims < toff+dur {
			ge = uk.GeStim
		}
		uk.NeuronUpdt(msec, ge, 0)
		uk.StatsDefault(0)
		if uk.StopNow() {
			break
		}
	}
	uk.GraphRun(uk.FinalSecs, 0)
	uk.Stopped()
}

func (uk *Urakubo) STDPSweepFun() {
	toff := 500
	dur := 1
	tott := uk.NReps * 1000

	uk.ResetDWtPlot()

	for dt := -uk.DeltaTRange; dt <= uk.DeltaTRange; dt += uk.DeltaTInc {
		psms := toff + 5 - dt // 5 is lag
		uk.ResetTimePlots()
		uk.Init()

		for msec := 0; msec < tott; msec++ {
			ims := msec % 1000
			if ims == psms {
				uk.Spine.States.PreSpike = 1
			} else {
				uk.Spine.States.PreSpike = 0
			}
			ge := float32(0.0)
			if ims >= toff && ims < toff+dur {
				ge = uk.GeStim
			}
			uk.NeuronUpdt(msec, ge, 0)
			uk.StatsDefault(0)
			if uk.StopNow() {
				uk.Stopped()
				return
			}
		}
		uk.GraphRun(uk.FinalSecs, 0)
		uk.StatsDWt(uk.Stats("DWtStats"), float64(dt), 0)
		uk.Plot("DWtPlot").GoUpdate()
	}

	uk.Stopped()
}

// STDPPacket runs a sequence of Dur pre-post spike packets with sweep of
// pre-post offset in < 1/2 SendHz ISI range, with ISI interval between packets, N reps,
// and varying the frequency of pre-post firing (X axis).
func (uk *Urakubo) STDPPacketSweepFun() {
	isi := int(1000.0 / uk.SendHz)
	hisi := isi / 2
	dr := hisi - 5 // allow for lag

	uk.ResetDWtPlot()

	for dt := -dr; dt <= dr; dt++ {
		rms := hisi
		sms := hisi + 5 - dt // 5 is lag
		uk.ResetTimePlots()
		uk.Init()

		for ri := 0; ri < uk.NReps; ri++ {
			for msec := 0; msec < uk.DurMsec; msec++ {
				ims := msec % isi
				if ims == sms {
					uk.Spine.States.PreSpike = 1
				} else {
					uk.Spine.States.PreSpike = 0
				}
				ge := float32(0.0)
				if ims == rms {
					ge = uk.GeStim
				}
				uk.NeuronUpdt(msec, ge, 0)
				uk.StatsDefault(0)
				if uk.StopNow() {
					uk.Stopped()
					return
				}
			}
		}
		uk.GraphRun(uk.FinalSecs, 0)
		uk.StatsDWt(uk.Stats("DWtStats"), float64(dt), float64(uk.SendHz))
		uk.Plot("DWtPlot").GoUpdate()
	}

	uk.Stopped()
}

func (uk *Urakubo) PoissonFun() {
	Sint := math32.Exp(-1000.0 / uk.SendHz)
	Rint := math32.Exp(-1000.0 / uk.RecvHz)

	tmsec := 0
	for ri := 0; ri < uk.NReps; ri++ {
		Sp := float32(1)
		Rp := float32(1)

		for msec := 0; msec < uk.DurMsec; msec++ {
			Sp *= rand.Float32()
			if Sp <= Sint {
				uk.Spine.States.PreSpike = 1
				Sp = 1
			} else {
				uk.Spine.States.PreSpike = 0
			}

			ge := float32(0.0)
			Rp *= rand.Float32()
			if Rp <= Rint {
				ge = uk.GeStim
				Rp = 1
			}

			uk.NeuronUpdt(tmsec, ge, 0)
			uk.StatsDefault(0)
			if uk.StopNow() {
				break
			}
			tmsec++
		}
		uk.Spine.States.PreSpike = 0
		uk.GraphRun(uk.ISISec, 0)
	}
	uk.GraphRun(uk.FinalSecs, 0)
	uk.Stopped()
}

func (uk *Urakubo) SPoissonRGClampFun() {
	Sint := math32.Exp(-1000.0 / uk.SendHz)

	for ri := 0; ri < uk.NReps; ri++ {
		Sp := float32(1)

		for msec := 0; msec < uk.DurMsec; msec++ {
			Sp *= rand.Float32()
			if Sp <= Sint {
				uk.Spine.States.PreSpike = 1
				Sp = 1
			} else {
				uk.Spine.States.PreSpike = 0
			}

			uk.NeuronUpdt(msec, uk.GeStim, 0)
			uk.StatsDefault(0)
			if uk.StopNow() {
				break
			}
		}
		uk.Spine.States.PreSpike = 0
		uk.GraphRun(uk.ISISec, 0)
	}
	uk.GraphRun(uk.FinalSecs, 0)
	uk.Stopped()
}

func (uk *Urakubo) PoissonHzSweepFun() {
	uk.ResetDWtPlot()

	for shz := 10; shz <= 100; shz += 10 {
		for rhz := 10; rhz <= 100; rhz += 10 {
			Sint := math32.Exp(-1000.0 / float32(shz))
			Rint := math32.Exp(-1000.0 / float32(rhz))

			uk.ResetTimePlots()
			uk.Init()
			for ri := 0; ri < uk.NReps; ri++ {
				Sp := float32(1)
				Rp := float32(1)

				for msec := 0; msec < uk.DurMsec; msec++ {
					Sp *= rand.Float32()
					if Sp <= Sint {
						uk.Spine.States.PreSpike = 1
						Sp = 1
					} else {
						uk.Spine.States.PreSpike = 0
					}

					ge := float32(0.0)
					Rp *= rand.Float32()
					if Rp <= Rint {
						ge = uk.GeStim
						Rp = 1
					}

					uk.NeuronUpdt(msec, ge, 0)
					uk.StatsDefault(0)
					if uk.StopNow() {
						uk.Stopped()
						return
					}
				}
				uk.Spine.States.PreSpike = 0
				uk.GraphRun(uk.ISISec, 0)
			}
			uk.GraphRun(uk.FinalSecs, 0)
			uk.StatsDWt(uk.Stats("DWtStats"), float64(rhz), float64(shz))
			uk.Plot("DWtPlot").GoUpdate()
		}
	}
	uk.Stopped()
}

func (uk *Urakubo) PoissonDurSweepFun() {
	uk.ResetDWtPlot()

	for dur := 200; dur <= 1000; dur += 100 {
		for rhz := 10; rhz <= 100; rhz += 10 {
			Sint := math32.Exp(-1000.0 / float32(uk.SendHz))
			Rint := math32.Exp(-1000.0 / float32(rhz))

			uk.ResetTimePlots()
			uk.Init()
			for ri := 0; ri < uk.NReps; ri++ {
				Sp := float32(1)
				Rp := float32(1)

				for msec := 0; msec < dur; msec++ {
					Sp *= rand.Float32()
					if Sp <= Sint {
						uk.Spine.States.PreSpike = 1
						Sp = 1
					} else {
						uk.Spine.States.PreSpike = 0
					}

					ge := float32(0.0)
					Rp *= rand.Float32()
					if Rp <= Rint {
						ge = uk.GeStim
						Rp = 1
					}

					uk.NeuronUpdt(msec, ge, 0)
					uk.StatsDefault(0)
					if uk.StopNow() {
						uk.Stopped()
						return
					}
				}
				uk.Spine.States.PreSpike = 0
				uk.GraphRun(uk.ISISec, 0)
			}
			uk.GraphRun(uk.FinalSecs, 0)
			uk.StatsDWt(uk.Stats("DWtStats"), float64(rhz), float64(dur))
			uk.Plot("DWtPlot").GoUpdate()
		}
	}
	uk.Stopped()
}

// OpPhase runs sending, recv in opposite phases (half interval off at start)
// This is what was used in the original XCAL Dwt function derivation in Genesis model
func (uk *Urakubo) OpPhaseDurSweepFun() {
	uk.ResetDWtPlot()

	for dur := 200; dur <= 1000; dur += 100 {
		for rhz := 10; rhz <= 100; rhz += 10 {
			Sint := 1000.0 / float32(uk.SendHz)
			Rint := 1000.0 / float32(rhz)

			uk.ResetTimePlots()
			uk.Init()
			for ri := 0; ri < uk.NReps; ri++ {
				Sp := Sint / 2
				Rp := Rint

				for msec := 0; msec < dur; msec++ {
					fms := float32(msec)
					if fms-Sp >= Sint {
						uk.Spine.States.PreSpike = 1
						Sp = fms
					} else {
						uk.Spine.States.PreSpike = 0
					}

					ge := float32(0.0)
					if fms-Rp >= Rint {
						ge = uk.GeStim
						Rp = fms
					}

					uk.NeuronUpdt(msec, ge, 0)
					uk.StatsDefault(0)
					if uk.StopNow() {
						uk.Stopped()
						return
					}
				}
				uk.Spine.States.PreSpike = 0
				uk.GraphRun(uk.ISISec, 0)
			}
			uk.GraphRun(uk.FinalSecs, 0)
			uk.StatsDWt(uk.Stats("DWtStats"), float64(rhz), float64(dur))
			uk.Plot("DWtPlot").GoUpdate()
		}
	}
	uk.Stopped()
}

func (uk *Urakubo) ThetaErrFun() {
	phsdur := []int{uk.DurMsec / 2, uk.DurMsec / 2}
	nphs := len(phsdur)

	// using send, recv for minus, plus
	sphz := []int{int(uk.SendHz), int(uk.RecvHz)}
	rphz := []int{int(uk.SendHz), int(uk.RecvHz)}

	tmsec := 0
	uk.ResetTimePlots()
	uk.Init()
	uk.RunQuiet(10)
	for ri := 0; ri < uk.NReps; ri++ {
		Sp := float32(1)
		Rp := float32(1)
		for pi := 0; pi < nphs; pi++ {
			dur := phsdur[pi]
			shz := sphz[pi]
			rhz := rphz[pi]
			Sint := math32.Exp(-1000.0 / float32(shz))
			Rint := math32.Exp(-1000.0 / float32(rhz))
			for msec := 0; msec < dur; msec++ {
				Sp *= rand.Float32()
				if Sp <= Sint {
					uk.Spine.States.PreSpike = 1
					Sp = 1
				} else {
					uk.Spine.States.PreSpike = 0
				}

				ge := float32(0.0)
				Rp *= rand.Float32()
				if Rp <= Rint {
					ge = uk.GeStim
					Rp = 1
				}
				if uk.RGClamp {
					ge = RGeStimForHz(float32(rhz))
				}

				uk.NeuronUpdt(tmsec, ge, 0)
				uk.StatsDefault(0)
				if uk.StopNow() {
					uk.Stopped()
					return
				}
				tmsec++
			}
		}
		uk.Spine.States.PreSpike = 0
		uk.GraphRun(uk.ISISec, 0)
		tmsec = uk.Msec
	}
	uk.GraphRun(uk.FinalSecs, 0)
	tmsec = uk.Msec
	uk.Stopped()
}

func (uk *Urakubo) ThetaErrCompFun() {
	uk.ResetTimePlots()
	for itr := 0; itr < 2; itr++ {
		phsdur := []int{uk.DurMsec / 2, uk.DurMsec / 2}
		nphs := len(phsdur)

		// using send, recv for minus, plus
		sphz := []int{int(uk.SendHz), int(uk.RecvHz)}
		rphz := []int{int(uk.SendHz), int(uk.RecvHz)}

		if itr == 1 {
			sphz[1] = int(uk.SendHz)
			rphz[1] = int(uk.SendHz)
		}

		tmsec := 0
		uk.Init()
		// ss.RunQuiet(10)

		for ri := 0; ri < uk.NReps; ri++ {
			Sp := float32(1)
			Rp := float32(1)
			for pi := 0; pi < nphs; pi++ {
				dur := phsdur[pi]
				shz := sphz[pi]
				rhz := rphz[pi]
				Sint := math32.Exp(-1000.0 / float32(shz))
				Rint := math32.Exp(-1000.0 / float32(rhz))
				for msec := 0; msec < dur; msec++ {
					Sp *= rand.Float32()
					if Sp <= Sint {
						uk.Spine.States.PreSpike = 1
						Sp = 1
					} else {
						uk.Spine.States.PreSpike = 0
					}

					ge := float32(0.0)
					Rp *= rand.Float32()
					if Rp <= Rint {
						ge = uk.GeStim
						Rp = 1
					}
					if uk.RGClamp {
						ge = RGeStimForHz(float32(rhz))
					}

					uk.NeuronUpdt(tmsec, ge, 0)
					uk.StatsDefault(itr)
					if uk.StopNow() {
						uk.Stopped()
						return
					}
					tmsec++
				}
			}
			uk.Spine.States.PreSpike = 0
			uk.GraphRun(uk.ISISec, itr)
			tmsec = uk.Msec
		}
		uk.GraphRun(uk.FinalSecs, itr)
		tmsec = uk.Msec
		uk.StatsDWt(uk.Stats("DWtStats"), float64(itr), 0)
		uk.Plot("DWtPlot").GoUpdate()
	}
	uk.Stopped()
}

func (uk *Urakubo) ThetaErrSweepFun() {
	uk.ResetDWtPlot()

	hz := []int{25, 50, 100}
	nhz := len(hz)

	phsdur := []int{uk.DurMsec / 2, uk.DurMsec / 2}
	nphs := len(phsdur)

	sphz := []int{0, 0}
	rphz := []int{0, 0}

	for smi := 0; smi < nhz; smi++ {
		sphz[0] = hz[smi] // minus phase
		rphz[0] = hz[smi] // minus phase
		for spi := 0; spi < nhz; spi++ {
			sphz[1] = hz[spi] // plus phase
			rphz[1] = hz[spi] // plus phase

			tmsec := 0
			uk.ResetTimePlots()
			uk.Init()
			for ri := 0; ri < uk.NReps; ri++ {
				Sp := float32(1)
				Rp := float32(1)
				for pi := 0; pi < nphs; pi++ {
					dur := phsdur[pi]
					shz := sphz[pi]
					rhz := rphz[pi]
					Sint := math32.Exp(-1000.0 / float32(shz))
					Rint := math32.Exp(-1000.0 / float32(rhz))
					for msec := 0; msec < dur; msec++ {
						Sp *= rand.Float32()
						if Sp <= Sint {
							uk.Spine.States.PreSpike = 1
							Sp = 1
						} else {
							uk.Spine.States.PreSpike = 0
						}

						ge := float32(0.0)
						Rp *= rand.Float32()
						if Rp <= Rint {
							ge = uk.GeStim
							Rp = 1
						}
						if uk.RGClamp {
							ge = RGeStimForHz(float32(rhz))
						}

						uk.NeuronUpdt(tmsec, ge, 0)
						uk.StatsDefault(0)
						if uk.StopNow() {
							uk.Stopped()
							return
						}
						tmsec++
					}
				}
				uk.Spine.States.PreSpike = 0
				uk.GraphRun(uk.ISISec, 0)
				tmsec = uk.Msec
			}
			uk.GraphRun(uk.FinalSecs, 0)
			tmsec = uk.Msec
			uk.StatsPhaseDWt(uk.Stats("PhaseDWtStats"), sphz, rphz)
			uk.Plot("PhaseDWtPlot").GoUpdate()
		}
	}
	uk.Stopped()
}

func (uk *Urakubo) ThetaErrAllSweepFun() {
	uk.ResetDWtPlot()

	hz := []int{25, 50, 100}
	nhz := len(hz)

	phsdur := []int{uk.DurMsec / 2, uk.DurMsec / 2}
	nphs := len(phsdur)

	sphz := []int{0, 0}
	rphz := []int{0, 0}

	for smi := 0; smi < nhz; smi++ {
		sphz[0] = hz[smi] // minus phase
		for spi := 0; spi < nhz; spi++ {
			sphz[1] = hz[spi] // plus phase
			for rmi := 0; rmi < nhz; rmi++ {
				rphz[0] = hz[rmi] // minus phase
				for rpi := 0; rpi < nhz; rpi++ {
					rphz[1] = hz[rpi] // plus phase

					uk.ResetTimePlots()
					uk.Init()
					for ri := 0; ri < uk.NReps; ri++ {
						Sp := float32(1)
						Rp := float32(1)
						for pi := 0; pi < nphs; pi++ {
							dur := phsdur[pi]
							shz := sphz[pi]
							rhz := rphz[pi]
							Sint := math32.Exp(-1000.0 / float32(shz))
							Rint := math32.Exp(-1000.0 / float32(rhz))
							for msec := 0; msec < dur; msec++ {
								Sp *= rand.Float32()
								if Sp <= Sint {
									uk.Spine.States.PreSpike = 1
									Sp = 1
								} else {
									uk.Spine.States.PreSpike = 0
								}

								ge := float32(0.0)
								Rp *= rand.Float32()
								if Rp <= Rint {
									ge = uk.GeStim
									Rp = 1
								}
								if uk.RGClamp {
									ge = RGeStimForHz(float32(rhz))
								}

								uk.NeuronUpdt(msec, ge, 0)
								uk.StatsDefault(0)
								if uk.StopNow() {
									uk.Stopped()
									return
								}
							}
						}
						uk.Spine.States.PreSpike = 0
						uk.GraphRun(uk.ISISec, 0)
					}
					uk.GraphRun(uk.FinalSecs, 0)
					uk.StatsPhaseDWt(uk.Stats("PhaseDWtStats"), sphz, rphz)
					uk.Plot("PhaseDWtPlot").GoUpdate()
				}
			}
		}
	}
	uk.Stopped()
}
