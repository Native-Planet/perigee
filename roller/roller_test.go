package roller

import (
	"testing"
)

func TestKeyfile(t *testing.T) {
	tests := []struct {
		name     string
		crypt    string
		auth     string
		ship     interface{}
		revision int
		want     string
		wantErr  bool
	}{
		{
			name:     "Known good input",
			crypt:    "a5f1a3a6662b675c109be8ce642c1b4add53df18609738b8d59bbc1d8d508831",
			auth:     "bd44cee24b74632e65d842ce0fc0bf6e109675e6aca4a9d94c33657e1cc24231",
			ship:     "~zod",
			revision: 1,
			want:     "0w1kLzh.QPclI.WU8jv.hDcxo.dFmWF.XUMMi.VNsqI.Tu3Iq.Eh1zu.ECtN9.rENBP.bI8ms.7U5-T.24IWY.Rpile.OC6ra.~3C4x.6a428.0cmr5",
			wantErr:  false,
		},
		{
			name:     "Invalid Ship",
			crypt:    "01",
			auth:     "02",
			ship:     "invalid",
			revision: 1,
			want:     "",
			wantErr:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Keyfile(tt.crypt, tt.auth, tt.ship, tt.revision)
			if (err != nil) != tt.wantErr {
				t.Errorf("Keyfile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Keyfile() = %v, want %v", got, tt.want)
			}
		})
	}
}
