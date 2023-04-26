# Trabalho Prático: Um estudo empírico sobre programação concorrente com threads

<sub>Última atualização: 26/04/2023</sub>

## Sumário

- [Visão geral e objetivos](#visão-geral-e-objetivos)  
- [Introdução](#introdução)
- [Tarefas](#tarefas)
  - [Implementação](#implementação)
  - [Experimentação](#experimentação)
  - [Relato](#relato)
- [Autoria e política de colaboração](#autoria-e-política-de-colaboração)
- [Entrega](#entrega)
- [Avaliação](#avaliação)
- [Dúvidas e informações](#dúvidas-e-informações)

## Visão geral e objetivos

O objetivo deste trabalho é verificar se é válida hipótese de que programas concorrentes implementados utilizando *threads* apresentam um melhor desempenho com relação a tempo de execução quando comparados a suas respectivas versões sequenciais. Para tal, será realizado um estudo empírico considerando a implementação das versões sequencial e concorrente de um algoritmo e realizados testes para diferentes cargas (*workloads*). Através da análise dos resultados relativos a tempo de execução registrado em cada um dos testes, será possível concluir acercada hipótese inicialmente dada.

## Introdução

O *merge sort* é um conhecido eficiente algoritmo de ordenação baseado no paradigma de divisão e conquista (em inglês, *divide-and-conquer*) proposto por John von Neumann em 1945. Na perspectiva desse paradigma, o problema é recursivamente decomposto em dois ou mais subproblemas menores até que sejam resolvidos de forma simples e direta, e as soluções desses subproblemas são então combinadas até comporem a solução para o problema original. No caso do *merge sort*, o (sub)arranjo de elementos a serem ordenados é particionado, de forma recursiva, em duas partes até que se tenha um subarranjo composto de um único elemento, o qual já se considera ordenado. Em seguida, uma operação chamada merge une duas partes ordenadas do arranjo em uma outra também ordenada. A figura a seguir ilustra como o *merge sort* funciona.

<p align="center">
   <img width="40%" src="https://upload.wikimedia.org/wikipedia/commons/thumb/e/e6/Merge_sort_algorithm_diagram.svg/1200px-Merge_sort_algorithm_diagram.svg.png" alt="Merge sort algorithm diagram" /><br/>
 <span><em>Exemplo de execução do algoritmo <em>merge sort</em> para ordenação <em>top-down</em> de um arranjo de inteiros com sete elementos (fonte: Wikipedia)</em></span>
</p>

Independentemente do caso (melhor ou pior), a complexidade do *merge sort* em termos de tempo de execução é sempre da ordem de $n * \log n$, em que $n$ é o número de elementos do arranjo a ser ordenado, fazendo com que esse algoritmo seja uma eficiente opção para a ordenação de conjuntos de dados. Além disso, um diferencial importante do *merge sort* é que ele é considerado um algoritmo estável, isto é, a ordenação mantém a ordem dos elementos iguais em que caso o elemento mais à esquerda seja menor ou igual ao mais à direita, ele deve ser colocado primeiro no arranjo ordenado. 


## Tarefas

### Implementação

Dentre as tarefas deste trabalho, deverão ser implementadas uma versão sequencial e uma versão concorrente do algoritmo de ordenação *merge sort*. O programa deverá receber como entrada um arquivo de texto dentre os disponíveis no diretório [`input`](input) deste repositório, cada um deles contendo um conjunto de valores inteiros a ser ordenado. Para cada arquivo, o programa deverá (i) ler os valores a partir do arquivo, (ii) realizar a ordenação dos valores, (iii) medir o tempo de execução despendido pela ordenação e (iv) gerar, como saída, um novo arquivo no diretório [`output`](output) com o conjunto ordenado de valores e o mesmo nome do arquivo fornecido como entrada. Ou seja, se o arquivo `A.in` foi utilizado como entrada, deve ser gerado um arquivo `A.out` como saída.

Especificamente para a versão concorrente da implementação do *merge sort*, deverão ser utilizadas *threads* **apenas** na etapa de ordenação propriamente dita, ou seja, a entrada e saída de dados devem ser sequenciais. Por questões de simplicidade, o procedimento de junção das partições ordenadas do arranjo (*merge*) deve ser feita de forma sequencial/iterativa, de modo que essencialmente novas *threads* serão alocadas apenas para tratar as partições do arranjo que estão sendo ordenadas.

Os programas poderão ser implementados utilizando facilidades providas pelas linguagens de programação C/C++, Java **ou** Python. Outra linguagem de programação diferente dessas três poderá ser utilizada contanto a proposta seja previamente validada com o docente. O desenvolvimento da solução deve de antemão visar pela busca de desenvolvimento de software de qualidade, isto é, funcionando correta e eficientemente, exaustivamente testado, bem documentado e com tratamento adequado de eventuais exceções. Por fim, é importante destacar que cada uma das versões deverá ser implementada como um **programa distinto**, resultando em **dois executáveis independentes entre si**.

### Experimentação

A tarefa de experimentação consiste em realizar sucessivas execuções das versões sequenciais e concorrente do algoritmo com vistas a realizar uma análise quantitativa entre elas. Um total de dez cenários deverá ser considerado na experimentação, cada um deles associado a uma quantidade $2^x$ de números inteiros a serem ordenados, sendo $10 \leq x < 20$, os quais estão dispostos em arquivos individuais disponíveis no diretório [`input`](input) deste repositório e estão nomeados alfabeticamente (`A`, `B`, `C`, ...). Assim sendo, o primeiro cenário, associado ao arquivo de entrada [`A.in`](input/A.in), considera a ordenação de um conjunto de $2^{10} = 1024$ valores; o segundo, associado ao arquivo de entrada [`B.in`](input/B.in), contém $2^{11} = 2048$ valores; e assim sucessivamente.

No intuito de tornar os experimentos representativos do ponto de vista estatístico, será necessário executar cada uma das versões em um total de vinte vezes para cada cenário. Os tempos de execução de cada versão em cada uma dessas vinte execuções deverá ser registrado a fim de calcular os valores máximo, mínimo e médio dos tempos nas vinte execuções, além do desvio padrão observado. É importante utilizar uma unidade de medida em um nível de granularidade que torne possível a observação dos valores. Por exemplo, a utilização de segundos (ou mesmo milissegundos, dependendo do caso) como unidade de medida pode resultar em valores significativamente pequenos que sejam expressos praticamente como zero, o que não é desejável para o propósito da análise.

A equação a seguir pode ser utilizada para calcular o desvio padrão $\sigma$ dos tempos de execução das versões sequencial e concorrente:

$$ \sigma = \sqrt{\frac{1}{20} \sum_{i=1}^{20} \(t_i - t_M\)^2} $$

sendo $t_i$ o tempo registrado na *i*-ésima execução da versão em questão e $t_M$ o tempo médio das vinte execuções. Um baixo desvio padrão indica que os pontos dos dados tendem a estar próximos da média desses dados, enquanto um alto desvio padrão indica que os pontos dos dados estão espalhados por uma ampla gama de valores.

Especificamente com relação à análise comparativa entre as versões sequencial e concorrente, é necessário ainda determinar o ganho de desempenho (*speed-up*) eventualmente obtido com a versão concorrente. Esse cálculo pode ser realizado através da seguinte equação:

$$ S = \frac{T_s}{T_c} $$

sendo $S$ o *speed-up*, $T_s$ o tempo médio despendido pela versão sequencial e $T_c$ o tempo médio despendido pela versão concorrente. Caso o valor do *speed-up* seja superior a 1, tem-se que a versão concorrente é, em média, de fato melhor em termos de desempenho do que a sequencial.

**Observação:** É bem sabido que o início da execução de um programa implementado na linguagem de programação Java é afetado de forma relativamente prejudicial pelo carregamento de classes na memória realizado pela máquina virtual Java (JVM) antes da execução propriamente dita do programa, ao que se chama *warm-up*. Com isso, apenas após esse processo de carregamento ter sido concluído é que se pode mensurar de forma confiável o desempenho do programa. Caso os programas objeto deste trabalho tenham sido implementados nessa linguagem de programação, a estratégia mais simples para uma medição confiável (uma vez que os programas em questão não possuem requisitos estritos de latência) é desconsiderar primeiras execuções na etapa de experimentação, justamente pelo fato de os tempos de execução observados serem certamente influenciados pelo tempo de *warm-up* da JVM. Assim sendo, serão necessárias execuções adicionais de cada um dos programas, sendo, contudo, consideradas na computação das estatísticas apenas as últimas vinte execuções.

### Relato

Uma vez realizadas as tarefas de implementação e de experimentação, deverá ser elaborado um relatório contendo, no mínimo, as seguintes seções:

1. **Introdução.** Apresentar uma visão geral do programa implementado, inclusive descrevendo o serviço utilizado como objeto e como o programa realiza o tratamento e apresentação das respostasobtidas.
2. **Metodologia.** Indicar o método adotado para realizar os experimentos com as versões e analisar os resultados obtidos. Deverá ser apresentada a caracterização técnica do computador utilizado (processador, sistema operacional, quantidade de memória RAM), a linguagem de programação e a versão do compilador empregados, os cenários considerados, entre outras informações. Deverão também ser descritos qual o procedimento adotado para gerar os resultados, como a comparação entre as versões foi feita etc.
3. **Resultados.** Apresentar os resultados obtidos na forma de gráfico de linha e tabelas. Para cada solução deverá ser apresentada uma tabela contendo os tempos mínimo, médio e máximo obtidos em cada cenário de experimentação, além dos valores de desvio padrão. Por sua vez, os gráficos de linha deverão exibir apenas os tempos médios obtidos nas execuções de cada versão.
4. **Conclusões.** Discutir os resultados, ou seja, o que foi possível concluir através dos resultados obtidos através dos experimentos, incluindo uma análise do ganho de desempenho (*speed-up*).

## Autoria e política de colaboração

Este trabalho deverá necessariamente ser realizado em equipe composta de **até dois estudantes**, sendo importante, dentro do possível, dividir as tarefas igualmente entre os integrantes da equipe. Após a implementação das soluções para os problemas propostos, o arquivo [`author.md`](https://github.com/ufrn-concprog/mergesort/tree/master/author.md) presente no repositório deverá ser editado preenchendo as informações de identificação dos integrantes da equipe, na seção [Informações de Autoria](https://github.com/ufrn-concprog/mergesort/tree/master/author.md#identificação-de-autoria).

O trabalho em cooperação entre estudantes da turma é estimulado, sendo admissível a discussão de ideias e estratégias. Contudo, tal interação não deve ser entendida como permissão para utilização de (parte de) código fonte de colegas, o que pode caracterizar situação de plágio. **Trabalhos copiados no todo ou em parte de outros colegas ou da Internet serão anulados e receberão nota zero.**

## Entrega

O sistema de controle de versões [Git](https://git-scm.com) e o serviço de hospedagem de repositórios [GitHub](https://git-scm.com) serão utilizados para possibilitar a entrega da implementação realizadas. Para possibilitar a associação de repositórios Git para cada equipe e reuni-los sob uma mesma infraestrutura, foi criada uma atividade (*assignment*) no GitHub Classroom. Cada integrante de equipe deverá acessar este [*link*](https://classroom.github.com/a/iglvzwKC), aceitar o convite para ingressar no GitHub Classroom e finalmente seguir as instruções em tela para acessar a atividade e ingressar em uma equipe existente ou criar outra. Este [vídeo](https://youtu.be/ObaFRGp_Eko) demonstra como ocorre esse processo.

No momento de criação de uma equipe, o GitHub Classroom cria um repositório Git privado acessível unicamente pelos integrantes da equipe e pelo docente, sob a organização [`ufrn-concprog`](https://github.com/ufrn-concprog). A fim de garantir a boa manutenção do repositório, deve-se ainda configurar corretamente o arquivo `.gitignore` para desconsiderar arquivos que não devam ser versionados, a exemplo dos arquivos executáveis gerado a partir da compilação do código fonte. Também não é necessário versionar os arquivos de saída resultantes do processo de ordenação.

A entrega deste trabalho deverá ser realizada até as **23:59 do dia 7 de maio de 2023** no respectivo repositório Git da equipe. O relatório elaborado, preferencialmente em formato *Adobe Portable Format* (PDF), deverá ser enviado através da opção *Tarefas* da Turma Virtual do SIGAA, juntamente com, no campo *Comentários*, o endereço do repositório. Um único membro da equipe deve realizar esse envio e não serão aceitos envios por outros meios ou repositórios que não sejam os descritos nesta especificação.

## Avaliação

A avaliação do trabalho contabilizará nota de até 4,0 pontos na 1ª Unidade da disciplina. O trabalho será avaliado de acordo com os seguintes critérios:

- utilização correta dos conceitos de *threads* vistos na disciplina;
- corretude da execução do programa implementado, que deve apresentar saída em conformidade com a especificação;
- aplicação de boas práticas de programação, incluindo legibilidade, organização e documentação de código fonte;
- correta utilização do repositório Git, no qual deverá ser registrado todo o histórico da implementação por meio de *commits*, e;
- qualidade do relatório produzido

## Dúvidas e informações

Caso haja qualquer dúvida, questionamento ou necessidade de informação adicional, é possível:

- enviar *e-mail* para o endereço everton.cavalcante@ufrn.br;
- enviar mensagem privada diretamente ao docente, utilizando o servidor Discord;
- enviar mensagem no canal de texto `#duvidas` no servidor Discord, ou;
- agendar encontros síncronos pelo canal de áudio `Fale com Prof. Everton` no servidor Discord.
