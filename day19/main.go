package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

const (
	example = `--- scanner 0 ---
404,-588,-901
528,-643,409
-838,591,734
390,-675,-793
-537,-823,-458
-485,-357,347
-345,-311,381
-661,-816,-575
-876,649,763
-618,-824,-621
553,345,-567
474,580,667
-447,-329,318
-584,868,-557
544,-627,-890
564,392,-477
455,729,728
-892,524,684
-689,845,-530
423,-701,434
7,-33,-71
630,319,-379
443,580,662
-789,900,-551
459,-707,401

--- scanner 1 ---
686,422,578
605,423,415
515,917,-361
-336,658,858
95,138,22
-476,619,847
-340,-569,-846
567,-361,727
-460,603,-452
669,-402,600
729,430,532
-500,-761,534
-322,571,750
-466,-666,-811
-429,-592,574
-355,545,-477
703,-491,-529
-328,-685,520
413,935,-424
-391,539,-444
586,-435,557
-364,-763,-893
807,-499,-711
755,-354,-619
553,889,-390

--- scanner 2 ---
649,640,665
682,-795,504
-784,533,-524
-644,584,-595
-588,-843,648
-30,6,44
-674,560,763
500,723,-460
609,671,-379
-555,-800,653
-675,-892,-343
697,-426,-610
578,704,681
493,664,-388
-671,-858,530
-667,343,800
571,-461,-707
-138,-166,112
-889,563,-600
646,-828,498
640,759,510
-630,509,768
-681,-892,-333
673,-379,-804
-742,-814,-386
577,-820,562

--- scanner 3 ---
-589,542,597
605,-692,669
-500,565,-823
-660,373,557
-458,-679,-417
-488,449,543
-626,468,-788
338,-750,-386
528,-832,-391
562,-778,733
-938,-730,414
543,643,-506
-524,371,-870
407,773,750
-104,29,83
378,-903,-323
-778,-728,485
426,699,580
-438,-605,-362
-469,-447,-387
509,732,623
647,635,-688
-868,-804,481
614,-800,639
595,780,-596

--- scanner 4 ---
727,592,562
-293,-554,779
441,611,-461
-714,465,-776
-743,427,-804
-660,-479,-426
832,-632,460
927,-485,-438
408,393,-506
466,436,-512
110,16,151
-258,-428,682
-393,719,612
-211,-452,876
808,-476,-593
-575,615,604
-485,667,467
-680,325,-822
-627,-443,-432
872,-547,-609
833,512,582
807,604,487
839,-516,451
891,-625,532
-652,-548,-490
30,-46,-14`
)

func readFile(path string) (string, error) {
	input, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer input.Close()

	var inputLines []string
	buffer := bufio.NewScanner(input)
	for buffer.Scan() {
		inputLines = append(inputLines, buffer.Text())
	}

	return strings.Join(inputLines, "\n"), buffer.Err()
}

func computeScanners(rawScanners []string) []Scanner {
	var scanners []Scanner
	for i, rawScanner := range rawScanners {
		var beacons []Point
		for _, line := range strings.Split(rawScanner, "\n") {
			if line[0:3] == "---" {
				continue
			}
			var x, y, z int
			fmt.Sscanf(strings.TrimSpace(line), "%d,%d,%d", &x, &y, &z)
			beacons = append(beacons, Point{x, y, z})
		}
		scanners = append(scanners, Scanner{i, Point{0, 0, 0}, beacons})
	}
	return scanners
}

type CoOrd struct {
	x int
	y int
}

type Point struct {
	x int
	y int
	z int
}

type PointSlice []Point

// Len is part of sort.Interface.
func (d PointSlice) Len() int {
	return len(d)
}

// Swap is part of sort.Interface.
func (d PointSlice) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

// Less is part of sort.Interface. We use count as the value to sort by
func (d PointSlice) Less(i, j int) bool {
	return d[i].x < d[j].x
}

func (p *Point) Vector(o Point) Point {
	return Point{o.x - p.x, o.y - p.y, o.z - p.z}
}

// func (p *Point) String() string {
// 	return fmt.Sprintf("%d,%d,%d", p.x, p.y, p.z)
// }

func (p *Point) Translate(t Point) Point {
	return Point{p.x + t.x, p.y + t.y, p.z + t.z}
}

func (p *Point) Distance(d Point) int {
	return int(math.Abs(float64(d.x-p.x)) + math.Abs(float64(d.y-p.y)) + math.Abs(float64(d.z-p.z)))
}

func (p *Point) Rotate(rotation int) Point {
	var x, y, z = p.x, p.y, p.z
	switch rotation {
	case 0:
		return Point{x, y, z}
	case 1:
		return Point{y, -x, z}
	case 2:
		return Point{-x, -y, z}
	case 3:
		return Point{-y, x, z}
	case 4:
		return Point{z, y, -x}
	case 5:
		return Point{y, -z, -x}
	case 6:
		return Point{-z, -y, -x}
	case 7:
		return Point{-y, z, -x}
	case 8:
		return Point{z, -x, -y}
	case 9:
		return Point{-x, -z, -y}
	case 10:
		return Point{-z, x, -y}
	case 11:
		return Point{x, z, -y}
	case 12:
		return Point{z, -y, x}
	case 13:
		return Point{-y, -z, x}
	case 14:
		return Point{-z, y, x}
	case 15:
		return Point{y, z, x}
	case 16:
		return Point{z, x, y}
	case 17:
		return Point{x, -z, y}
	case 18:
		return Point{-z, -x, y}
	case 19:
		return Point{-x, z, y}
	case 20:
		return Point{-x, y, -z}
	case 21:
		return Point{y, x, -z}
	case 22:
		return Point{x, -y, -z}
	case 23:
		return Point{-y, -x, -z}
	default:
		panic("Rut Roh! Can't rotate to that!")
	}
}

