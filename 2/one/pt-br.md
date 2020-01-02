# --- Dia 2: Alarme do Programa 1202 ---

No seu trajeto para assistência gravitacional ao redor da Lua, o computador da sua nave apita repetidamente sobre um "alarme do programa 1202". No rádio, um elfo já explica como lidar com a situação: "Não se preocupe, isso é perfeitamente norma--" O computador explode.

Você notifica os elfos que a fumaça mágica do computador parece ter escapado. "Esse computador executava programas Intcode enquanto o programa de assistência gravitacional funcionava; com certeza há peças sobrando por aqui para construir um novo computador Intcode!"

Um programa Intcode é uma lista de inteiros separados por vírgula (como 1,0,0,3,99). Para executar um, é necessário verificar o primeiro inteiro (chamado posição 0). Aqui, você vai encontrar um opcode - seja 1, 2 ou 99. O opcode indica o que fazer; por exemplo, 99 significa que o programa está finalizado e deve ser parado imediatamente.
Encontrar um código desconhecido significa que algo deu errado.

Opcode 1 adiciona números lidos das duas primeiras posições e armazena o resultado em uma terceira posição.

Os três inteiros imediatamente após o opcode te mostram essas três posições - os dois primeiros indicam as posições que você dev ler os valores de entrada, e o terceiro indica a posição que a saída deve ser armazenada.

Por exemplo, se seu computador Intcode encontrar 1,10,20,30, ele deve ler os valores nas posições 10 e 20, adicionar esses valores, e então sobrescrever o valor na posição 30 com sua soma.

O opcode2 funciona exatamente como o opcode 1, com a diferença que multiplica as duas entradas ao invés de adicioná-las. De novo, os três inteiros após o opcode indicam onde as entradas e saídas ficam, não seus valores.

Uma vez que tenha finalizado um opcode, mude para o próximo pulando 4 posições.

Por exemplo, supondo que você tenha o programa abaixo:

``1,9,10,3,2,3,11,0,99,30,40,50`

Para propósitos de ilustração, aqui está o mesmo programa separado em diversas linhas:

```text
1,9,10,3,
2,3,11,0,
99,
30,40,50
```

Os primeiros quatro inteiros, 1,9,10,3, estão as posições 0,1,2 e 3. No entanto, eles representam o primeiro opcode (1, adição), as posições das duas entradas (9 e 10) e a posição da saída (3). Para lidar com esse opcode, primeiro ovcê precisa obter os valores nas posições de entrada: a posição 9 contém 30 e a posição 10 contém 40. Adicione esses números para obter 70. Depois, armazene esse valor na posição de saída; aqui, a posiçãode saída (3) está na posição 3, então ele apenas se sobrescreve. Depois, o programa deve se parecer com isso:

```text
1,9,10,70,
2,3,11,0,
99,
30,40,50
```

Pule 4 posições para chegar no próximo opcode, 2. O opcode trabalha da mesma forma que o anterior, mas multiplica ao invés de adicionar. As entradas ficam nas posições 3 e 11; essas posições contém 70 e 50 respectivamente. Multiplicá-los resulta em 3500; isso é armazenado na posição 0:

```text
3500,9,10,70,
2,3,11,0,
99,
30,40,50
```

Pular mais 4 posições dá o opcode 99, parando o programa.

Aqui estão os estados inicias e finais de alguns programas menores:

* `1,0,0,0,99` vira `2,0,0,0,99` (1 + 1 = 2).
* `2,3,0,3,99` vira `2,3,0,6,99` (3 * 2 = 6).
* `2,4,4,5,99,0` vira `2,4,4,5,99,9801` (99 * 99 = 9801).
* `1,1,1,4,99,5,6,0,99` vira `30,1,1,4,2,5,6,0,99`.

Quando tiver o computador funcionando, a primeira etapa é restaurar a o programa de assistência gravitacional (seu desafio de entrada) para o estado que o "alarme do programa 1202" tinha antes do último computador pegar fogo. Para fazer isso, antes de executar o programa, substitua a posição 1 com o valor 12 e a posição 2 com o valor 2. Qual valor é deixado na posição 0 após o programa parar?