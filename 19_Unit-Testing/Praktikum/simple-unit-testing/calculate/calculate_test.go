package calculate

import "testing"

func TestAddition(t *testing.T) {

	testCases := []struct {
		name     string
		input    struct{ a, b float64 }
		expected float64
	}{
		{
			name:     "Test Case 1 : Addition of Positive Numbers",
			input:    struct{ a, b float64 }{7, 7},
			expected: 14,
		},
		{
			name:     "Test Case 2 : Addition of Negative Numbers",
			input:    struct{ a, b float64 }{-7, -7},
			expected: -14,
		},
		{
			name:     "Test Case 3 : Addition of Positive and Negative Numbers",
			input:    struct{ a, b float64 }{10, -3},
			expected: 7,
		},
		{
			name:     "Test Case 4 : Addition of Fractional Numbers",
			input:    struct{ a, b float64 }{2.5, 1.5},
			expected: 4.0,
		},
	}

	calc := Calculate{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calc.Addition(tc.input.a, tc.input.b)
			if result != tc.expected {
				t.Errorf("%s: Addition(%f, %f) = %f; Expected %f", tc.name, tc.input.a, tc.input.b, result, tc.expected)
			}
		})
	}
}

func TestSubtraction(t *testing.T) {

	testCases := []struct {
		name     string
		input    struct{ a, b float64 }
		expected float64
	}{
		{
			name:     "Test Case 1 : Subtraction of Positive Numbers",
			input:    struct{ a, b float64 }{10, 3},
			expected: 7,
		},
		{
			name:     "Test Case 2 : Subtraction of Negative Numbers",
			input:    struct{ a, b float64 }{-7, -3},
			expected: -4,
		},
		{
			name:     "Test Case 3 : Subtraction of Positive and Negative Numbers",
			input:    struct{ a, b float64 }{10, -4},
			expected: 14,
		},
		{
			name:     "Test Case 4 : Subtraction of Fractional Numbers",
			input:    struct{ a, b float64 }{3.5, 1.5},
			expected: 2.0,
		},
	}

	calc := Calculate{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calc.Subtraction(tc.input.a, tc.input.b)
			if result != tc.expected {
				t.Errorf("%s: Subtraction(%f, %f) = %f; Expected %f", tc.name, tc.input.a, tc.input.b, result, tc.expected)
			}
		})
	}
}

func TestDivision(t *testing.T) {

	testCases := []struct {
		name      string
		input     struct{ a, b float64 }
		expected  float64
		expectErr bool
	}{
		{
			name:      "Test Case 1 : Division of Positive Numbers",
			input:     struct{ a, b float64 }{10, 2},
			expected:  5,
			expectErr: false,
		},
		{
			name:      "Test Case 2 : Division by Zero",
			input:     struct{ a, b float64 }{7, 0},
			expected:  0,
			expectErr: true,
		},
		{
			name:      "Test Case 3 : Division of Positive and Negative Numbers",
			input:     struct{ a, b float64 }{10, -2},
			expected:  -5,
			expectErr: false,
		},
		{
			name:      "Test Case 4 : Division of Fractional Numbers",
			input:     struct{ a, b float64 }{3.0, 1.5},
			expected:  2.0,
			expectErr: false,
		},
	}

	calc := Calculate{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result, err := calc.Division(tc.input.a, tc.input.b)
			if tc.expectErr {
				if err == nil {
					t.Errorf("%s: Division(%f, %f) Didn't Return an Error as Expected", tc.name, tc.input.a, tc.input.b)
				}
				return
			}

			if err != nil {
				t.Errorf("%s: Division(%f, %f) Returned an Unexpected Error: %v", tc.name, tc.input.a, tc.input.b, err)
				return
			}

			if result != tc.expected {
				t.Errorf("%s: Division(%f, %f) = %f; Expected %f", tc.name, tc.input.a, tc.input.b, result, tc.expected)
			}
		})
	}
}

func TestMultiplication(t *testing.T) {

	testCases := []struct {
		name     string
		input    struct{ a, b float64 }
		expected float64
	}{
		{
			name:     "Test Case 1 : Multiplication of Positive Numbers",
			input:    struct{ a, b float64 }{7, 5},
			expected: 35,
		},
		{
			name:     "Test Case 2 : Multiplication of Negative Numbers",
			input:    struct{ a, b float64 }{-7, -3},
			expected: 21,
		},
		{
			name:     "Test Case 3 : Multiplication of Positive and Negative Numbers",
			input:    struct{ a, b float64 }{10, -7},
			expected: -70,
		},
		{
			name:     "Test Case 4 : Multiplication of Fractional Numbers",
			input:    struct{ a, b float64 }{3.5, 2},
			expected: 7.0,
		},
	}

	calc := Calculate{}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := calc.Multiplication(tc.input.a, tc.input.b)
			if result != tc.expected {
				t.Errorf("%s: Multiplication(%f, %f) = %f; Expected %f", tc.name, tc.input.a, tc.input.b, result, tc.expected)
			}
		})
	}
}