type Scanner struct {
	id          int
	orientation Point
	beacons     []Point
}

func buildVectors(beacons map[Point]Point) map[Point]Point {
	vectors := make(map[Point]Point)
	for _, beacon1 := range beacons {
		for _, beacon2 := range beacons {
			if beacon1 == beacon2 {
				continue
			}
			newVector := beacon2.Vector(beacon1)
			if _, exists := vectors[newVector]; !exists {
				vectors[newVector] = beacon2
			}
		}
	}
	return vectors
}

func testRotation(knownVectors map[Point]Point, beacons []Point, rotation int) (bool, Point) {
	matchCounter := 0

	for _, beacon1 := range beacons {
		beacon1Rotated := beacon1.Rotate(rotation)

		for _, beacon2 := range beacons {
			if beacon1 == beacon2 {
				continue
			}

			beacon2Rotated := beacon2.Rotate(rotation)

			vector := beacon1Rotated.Vector(beacon2Rotated)

			if _, exists := knownVectors[vector]; exists {
				matchCounter++
			}
			if matchCounter == 11 {
				return true, beacon1Rotated.Vector(knownVectors[vector])
			}
		}
	}

	return false, Point{}
}

func rotatePoints(beacons []Point, rotation int) []Point {
	var results []Point
	for _, beacon := range beacons {
		results = append(results, beacon.Rotate(rotation))
	}
	return results
}

func translatePoints(beacons []Point, translate Point) []Point {
	var results []Point
	for _, beacon := range beacons {
		results = append(results, beacon.Translate(translate))
	}
	return results
}

func printBeacons(knownBeacons map[Point]Point) {
	// Can't sort a map so do some magical StackOverflow turning into an array
	knownBeaconsSorted := make(PointSlice, 0, len(knownBeacons))
	for _, knownBeacon := range knownBeacons {
		knownBeaconsSorted = append(knownBeaconsSorted, knownBeacon)
	}

	sort.Sort(knownBeaconsSorted)

	fmt.Printf("Known Beacons:\n")
	for _, knownBeacon := range knownBeaconsSorted {
		fmt.Printf("%d,%d,%d\n", knownBeacon.x, knownBeacon.y, knownBeacon.z)
	}
}

func compute(scanners []Scanner) (map[Point]Point, []Scanner) {
	// Assume all Beacons in Scanner 0 are correct and it's the starting point
	knownScanners := []Scanner{scanners[0]}
	knownBeacons := make(map[Point]Point)

	for _, kb := range knownScanners[0].beacons {
		knownBeacons[kb] = kb
	}

	printBeacons(knownBeacons)

	unknownScanners := scanners[1:]

	for {
		unknownScanner := unknownScanners[0]
		unknownScanners = unknownScanners[1:]

		knownVectors := buildVectors(knownBeacons)

		workingRotation := -1
		var workingVector Point

		for rotation := 0; rotation < 24; rotation++ {
			passed, vector := testRotation(knownVectors, unknownScanner.beacons, rotation)
			if passed {
				workingRotation = rotation
				workingVector = vector
				break
			}
		}

		if workingRotation != -1 {
			rotatedScannerBeacons := rotatePoints(unknownScanner.beacons, workingRotation)
			translatedScannerBeacons := translatePoints(rotatedScannerBeacons, workingVector)
			for _, beacon := range translatedScannerBeacons {
				knownBeacons[beacon] = beacon
			}
			knownScanners = append(knownScanners, Scanner{id: unknownScanner.id, orientation: workingVector, beacons: translatedScannerBeacons})
		} else {
			unknownScanners = append(unknownScanners, unknownScanner)
		}

		if len(unknownScanners) == 0 {
			break
		}
	}

	// printBeacons(knownBeacons)

	return knownBeacons, knownScanners
}

func part1(beacons map[Point]Point) int {
	return len(beacons)
}

func part2(scanners []Scanner) int {
	maxDistance := 0
	for _, a := range scanners {
		for _, b := range scanners {
			if a.id == b.id {
				continue
			}
			distance := a.orientation.Distance(b.orientation)
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return maxDistance
}

func main() {
	input, _ := readFile("input")

	rawScanners := strings.Split(input, "\n\n")
	scanners := computeScanners(rawScanners)

	resultBeacons, resultScanners := compute(scanners)

	resultPart1 := part1(resultBeacons)
	resultPart2 := part2(resultScanners)

	fmt.Printf("Part 1 = %d\n", resultPart1)

	fmt.Printf("Part 2 = %+v\n", resultPart2)
}
