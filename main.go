package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"image/color"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"text/tabwriter"
	"time"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"gonum.org/v1/plot/vg/draw"
)

type RatingTicker struct{}

func (ticker RatingTicker) Ticks(min, max float64) []plot.Tick {
	low := int(min + 1) // 4 -> 4, 5, 6, 7, 8
	high := int(max)    // 8
	ticks := make([]plot.Tick, high-low+1)
	for low <= high {
		ticks = append(ticks,
			plot.Tick{Value: float64(low),
				Label: strconv.Itoa(low),
			})
		low++
	}
	return ticks
}

var reverse bool

func main() {
	var chart, poor, sortTime bool
	var targetScore float64
	var fileLoc string
	flag.Float64Var(&targetScore, "target", 100.0, "Print only scores lower than this")
	flag.StringVar(&fileLoc, "profile", "", "Profile location")
	flag.BoolVar(&reverse, "reverse", false, "Reverse order")
	flag.BoolVar(&chart, "chart", false, "Generate a chart")
	flag.BoolVar(&poor, "poor", false, "List bad scores")
	flag.BoolVar(&sortTime, "time", false, "Sort by time")
	flag.Parse()

	bytes, err := ioutil.ReadFile(fileLoc)
	if err != nil {
		log.Fatal(err)
	}

	var stats Stats
	if err := xml.Unmarshal(bytes, &stats); err != nil {
		log.Fatal(err)
	}

	var validCount int
	var invalidCount int
	var totalWife float32

	var scores Scores
	var bestScores BestScores
	if chart {
		scores = make(Scores, 0, 2000)
	}
	if poor {
		bestScores = make(BestScores, 0, 2000)
	}

	streams := make(FloatArray, 0, 2000)
	jumpstreams := make(FloatArray, 0, 2000)
	handstreams := make(FloatArray, 0, 2000)
	stamina := make(FloatArray, 0, 2000)
	jacks := make(FloatArray, 0, 2000)
	chordjacks := make(FloatArray, 0, 2000)
	technical := make(FloatArray, 0, 2000)

	for _, song := range stats.PlayerScores.Chart {
		for _, scoresAt := range song.ScoresAt {
			for _, score := range scoresAt.Scores {
				if chart {
					scores = append(scores, score)
					continue
				}

				if score.Key != scoresAt.PBKey {
					continue
				}

				if poor {
					t, err := time.Parse("2006-01-02 15:04:05", score.DateTime)
					if err != nil {
						log.Println("Unable to parse time:", err.Error())
					}
					bestScores = append(bestScores, BestScore{
						WifeScore:      score.WifeScore,
						SSRNormPercent: score.SSRNormPercent,
						EtternaValid:   score.EtternaValid,
						Pack:           song.Pack,
						Song:           song.Song,
						Rate:           scoresAt.Rate,
						Steps:          song.Steps,
						Date:           t,
						Grade:          score.Grade,
						Overall:        score.SkillsetSSRs.Overall,
						Stream:         score.SkillsetSSRs.Stream,
						Jumpstream:     score.SkillsetSSRs.Jumpstream,
						Handstream:     score.SkillsetSSRs.Handstream,
						Stamina:        score.SkillsetSSRs.Stamina,
						JackSpeed:      score.SkillsetSSRs.JackSpeed,
						Chordjack:      score.SkillsetSSRs.Chordjack,
						Technical:      score.SkillsetSSRs.Technical,
					})
					continue
				}

				// General stats
				if score.EtternaValid == 0 {
					invalidCount++
				}
				validCount++
				totalWife += score.SSRNormPercent
				streams = append(streams, score.SkillsetSSRs.Stream)
				jumpstreams = append(jumpstreams, score.SkillsetSSRs.Jumpstream)
				handstreams = append(handstreams, score.SkillsetSSRs.Handstream)
				stamina = append(stamina, score.SkillsetSSRs.Stamina)
				jacks = append(jacks, score.SkillsetSSRs.JackSpeed)
				chordjacks = append(chordjacks, score.SkillsetSSRs.Chordjack)
				technical = append(technical, score.SkillsetSSRs.Technical)
			}
		}
	}

	if poor {
		if sortTime {
			timeScores := BestScoresByTime(bestScores)
			sort.Sort(timeScores)
			bestScores = BestScores(timeScores)
		} else {
			sort.Sort(bestScores)
		}
		w := new(tabwriter.Writer)
		// minwidth, tabwidth, padding, padchar, flags
		w.Init(os.Stdout, 8, 8, 0, '\t', 0)
		defer w.Flush()

		printedCount := 0
		for _, score := range bestScores {
			if score.SSRNormPercent*100 > float32(targetScore) {
				continue
			}
			var difficulty string
			switch score.Steps {
			case "Beginner":
				difficulty = "\033[0;36mBeginner\033[0m"
			case "Easy":
				difficulty = "\033[0;32mEasy    \033[0m"
			case "Medium":
				difficulty = "\033[0;33mMedium  \033[0m"
			case "Hard":
				difficulty = "\033[0;31mHard    \033[0m"
			case "Challenge":
				difficulty = "\033[0;35mInsane  \033[0m"
			default:
				difficulty = "\033[0;35m" + score.Steps + "  \033[0m"
			}

			valid := "\033[0;32m✓\033[0m"
			if 0 == score.EtternaValid {
				valid = "\033[0;31m✘\033[0m"
			}

			fmt.Fprintf(w, "\n %s %.2f%%\t%.2f\t%s\t%s", valid, score.SSRNormPercent*100, score.Rate, score.HardestSkill(), score.Song)
			fmt.Fprintf(w, "\n %s\t%s\t\t%s", difficulty, score.Date.Format("2006, Jan 2"), score.Pack)
			fmt.Fprintf(w, "\n \033[0;37m─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────\033[0m")
			printedCount++
		}
		fmt.Fprintln(w, "\n Total:", printedCount)
		return
	}

	if chart {
		p, err := plot.New()
		if err != nil {
			panic(err)
		}

		p.Title.Text = "Etterna Overall Scores"
		p.X.Label.Text = "Song"
		p.Y.Label.Text = "Rating"

		p.Y.Tick.Marker = RatingTicker{}

		p.Add(plotter.NewGrid())

		sort.Sort(scores)

		s, err := NewScatter(scores)
		if err != nil {
			panic(err)
		}

		p.Add(s)

		if err := p.Save(64*vg.Inch, 36*vg.Inch, "points.png"); err != nil {
			panic(err)
		}

		return
	}

	// Default behaviour
	sort.Sort(streams)
	sort.Sort(jumpstreams)
	sort.Sort(handstreams)
	sort.Sort(stamina)
	sort.Sort(jacks)
	sort.Sort(chordjacks)
	sort.Sort(technical)

	ss := streams.Score()
	jss := jumpstreams.Score()
	hss := handstreams.Score()
	sts := stamina.Score()
	js := jacks.Score()
	cjs := chordjacks.Score()
	ts := technical.Score()
	os := ((ss + jss + hss + sts + js + cjs + ts) / 7.0)

	fmt.Println("Valid Scores:", validCount)
	fmt.Println("Invalid Scores:", invalidCount)
	fmt.Printf("Average Score: %.3f\n", 100.0*totalWife/float32(validCount))
	fmt.Printf("Overall: %.2f\n", os)
	fmt.Printf("Stream: %.2f\n", ss)
	fmt.Printf("Jumpstream: %.2f\n", jss)
	fmt.Printf("Handstream: %.2f\n", hss)
	fmt.Printf("Stamina: %.2f\n", sts)
	fmt.Printf("Jackspeed: %.2f\n", js)
	fmt.Printf("Chordjack: %.2f\n", cjs)
	fmt.Printf("Technical: %.2f\n", ts)
}

