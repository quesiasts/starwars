# Requisitos:

- A API deve ser REST
- Cada planeta, os seguintes dados devem ser obtidos do banco de dados da aplicação, sendo inserido manualmente:
Nome
Clima
Terreno
- Para cada planeta também devemos ter a quantidade de aparições em filmes, que podem ser obtidas pela API pública do Star Wars: https://swapi.co/

# Funcionalidades desejadas: 

- Adicionar um planeta (com nome, clima e terreno)
- Listar Planetas
- Buscar por nome
- Buscar por ID
- Remover planeta

Linguagem utilizada: GO

Favor ler o arquivo endPoints
  
### Rodando o sistema

Abra o terminal e rode o comando abaixo para subir a aplicação com o docker
```sh
$ docker-compose up
```

Abra outro terminal e rode o comando abaixo para subir o frontend
```sh
$ go run main.go -- Rodar o sistema.
```

© 2022 GitHub, Inc.
Terms
Privacy