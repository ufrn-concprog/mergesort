# Lista 2: Estruturas de dados genéricas em Go

<sub>Última atualização: 16/04/2023</sub>

## Sumário

- [Visão geral e objetivos](#visão-geral-e-objetivos)  
- [Tarefas](#tarefas)
- [Requisitos](#requisitos)
- [Orientações gerais](#orientações-gerais)
- [Exemplos de entradas e saídas](#exemplos-de-entradas-e-saídas)
- [Autoria e política de colaboração](#autoria-e-política-de-colaboração)
- [Entrega](#entrega)
- [Avaliação](#avaliação)
- [Dúvidas e informações](#dúvidas-e-informações)

## Visão geral e objetivos

O objetivo deste exercício de programação é implementar, na linguagem de programação Go, um programa chamado `genmap` para gerenciar elementos de tipos quaisquer a serem armazenados em um *mapa*. Um mapa é uma estrutura de dados que realiza a correspondência de valores para chaves, armazenando, assim, entradas como pares chave--valor.

Este exercício explora os seguintes elementos da programação em Go, cujos conhecimentos são, portanto, ora necessários:

- Entrada e saída formatada via console
- Estruturas de controle de fluxo
- Tipos estruturados (*structs*)
- Funções
- *Slices*
- Manipulação de *strings*
- Interfaces
- Generics

## Tarefas

As seguintes tarefas deverão ser realizadas na implementação do programa `genmap`:

1. Implementar um tipo estruturado (*struct*) genérico chamado `Entry` que represente uma entrada do mapa, composta de uma chave e um valor. A chave deve ser unicamente do tipo inteiro ou *string*, enquanto o valor pode ser de qualquer tipo passível de utilização de operadores relacionais. Por questões de simplicidade, deve-se assumir que os valores são de tipos dentre os tipos primitivos preexistentes em Go.
2. Implementar um tipo estruturado (*struct*) genérico chamado `Map` que representa o mapa propriamente dito, composto por um conjunto de entradas, ou seja, elementos do tipo `Entry`. O número de entradas é indefinido *a priori*.
3. Implementar as seguintes funções genéricas para realizar operações sobre o mapa. A chamada às funções deve ser realizada por meio de uma abreviação (*case insensitive*) a ser informada pelo usuário através da entrada padrão. As abreviações devem ser **exatamente** as listados a seguir.

    1. **_addEntry_ (abreviação _add_):** Adiciona uma nova entrada ao mapa, exibindo na saída padrão o mapa atualizado com a entrada adicionada. A chave e o valor a serem armazenados na entrada deverão ser informados pelo usuário através da entrada padrão. Caso já exista uma entrada para a chave informada, o valor armazenado deve ser sobrescrito pelo novo valor. Esta função deve receber como parâmetro o mapa que será modificado, a chave e o valor a serem adicionados.
    2. **_getEntry_ (abreviação _get_):** Exibe na saída padrão uma entrada armazenada no mapa a partir de uma determinada chave, percorrendo o conjunto de entradas por meio da comparação com a chave informada. Caso a chave não exista, o programa deverá exibir uma mensagem na saída padrão informando que a chave é inexistente.
   3. **_mapSize_ (abreviação _size_):** Exibe na saída padrão o número de entradas atualmente armazenadas no mapa.
   4. **_printMap_ (abreviação _print_):** Exibe na saída padrão todas as entradas atualmente armazenadas no mapa e o número de entradas.

4. O programa deve continuamente aceitar novas abreviações providas pelo usuário através da entrada padrão. O programa deverá ainda exibir uma mensagem de erro caso seja fornecido um número inválido de argumentos a *add* ou a *get*, bem como caso seja fornecida uma abreviação para uma função não implementada. Por fim, caso o usuário informe a abreviação *exit*, o programa deve ser encerrado.

Por questões de simplicidade, na função principal do programa, deve-se admitir que o mapa irá armazenar chaves do tipo inteiro e valores do tipo *string* e, portanto, esses tipos são os que deverão ser utilizados na instanciação do mapa. Contudo, **todos os tipos estruturados e funções devem ser genéricos**.

## Requisitos

A implementação deste trabalho requer os seguintes elementos instalados no ambiente de desenvolvimento:

- [Git](https://git-scm.com), como sistema de controle de versões
- [Go](https://go.dev), incluindo compilador, ambiente de execução e outras ferramentas associadas à linguagem de programação Go

## Orientações gerais

Boas práticas de programação deverão ser constantemente aplicadas no desenvolvimento do programa. Assim, o programa deverá ser codificado de forma legível (com indentação de código fonte, nomes consistentes, etc.) e documentado adequadamente na forma de comentários.

O programa deverá ainda ser desenvolvido com qualidade, garantindo que o ele funcione de forma correta e eficiente. Deve-se também pensar nas possíveis entradas que poderão ser utilizadas para testar apropriadamente o programa, além de serem tratadas adequadamente possíveis entradas consideradas inválidas.

## Exemplos de entradas e saídas

```bash
$ genmap

> print
[]
0

> add 1 "Hello"
[1] "Hello"

> add 2 "World"
[1] "Hello"
[2] "World"

> get 2
[2] "World"

> size
2

> exit
```

Note que a exibição de cada entrada na saída padrão deve obedecer ao formato `[<chave>] valor` (ou seja, com colchetes circundando as chaves). Caso o mapa esteja vazio, devem ser impressos apenas os colchetes seguidos do número zero.

## Autoria e política de colaboração

Este trabalho deverá necessariamente ser realizado em equipe composta de **até dois estudantes**, sendo importante, dentro do possível, dividir as tarefas igualmente entre os integrantes da equipe. Após a implementação das soluções para os problemas propostos, o arquivo [`author.md`](https://github.com/ufrn-golang/lista2/tree/master/author.md) presente no repositório deverá ser editado preenchendo as informações de identificação dos integrantes da equipe, na seção [Informações de Autoria](https://github.com/ufrn-golang/lista2/tree/master/author.md#identificação-de-autoria). 

O trabalho em cooperação entre estudantes da turma é estimulado, sendo admissível a discussão de ideias e estratégias. Contudo, tal interação não deve ser entendida como permissão para utilização de (parte de) código fonte de colegas, o que pode caracterizar situação de plágio. Trabalhos copiados no todo ou em parte de outros colegas ou da Internet serão anulados e receberão nota zero.

## Entrega

O sistema de controle de versões [Git](https://git-scm.com) e o serviço de hospedagem de repositórios [GitHub](https://git-scm.com) serão utilizados para possibilitar a entrega da implementação realizadas. Para possibilitar a associação de repositórios Git para cada equipe e reuni-los sob uma mesma infraestrutura, foi criada uma atividade (*assignment*) no GitHub Classroom. Cada integrante de equipe deverá acessar este [*link*](https://classroom.github.com/a/5aSHjU-U), aceitar o convite para ingressar no GitHub Classroom e finalmente seguir as instruções em tela para acessar a atividade e ingressar em uma equipe existente ou criar outra. Este [vídeo](https://youtu.be/ObaFRGp_Eko) demonstra como ocorre esse processo.

No momento de criação de uma equipe, o GitHub Classroom cria um repositório Git privado acessível unicamente pelos integrantes da equipe e pelo docente, sob a organização [`ufrn-golang`](https://github.com/ufrn-golang). A fim de garantir a boa manutenção do repositório, deve-se ainda configurar corretamente o arquivo `.gitignore` para desconsiderar arquivos que não devam ser versionados, a exemplo do arquivo executável gerado a partir da compilação do código fonte.

A implementação do programa objeto deste trabalho deverá ser realizada **até as 23:59 do dia 18 de abril de 2023** no respectivo repositório Git da equipe. Para fins de registro, o endereço do repositório também deverá ser **obrigatoriamente** enviado através da opção *Tarefas* na Turma Virtual do SIGAA, devendo **um único membro da equipe** realizar esse envio. Além disso, **não serão aceitos envios por outros meios ou repositórios que não sejam os descritos nesta especificação.**

## Avaliação

A avaliação do programa implementado contabilizará nota de até 10,0 pontos. O programa implementado será avaliado de acordo com os seguintes critérios:

- utilização correta dos recursos providos pela linguagem de programação Go;
- corretude da execução do programa implementado, que deve apresentar saída em conformidade com a especificação;
- aplicação de boas práticas de programação, incluindo legibilidade, organização e documentação de código fonte, e;
- correta utilização do repositório Git, no qual deverá ser registrado todo o histórico da implementação por meio de *commits*.

O não cumprimento de algum dos critérios de avaliação especificados poderá resultar nos seguintes decréscimos, aplicados sobre a nota obtida até então na avaliação:

| Falta | Decréscimo |
| :--- | ---: |
| Falta de comentários no código fonte | -10% |
| Uso inadequado de controle de versão com Git | -20% |
| Falta de especificação ou especificação incorreta da autoria do trabalho (arquivo [`author.md`](https://github.com/ufrn-golang/lista2/tree/master/author.md)) | -20% |
| Código fonte com legibilidade prejudicada (por exemplo, com identação ou nomenclatura inadequada) | -30% |
| Implementação realizada em desacordo com as especificações postas para o trabalho | -50% |
| Programa apresenta erros de compilação, não executa ou apresenta saída incorreta | -70% |
| Plagiarismo | -100% |

## Dúvidas e informações

Caso haja qualquer dúvida, questionamento ou necessidade de informação adicional, é possível:

- enviar *e-mail* para o endereço everton.cavalcante@ufrn.br;
- enviar mensagem privada diretamente ao docente, utilizando o servidor Discord;
- enviar mensagem no canal de texto `#duvidas` no servidor Discord, ou;
- agendar encontros síncronos pelo canal de áudio `Fale com Prof. Everton` no servidor Discord.
