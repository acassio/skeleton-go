package enum

//ERRORS é um enumerador de erros
var ERRORS = map[string]int{
	Empty:           0,
	InvalidResponse: 1,
}

//InvalidResponse quando o response é diferente de 200
const InvalidResponse = "Response inválido."

//Empty quando nao ha mensagem
const Empty = ""
