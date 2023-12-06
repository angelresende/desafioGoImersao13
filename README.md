# Ordenação de Dados em Go

Informações do desafio
Objetivo:
Criar um programa em Go que lê informações de um arquivo. Realizar ordenação dos dados lidos por diferentes critérios (por exemplo, ordenação por nome e por idade). A ordenação por nome deve ser feita utilizando a tabela ASCII. Salvar o resultado ordenado em outro arquivo.


Requisitos Técnicos:
Use a biblioteca padrão do Go para manipulação de arquivos. Considere trabalhar com dados estruturados, como linhas de um arquivo CSV.


Estrutura do Programa:
Crie um arquivo de entrada contendo dados (por exemplo, CSV, cada linha representando uma entidade com algumas informações).
Implemente funções que leem o arquivo, realizam a ordenação dos dados por nome e por idade e retornam os resultados.
Salve os resultados ordenados em dois novos arquivos (um para cada critério de ordenação).


Exemplo Arquivo de Entrada (arquivo-origem.csv)
Nome,Idade,Pontuação
<p>• Carlos,30,80 </p>
<p>• Carlos,22,75</p>
<p>• Maria,30,95</p>
<p>• Joao,25,80</p>
<p>• carlos,15,90</p>

Exemplo de Arquivo de Saída(arquivo-destino.csv)
Nome,Idade,Pontuação
<p>• Carlos,22,75</p>
<p>• Carlos,30,80</p>
<p>• carlos,15,90</p>
<p>• Joao,25,80</p>
<p>• Maria,30,95</p>

Observações:
Utilize a lógica de manipulação de arquivos em Go.
Implemente a ordenação por nome e por idade.
Se desejar, crie funções separadas para leitura, processamento e escrita de arquivos.
Lide com erros de maneira apropriada.

Precisamos rodar o comando go run main.go ./arquivo-origem.csv ./arquivo-destino.csv. Será usado um arquivo CSV aleatório nosso para testar sua aplicação.

