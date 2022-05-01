package main

import (
	"fmt"

	c "github.com/ostafen/clover"
)

func main() {

	db, _ := c.Open("database")
	defer db.Close()

	db.CreateCollection("tabela")

	p1 := c.NewDocument()
	p1.Set("tabela", "tb_apolice")
	p1.Set("coluna", "nr_apolice")
	p1.Set("tipo", "number(4)")
	p1.Set("order", 1)

	p2 := c.NewDocument()
	p2.Set("tabela", "tb_apolice")
	p2.Set("coluna", "cd_cliente")
	p2.Set("tipo", "number(4)")
	p2.Set("order", 2)

	p3 := c.NewDocument()
	p3.Set("tabela", "tb_apolice")
	p3.Set("coluna", "vr_premio")
	p3.Set("tipo", "number(12,2)")
	p3.Set("order", 3)

	p4 := c.NewDocument()
	p4.Set("tabela", "tb_cliente")
	p4.Set("coluna", "cd_cliente")
	p4.Set("tipo", "number(4)")
	p4.Set("order", 1)

	p5 := c.NewDocument()
	p5.Set("tabela", "tb_cliente")
	p5.Set("coluna", "nm_cliente")
	p5.Set("tipo", "varchar(100)")
	p5.Set("order", 2)

	p6 := c.NewDocument()
	p6.Set("tabela", "tb_cliente")
	p6.Set("coluna", "nr_cpf")
	p6.Set("tipo", "number(11)")
	p6.Set("order", 3)

	p7 := c.NewDocument()
	p7.Set("tabela", "tb_sinistro")
	p7.Set("coluna", "nr_apolice")
	p7.Set("tipo", "number(11)")
	p7.Set("order", 1)

	// db.InsertOne("tabela", p1)
	// db.InsertOne("tabela", p2)
	// db.InsertOne("tabela", p3)
	// db.InsertOne("tabela", p4)
	// db.InsertOne("tabela", p5)
	// db.InsertOne("tabela", p6)

	// Update
	updates := make(map[string]interface{})
	updates["tipo"] = "number(6)"

	db.Query("tabela").Where(c.Field("tabela").Eq("tb_apolice")).Where((*c.Criteria)(c.Field("coluna").Eq("nr_apolice"))).Update(updates)

	// Delete
	db.Query("tabela").Where(c.Field("tabela").Eq("tb_sinistro")).Delete()

	// Find...
	// FindAll -> db.Query("tabela").FindAll()
	// Where -> db.Query("tabela").Where(c.Field("tabela").Eq("tb_cliente")).FindAll()
	// OrderBy -> db.Query("tabela").Sort(c.SortOption{Field: "tabela", Direction: 1}, c.SortOption{Field: "order", Direction: 1}).FindAll()

	tabelas, _ := db.Query("tabela").Sort(c.SortOption{Field: "tabela", Direction: 1}, c.SortOption{Field: "order", Direction: 1}).FindAll()
	Tabela := &struct {
		Tabela string `json:"tabela"`
		Coluna string `json:"coluna"`
		Tipo   string `json:"tipo"`
		Order  int    `json:"order"`
	}{}

	for _, p := range tabelas {
		p.Unmarshal(Tabela)
		p := fmt.Sprintf("%s %s %s %d", Tabela.Tabela, Tabela.Coluna, Tabela.Tipo, Tabela.Order)
		println(p)
	}

}
