# Go com Local Database CloverDB

[CloverDB](https://github.com/ostafen/clover)

## Criação de uma instância

```go
db, _ := c.Open("database")
defer db.Close()
```

Será criado um diretório chamado "database" com os arquivos de banco de dados.

## Criação de um documento (tabela) e inclusão de um registro

```go
db.CreateCollection("tabela")

p1 := c.NewDocument()
p1.Set("tabela", "tb_apolice")
p1.Set("coluna", "nr_apolice")
p1.Set("tipo", "number(4)")
p1.Set("order", 1)

db.InsertOne("tabela", p1)
```

Mais operações, verificar o código, por favor.
