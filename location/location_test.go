package location

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCEP_Validate(t *testing.T) {
	tests := []struct {
		name        string
		cep         CEP
		expectError bool
		errorMsg    string
	}{
		{
			name:        "CEP válido com 8 dígitos",
			cep:         CEP("12345678"),
			expectError: false,
		},
		{
			name:        "CEP válido com formato brasileiro (com hífen)",
			cep:         CEP("12345-678"),
			expectError: false,
		},
		{
			name:        "CEP com letras",
			cep:         CEP("1234567a"),
			expectError: true,
			errorMsg:    "cep cannot contain letters",
		},
		{
			name:        "CEP com letras no meio",
			cep:         CEP("123a5678"),
			expectError: true,
			errorMsg:    "cep cannot contain letters",
		},
		{
			name:        "CEP muito curto",
			cep:         CEP("1234567"),
			expectError: true,
			errorMsg:    "cep must be 8 numeric characters long",
		},
		{
			name:        "CEP muito longo",
			cep:         CEP("123456789"),
			expectError: true,
			errorMsg:    "cep must be 8 numeric characters long",
		},
		{
			name:        "CEP vazio",
			cep:         CEP(""),
			expectError: true,
			errorMsg:    "cep must be 8 numeric characters long",
		},
		{
			name:        "CEP com espaços",
			cep:         CEP("12345 678"),
			expectError: false, // Espaços serão removidos pelo regex
		},
		{
			name:        "CEP com pontos",
			cep:         CEP("123.45.678"),
			expectError: false, // Pontos serão removidos pelo regex
		},
		{
			name:        "CEP só com caracteres especiais",
			cep:         CEP("!@#$%^&*"),
			expectError: true,
			errorMsg:    "cep must be 8 numeric characters long",
		},
	}

	for _, item := range tests {
		err := item.cep.Validate()
		if item.expectError {
			require.Error(t, err, "Era esperado um erro para o CEP: %s", item.cep)
			assert.Equal(t, item.errorMsg, err.Error(), "Mensagem de erro não confere para:" + item.cep)
		} else {
			assert.NoError(t, err, "Não era esperado erro para o CEP válido: %s", item.cep)
		}
	}
}
