package main

import "testing"

func TestGetPacketVersion(t *testing.T) {
	result := getPacketVersion(hexToBinary(example1))

	if result != 1 {
		t.Fatalf("Expected 1 but got %d", result)
	}
}

func TestGetPacketTypeID(t *testing.T) {
	result := getPacketTypeID(hexToBinary(example1))

	if result != 6 {
		t.Fatalf("Expected 6 but got %d", result)
	}
}

func TestGetPacketLengthID(t *testing.T) {
	result := getPacketLengthID(hexToBinary(example1))

	if result != 0 {
		t.Fatalf("Expected 0 but got %d", result)
	}
}

func TestGetPacketLengthID2(t *testing.T) {
	result := getPacketLengthID(hexToBinary(example2))

	if result != 1 {
		t.Fatalf("Expected 1 but got %d", result)
	}
}

func TestParseLiteralPacket(t *testing.T) {
	result, _ := parsePacket(hexToBinary(example0))

	if result.value != 2021 {
		t.Fatalf("Expected 2021 but got %d", result)
	}
}

func TestPart1Example3(t *testing.T) {
	result := part1(example3)

	if result != 16 {
		t.Fatalf("Expected 16 but got %d", result)
	}
}

func TestPart1Example4(t *testing.T) {
	result := part1(example4)

	if result != 12 {
		t.Fatalf("Expected 12 but got %d", result)
	}
}

func TestPart1Example5(t *testing.T) {
	result := part1(example5)

	if result != 23 {
		t.Fatalf("Expected 23 but got %d", result)
	}
}

func TestPart1Example6(t *testing.T) {
	result := part1(example6)

	if result != 31 {
		t.Fatalf("Expected 31 but got %d", result)
	}
}

func TestPart2Example7(t *testing.T) {
	result := part2(example7)

	if result != 3 {
		t.Fatalf("expected 3 but got %d", result)
	}
}

func TestPart2Example8(t *testing.T) {
	result := part2(example8)

	if result != 54 {
		t.Fatalf("expected 54 but got %d", result)
	}
}

func TestPart2Example9(t *testing.T) {
	result := part2(example9)

	if result != 7 {
		t.Fatalf("expected 54 but got %d", result)
	}
}

func TestPart2Example10(t *testing.T) {
	result := part2(example10)

	if result != 9 {
		t.Fatalf("expected 9 but got %d", result)
	}
}

func TestPart2Example11(t *testing.T) {
	result := part2(example10)

	if result != 9 {
		t.Fatalf("expected 9 but got %d", result)
	}
}

func TestPart2Example12(t *testing.T) {
	result := part2(example12)

	if result != 0 {
		t.Fatalf("expected 0 but got %d", result)
	}
}

func TestPart2Example13(t *testing.T) {
	result := part2(example13)

	if result != 0 {
		t.Fatalf("expected 0 but got %d", result)
	}
}

func TestPart2Example14(t *testing.T) {
	result := part2(example14)

	if result != 1 {
		t.Fatalf("expected 0 but got %d", result)
	}
}
