# PBR-teste
Projetos feito baseado em material online com o objetivo de estudar e aprender uma nova tecnologia, buscando-se assim uma oportunidade na performance Brasil (PBR)

This project was bootstrapped with [Create React App](https://github.com/facebook/create-react-app).

## 1° Questão - projeto 1, utilizando golang

### `Metodologia de busca`

Por não ter nenhuma experiência com a linguagem ( tendo uma ajuda  no entendimento via a familiaridade com a linguagem C) GoLang  optei por procurar tutoriais, video aula e documentação da tecnologia para ter um auxílio na construção do projeto, assim poderia trazer algo produtivo para a resolução da atividade e para o meu aprendizado pessoal.

### `Projeto`
API com cadastro de pessoas:

(descrição) - Trata-se de um sistema simples focado em back-end, onde busco criar um usuário / pessoa.

### `Testar`
OBS: os testes poderão ser realizados por ferramentas apropriadas para testes que  utilizem das rotas criadas no desenvolvimento, ferramentas indicadas são o insomnia ou postman, como utilizei na 2° questão projeto 1 uma (o próximo projeto da lista) metodologia de passo a passo de teste, para api passarei apenas as rotas para requisição.

rotas a serem utilizadas : localhost:3838/persons , os métodos disponíveis são POST E GET.

## 2° Questão - projeto 1, utilizando golang

### `Projeto`
Crud de filmes:

(descrição) - Trata se de um sistema simples focado em back-end, onde busco criar, editar, deletar e buscar/listar filmes (identificado nas linhas de código como movie, sua escrita na língua inglesa);

### `Testar`
Obs: os testes poderão ser realizados por ferramentas apropriadas para testes que  utilizem das rotas criadas no desenvolvimento, ferramentas indicadas são o insomnia ou postman.

1° start o seu projeto utilizando o comando go run main.go

2° teste em seu navegador se o projeto foi validado ao startar digitando o endereço http://localhost:3000/api/v1/movies.

3° Agora utilizando um aplicativo de teste (ex.: postman,..), crie uma nova requisição com o método POST na rota http://localhost:3000/api/v1/movies, selecione a opção body e passe um json com suas informações correspondentes ao exemplo abaixo, o resultado deve retornar um json de informações construídas em sua base de dados :

{
  "name": "Senhos dos anéis - O retorno do rei ",
  "description": "O Senhor dos Anéis: O Retorno do Rei é uma conclusão emocionante e satisfatória para uma grande trilogia". O filme possui uma pontuação de 94 em 100 no Metacritic",
  "thumb_image": "https://www.google.com/imgres?imgurl=https%3A%2F%2Fupload.wikimedia.org%2Fwikipedia%2Fpt%2F0%2F0d%2FEsdlaIII.jpg&imgrefurl=https%3A%2F%2Fpt.wikipedia.org%2Fwiki%2FThe_Lord_of_the_Rings%3A_The_Return_of_the_King&tbnid=X5--ya-kfzowyM&vet=12ahUKEwiM4_zOvd7tAhUFELkGHcWrAscQMygAegUIARCqAQ..i&docid=GI4zg7PjhoUfaM&w=310&h=459&q=senhor%20dos%20aneis%20o%20retorno%20do%20rei&client=firefox-b-d&ved=2ahUKEwiM4_zOvd7tAhUFELkGHcWrAscQMygAegUIARCqAQ",
  "active": true
}

4° Após a criação da lista de filmes, você poderá acessar a rota http://localhost:3000/api/v1/movies para ter acesso a listagem de todos os filmes cadastrados 

5° Caso você queira ver um em especifico a rota continua a mesma porém, deverá ser passado o id gerado por o filme, modificando sua rota para http://localhost:3000/api/v1/movies/{id}.

6° Para atualização dos dados, você irá trazer um json similar ao do método POST, porém, com a adição do id na primeira linha do json, a rota utilizada é a http://localhost:3000/api/v1/movies/{id} e o método utilizado será o  é PUT  

7°  A requisição de deletar a informação é simples, deverá ser selecionada uma requisição de delete, passando a rota http://localhost:3000/api/v1/movies/{id} especificando seu id e a remoção será realizada.







