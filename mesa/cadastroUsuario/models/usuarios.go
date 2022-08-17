package models

import "awesomeProject/Login/db"

type Usuario struct {
	Id, Cpf, Squad    int
	Nome, Categoria   string
}

func BuscaTodosOsUsuarios() []Usuario {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsUsuarios, err := db.Query("select * from usuario")
	if err != nil {
		panic(err.Error())
	}

	u := Usuario{}
	usuarios := []Usuario{}

	for selectDeTodosOsUsuarios.Next() {
		var id, cpf, squad int
		var nome, categoria string


		err = selectDeTodosOsUsuarios.Scan(&id, &nome, &cpf, &squad, &categoria)
		if err != nil {
			panic(err.Error())
		}

		u.Id = id
		u.Nome = nome
		u.Cpf = cpf
		u.Squad = squad
		u.Categoria = categoria

		usuarios = append(usuarios, u)
	}
	defer db.Close()
	return usuarios
}
func CriaNovoUsuario(nome string, cpf, squad int, categoria string ) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into usuario(nome, cpf, squad, categoria) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, cpf, squad, categoria)
	defer db.Close()

}

func DeletaUsuario(id string) {
	db := db.ConectaComBancoDeDados()

	deletarOUsuario, err := db.Prepare("delete from usuario where id=$1")
	if err != nil {
		panic(err.Error())
	}

	deletarOUsuario.Exec(id)
	defer db.Close()

}

func EditaUsuario(id string) Usuario {
	db := db.ConectaComBancoDeDados()

	usuarioDoBanco, err := db.Query("select * from usuario where id=$1", id)
	if err != nil {
		panic(err.Error())
	}

	usuarioParaAtualizar := Usuario{}

	for usuarioDoBanco.Next() {
		var id, cpf, squad int
		var nome, categoria string

		err = usuarioDoBanco.Scan(&id, &nome, &cpf, &squad, &categoria)
		if err != nil {
			panic(err.Error())
		}
		usuarioParaAtualizar.Id = id
		usuarioParaAtualizar.Nome = nome
		usuarioParaAtualizar.Cpf = cpf
		usuarioParaAtualizar.Squad = squad
		usuarioParaAtualizar.Categoria = categoria
	}
	defer db.Close()
	return usuarioParaAtualizar
}

func AtualizaUsuario(id int, nome string, cpf,squad int, categoria string) {
	db := db.ConectaComBancoDeDados()

	AtualizaUsuario, err := db.Prepare("update usuario set nome=$1, cpf=$2, squad=$3, categoria=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaUsuario.Exec(nome, cpf, squad, categoria, id)
	defer db.Close()
}