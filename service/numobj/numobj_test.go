package numobj

import (
	"bignum-service/domain/bignum"
	"bignum-service/domain/bignum/mocks"
	"bignum-service/lib/ctxlib"
	"fmt"
	"math/big"
	"testing"
)

func getBigNum(name, validBigFloatStr string) *bignum.BigNum {
	num, _ := bignum.NewFromString(name, validBigFloatStr)
	return num
}

func getBigFloat(validBigFloatStr string) *big.Float {
	floatNum, _ := bignum.ParseFloat(validBigFloatStr)
	return floatNum
}

func TestMultiply(t *testing.T) {
	mockRepo := &mocks.Repository{}
	svc, err := NewService(WithBignumMemRepo(mockRepo))
	ctx := ctxlib.Background()

	if err != nil {
		t.Errorf("new numobj service: %s", err)
	}
	cases := []struct {
		name           string
		num1Str        string
		num2Str        string
		mockFunc       func()
		expectedErr    error
		expectedResult *big.Float
	}{
		{
			name:           "both are big float numbers",
			num1Str:        "1.2",
			num2Str:        "2.3",
			expectedErr:    nil,
			expectedResult: new(big.Float).Mul(getBigFloat("1.2"), getBigFloat("2.3")),
		},
		{
			name:    "one number object and one float number, number object is available",
			num1Str: "planet_mass",
			num2Str: "2.3",
			mockFunc: func() {
				mockRepo.On("GetNum", ctx, "planet_mass").Return(getBigNum("planet_mass", "6416930923733925522307001.29472615"), nil).Once()
			},
			expectedErr:    nil,
			expectedResult: getBigFloat("1.4758941124588028701e+25"),
		},
		{
			name:    "one number object and one float number, number object not found",
			num1Str: "planet_mass",
			num2Str: "2.3",
			mockFunc: func() {
				mockRepo.On("GetNum", ctx, "planet_mass").Return(nil, ErrNumberObjectNotFound).Once()
			},
			expectedErr:    ErrNumberObjectNotFound,
			expectedResult: nil,
		},
		{
			name:    "two number object, number objects are available",
			num1Str: "planet_mass",
			num2Str: "grav_const",
			mockFunc: func() {
				mockRepo.On("GetNum", ctx, "planet_mass").Return(getBigNum("planet_mass", "6416930923733925522307001.29472615"), nil).Once()
				mockRepo.On("GetNum", ctx, "grav_const").Return(getBigNum("grav_const", "0.000000000066731039356729"), nil).Once()
			},
			expectedErr:    nil,
			expectedResult: getBigFloat("4.2820847002109996115e+14"),
		},
	}
	for idx, tc := range cases {
		if tc.mockFunc != nil {
			tc.mockFunc()
		}
		t.Run(fmt.Sprintf("TC%02d %s", idx, tc.name), func(t *testing.T) {
			result, err := svc.Multiply(ctx, tc.num1Str, tc.num2Str)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if (result == nil && tc.expectedResult != nil) ||
				(result != nil && tc.expectedResult == nil) {
				t.Errorf("Expected result %v, got %v", tc.expectedResult, result)
			}
			ctx.Logger.Debug().Interface("result", result).Interface("expected", tc.expectedResult).Msg("")
			if result != nil && tc.expectedResult != nil && result.Cmp(tc.expectedResult) != 0 {
				t.Errorf("Expected result %v, got %v", tc.expectedResult, result)
			}
		})
	}
}
