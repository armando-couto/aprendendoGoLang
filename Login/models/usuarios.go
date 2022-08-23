package models

import "mesa/db"

type Usuario struct {
	Id, Squad              int
	Nome, Email, Permissao string
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
		var id, squad int
		var nome, email, permissao string

		err = selectDeTodosOsUsuarios.Scan(&id, &nome, &email, &squad, &permissao)
		if err != nil {
			panic(err.Error())
		}

		u.Id = id
		u.Nome = nome
		u.Email = email
		u.Squad = squad
		u.Permissao = permissao

		usuarios = append(usuarios, u)
	}
	defer db.Close()
	return usuarios
}

func CriaNovoUsuario(nome string, email string, squad int, permissao string) {
	db := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := db.Prepare("insert into usuario(nome, email, squad, permissao) values($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}

	insereDadosNoBanco.Exec(nome, email, squad, permissao)
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
		var id, squad int
		var nome, email, permissao string

		err = usuarioDoBanco.Scan(&id, &nome, &email, &squad, &permissao)
		if err != nil {
			panic(err.Error())
		}
		usuarioParaAtualizar.Id = id
		usuarioParaAtualizar.Nome = nome
		usuarioParaAtualizar.Email = email
		usuarioParaAtualizar.Squad = squad
		usuarioParaAtualizar.Permissao = permissao
	}
	defer db.Close()
	return usuarioParaAtualizar
}

func AtualizaUsuario(id int, nome string, email string, squad int, permissao string) {
	db := db.ConectaComBancoDeDados()

	AtualizaUsuario, err := db.Prepare("update usuario set nome=$1, email=$2, squad=$3, permissao=$4 where id=$5")
	if err != nil {
		panic(err.Error())
	}
	AtualizaUsuario.Exec(nome, email, squad, permissao, id)
	defer db.Close()
}
