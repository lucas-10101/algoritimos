# Repositório de código para algoritimos encontrados durante a leitura do livro "Algoritimos Teoria e prática 3º edição"

Inicialmente as implementações serão feitas em golang

## Por que golang

Entre os motivos que eu posso listar para o por que utilizar golang para criação deste projeto, posso destacar:

- Sua natureza baixo nível
- Performance nas operações
- Simplicidade no gerenciamento de memória
- Menos "coisas prontas"

Assim é possível focar mais em como fazer do que quais bibliotecas utilizar

## Testes

Para validação das implementações também foram adicionados testes unitários para as funções seguindo o padrão da linguagem

## Funções implementadas

#### Insertion Sort (sort/almost_selection_sort.go) : Um quase algoritimo de ordenação por seleção

#### Insertion Sort (sort/insertion_sort.go) : Ordenação de inteiros em ordem crescente

#### Linear Search (search/linear_search.go) : Busca linear de indice em vetor, tempo de execução O(n)

#### BinaryIntSum (binaries/binary_int_sum.go): Realiza a soma entre dois vetores de bits (representados por booleans) (fui lembrar do byte depois, porém esta ótimo para teste de lógica)

#### Selection Sort (sort/select_sort.go) : Ordenação de arranjos por seleção de menor elemento O(n^2)
