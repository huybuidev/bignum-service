package bignummem

import (
	"bignum-service/domain/bignum"
	"bignum-service/lib/ctxlib"
	"fmt"
	"math/big"
	"reflect"
	"testing"
)

func TestGetNum(t *testing.T) {
	testCases := []struct {
		name          string
		numObjectName string
		expectedErr   error
		expectedValue *big.Float
		numObjects    map[string]*big.Float
	}{
		{
			name:          "number object found",
			numObjectName: "abc",
			expectedErr:   nil,
			expectedValue: big.NewFloat(1.23),
			numObjects:    map[string]*big.Float{"abc": big.NewFloat(1.23)},
		},
		{
			name:          "number object not found",
			numObjectName: "abc",
			expectedErr:   ErrNotFound,
			expectedValue: nil,
			numObjects:    map[string]*big.Float{"abcd": big.NewFloat(1.23)},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("TC%02d-%s", idx, tc.name), func(t *testing.T) {
			ctx := ctxlib.Background()
			repo := &BignumMemoryRepo{
				numObjects: tc.numObjects,
			}
			num, err := repo.GetNum(ctx, tc.numObjectName)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil {
				if num.Value().String() != tc.expectedValue.String() {
					t.Errorf("Expected value %v, got %v", tc.expectedValue, num.Value())
				}
			}
		})
	}
}

func TestUpdateNum(t *testing.T) {
	testCases := []struct {
		name               string
		numObjectName      string
		numObjectValue     *big.Float
		expectedErr        error
		numObjects         map[string]*big.Float
		expectedNumObjects map[string]*big.Float
	}{
		{
			name:               "number object found",
			numObjectName:      "abc",
			numObjectValue:     big.NewFloat(4.56),
			expectedErr:        nil,
			numObjects:         map[string]*big.Float{"abc": big.NewFloat(1.23)},
			expectedNumObjects: map[string]*big.Float{"abc": big.NewFloat(4.56)},
		},
		{
			name:               "number object not found",
			numObjectName:      "abc",
			numObjectValue:     big.NewFloat(4.56),
			expectedErr:        ErrNotFound,
			numObjects:         map[string]*big.Float{"a": big.NewFloat(1.23)},
			expectedNumObjects: map[string]*big.Float{"a": big.NewFloat(1.23)},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("TC%02d-%s", idx, tc.name), func(t *testing.T) {
			ctx := ctxlib.Background()
			repo := &BignumMemoryRepo{
				numObjects: tc.numObjects,
			}
			num, _ := bignum.New(tc.numObjectName, tc.numObjectValue)
			err := repo.UpdateNum(ctx, num)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil {
				if !reflect.DeepEqual(tc.expectedNumObjects, repo.numObjects) {
					t.Errorf("Expected map %v, got %v", tc.expectedNumObjects, repo.numObjects)
				}
			}
		})
	}
}

func TestPutNum(t *testing.T) {
	testCases := []struct {
		name               string
		numObjectName      string
		numObjectValue     *big.Float
		expectedErr        error
		numObjects         map[string]*big.Float
		expectedNumObjects map[string]*big.Float
	}{
		{
			name:               "number object created from empty map",
			numObjectName:      "abc",
			numObjectValue:     big.NewFloat(4.56),
			expectedErr:        nil,
			numObjects:         make(map[string]*big.Float),
			expectedNumObjects: map[string]*big.Float{"abc": big.NewFloat(4.56)},
		},
		{
			name:               "number object created from non empty map",
			numObjectName:      "abc",
			numObjectValue:     big.NewFloat(4.56),
			expectedErr:        nil,
			numObjects:         map[string]*big.Float{"a": big.NewFloat(1.23)},
			expectedNumObjects: map[string]*big.Float{"a": big.NewFloat(1.23), "abc": big.NewFloat(4.56)},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("TC%02d-%s", idx, tc.name), func(t *testing.T) {
			ctx := ctxlib.Background()
			repo := &BignumMemoryRepo{
				numObjects: tc.numObjects,
			}
			num, _ := bignum.New(tc.numObjectName, tc.numObjectValue)
			err := repo.PutNum(ctx, num)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil {
				if !reflect.DeepEqual(tc.expectedNumObjects, repo.numObjects) {
					t.Errorf("Expected map %v, got %v", tc.expectedNumObjects, repo.numObjects)
				}
			}
		})
	}
}

func TestDeleteNum(t *testing.T) {
	testCases := []struct {
		name               string
		numObjectName      string
		expectedErr        error
		numObjects         map[string]*big.Float
		expectedNumObjects map[string]*big.Float
	}{
		{
			name:               "successfully deleted",
			numObjectName:      "abc",
			expectedErr:        nil,
			numObjects:         map[string]*big.Float{"abc": big.NewFloat(4.56)},
			expectedNumObjects: make(map[string]*big.Float),
		},
		{
			name:               "number object not in map",
			numObjectName:      "abc",
			expectedErr:        ErrNotFound,
			numObjects:         map[string]*big.Float{"a": big.NewFloat(1.23)},
			expectedNumObjects: map[string]*big.Float{"a": big.NewFloat(1.23)},
		},
	}

	for idx, tc := range testCases {
		t.Run(fmt.Sprintf("TC%02d-%s", idx, tc.name), func(t *testing.T) {
			ctx := ctxlib.Background()
			repo := &BignumMemoryRepo{
				numObjects: tc.numObjects,
			}
			err := repo.DeleteNum(ctx, tc.numObjectName)
			if err != tc.expectedErr {
				t.Errorf("Expected error %v, got %v", tc.expectedErr, err)
			}
			if err == nil {
				if !reflect.DeepEqual(tc.expectedNumObjects, repo.numObjects) {
					t.Errorf("Expected map %v, got %v", tc.expectedNumObjects, repo.numObjects)
				}
			}
		})
	}
}