func (score BestScore) HardestSkill() string {
	name, rating := "Chordjack ", score.Chordjack
	name, rating = ifHigher(name, rating, "Stream    ", score.Stream)
	name, rating = ifHigher(name, rating, "Jumpstream", score.Jumpstream)
	name, rating = ifHigher(name, rating, "Handstream", score.Handstream)
	name, rating = ifHigher(name, rating, "Stamina   ", score.Stamina)
	name, rating = ifHigher(name, rating, "Jackspeed ", score.JackSpeed)
	name, rating = ifHigher(name, rating, "Technical ", score.Technical)
	return name
}

func ifHigher(oldname string, high float32, name string, rating float32) (string, float32) {
	if rating > high {
		return name, rating
	}
	return oldname, high
}

func (a FloatArray) Score() float32 {
	total := float32(0.0)
	for i := 5; i < 10; i++ {
		total += a[i]
	}

	return total / 5.0
}

func (scores Scores) XY(i int) (x, y float64) {
	return float64(i), float64(scores[i].SkillsetSSRs.Overall)
}

type FloatArray []float32

func (a Scores) Len() int      { return len(a) }
func (a Scores) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a Scores) Less(i, j int) bool {
	if !reverse {
		return a[i].SkillsetSSRs.Overall < a[j].SkillsetSSRs.Overall
	}
	return a[i].SkillsetSSRs.Overall > a[j].SkillsetSSRs.Overall
}

func (a BestScores) Len() int      { return len(a) }
func (a BestScores) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BestScores) Less(i, j int) bool {
	if !reverse {
		return a[i].SSRNormPercent > a[j].SSRNormPercent
	}
	return a[i].SSRNormPercent < a[j].SSRNormPercent
}

func (a BestScoresByTime) Len() int      { return len(a) }
func (a BestScoresByTime) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a BestScoresByTime) Less(i, j int) bool {
	if !reverse {
		return a[i].Date.After(a[j].Date)
	}
	return a[i].Date.Before(a[j].Date)
}

func (a FloatArray) Len() int           { return len(a) }
func (a FloatArray) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a FloatArray) Less(i, j int) bool { return a[i] > a[j] }

func NewScatter(scores Scores) (*plotter.Scatter, error) {
	data, err := plotter.CopyXYs(scores)
	if err != nil {
		return nil, err
	}
	return &plotter.Scatter{
		XYs: data,
		GlyphStyleFunc: func(i int) draw.GlyphStyle {
			score := scores[i]
			wife := float64(((score.SSRNormPercent * 10) - 8) * 5)

			var parts float64
			if wife > 5 {
				parts = (1 - ((wife - 5) / 5))
			} else {
				parts = wife / 5
			}
			comp := uint8(parts * 255)

			var c color.RGBA
			if score.EtternaValid != 1 {
				c = color.RGBA{R: 0, G: 0, B: 0, A: 255}
			} else if wife < 5 {
				c = color.RGBA{R: 255, G: comp, B: 0, A: 255}
			} else {
				c = color.RGBA{R: comp, G: 255 - ((255 - comp) / 2), B: 255 - comp, A: 255}
			}

			return draw.GlyphStyle{
				Color:  c,
				Radius: vg.Points(6.0),
				Shape:  draw.PyramidGlyph{},
			}
		},
	}, err
}
