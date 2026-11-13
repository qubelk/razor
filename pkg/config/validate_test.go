package config

import (
	"fmt"
	"testing"
)

func TestValidateCompiler(t *testing.T) {
	testCases := []struct {
		compiler string
		wantErr  bool
	}{
		{
			compiler: "gcc",
			wantErr:  false,
		},
		{
			compiler: "clang",
			wantErr:  false,
		},
		{
			compiler: "tcc",
			wantErr:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.compiler, func(t *testing.T) {
			t.Parallel()
			cfg := Config{
				Conf: TomlConfig{
					Compiler: tc.compiler,
				},
			}

			err := cfg.validateCompiler()
			if (err != nil) != tc.wantErr {
				t.Errorf("WantErr %v but got %v\n", tc.wantErr, err)
			}
		})
	}
}

func TestValidateStandart(t *testing.T) {
	testCases := []struct {
		standart int
		wantErr  bool
	}{
		{
			standart: 89,
			wantErr:  false,
		},
		{
			standart: 99,
			wantErr:  false,
		},
		{
			standart: 11,
			wantErr:  false,
		},
		{
			standart: 17,
			wantErr:  false,
		},
		{
			standart: 23,
			wantErr:  false,
		},
		{
			standart: 29,
			wantErr:  true,
		},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("c%d", tc.standart), func(t *testing.T) {
			t.Parallel()
			cfg := Config{
				Conf: TomlConfig{
					Standard: tc.standart,
				},
			}

			err := cfg.validateStandard()
			if (err != nil) != tc.wantErr {
				t.Errorf("WantErr %v but got %v\n", tc.wantErr, err)
			}
		})
	}
}
