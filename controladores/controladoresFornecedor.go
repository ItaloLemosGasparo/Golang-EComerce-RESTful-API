package controladores

import (
	"projeto/inicializadores"
	"projeto/modelos"

	"github.com/gin-gonic/gin"
)

func CadastrarFornecedor(c *gin.Context) {
	var fornecedorTemp struct {
		Nome      string
		Email     string
		Telefone  string
		TelefoneB string
		Cpf       string
		Cnpj      string
	}

	c.Bind(&fornecedorTemp)

	fornecedor := modelos.Fornecedor{
		Nome:      fornecedorTemp.Nome,
		Email:     fornecedorTemp.Email,
		Telefone:  fornecedorTemp.Telefone,
		TelefoneB: fornecedorTemp.TelefoneB,
		CPF:       fornecedorTemp.Cpf,
		CNPJ:      fornecedorTemp.Cnpj,
	}

	if result := inicializadores.BD.Create(&fornecedor); result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"Fornecedor": fornecedor,
	})
}

func DeletarFornecedor(c *gin.Context) {
	id := c.Param("id")

	inicializadores.BD.Delete(&modelos.Fornecedor{}, id)

	c.Status(200)
}

func BuscarFornecedores(c *gin.Context) {
	var Fornecedores []modelos.Fornecedor
	inicializadores.BD.Find(&Fornecedores)

	c.JSON(200, gin.H{
		"fornecedores": Fornecedores,
	})
}

func BuscarFornecedore(c *gin.Context) {
	id := c.Param("id")

	var fornecedor modelos.Fornecedor
	inicializadores.BD.First(&fornecedor, id)

	c.JSON(200, gin.H{
		"fornecedor": fornecedor,
	})
}

func AtualizarFornecedor(c *gin.Context) {
	id := c.Param("id")

	var fornecedorTemp struct {
		Nome      string
		Email     string
		Telefone  string
		TelefoneB string
		Cpf       string
		Cnpj      string
	}
	c.Bind(&fornecedorTemp)

	var fornecedor modelos.Fornecedor
	inicializadores.BD.First(&fornecedor, id)

	inicializadores.BD.Model(&fornecedor).Updates(modelos.Fornecedor{
		Nome:      fornecedorTemp.Nome,
		Email:     fornecedorTemp.Email,
		Telefone:  fornecedorTemp.Telefone,
		TelefoneB: fornecedorTemp.TelefoneB,
		CPF:       fornecedorTemp.Cpf,
		CNPJ:      fornecedorTemp.Cnpj,
	})

	c.JSON(200, gin.H{
		"fornecedor": fornecedor,
	})

}