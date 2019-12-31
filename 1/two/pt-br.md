# --- Segunda Parte ---

Durante a segunda verificação de lançamento, o Elfo responsável pela Equação de Verificação Dupla do Foguete pára a sequência de lançamento. Aparentemente, você esqueceu de incluir combustível a mais para o combustível que acabou de adicionar.

O combustível em si necessita combustível como um módulo - pega a massa, divide por três, arredonda para baixo e subtrai 2. No entanto, o combustível também requisita de combustível, e esse combustível requisita combustível, e assim por diante. Qualquer massa que requisitar combustível negativo deveria ser tratada como se requisitasse nenhum combustível; a massa restante, se houver, é lidada com muito desejo, o que não tem massa e está fora do escopo de cálculo.

Então, para cada módulo de massa, calcule seu combustível e adicione para o total. Depois, trate a quantidade de combustível que você acabou de calcular como entrada de massa e repita o processo, continuando até o requerimento de combustível ser zero ou negativo. Por exemplo:

Um módulo de massa 14 requer 2 de combustível. Esse combustível não requisita mais combustível (2 dividido por 3 e arredondado para baixo dá 0, o que se enquadra no combustível negativo), então o combustível total requisitado ainda é apenas 2.
A princípio, um módulo de massa 1969 requer 654 de combustível. Depois, esse combustível requisita mais 216 de combustível (654 / 3 - 2). 216 requisita mais 70 de combustível, que requer 21 de combustível, que requer 5 de combustível, que não reque mais nenhum combustível. Logo, o combustível total necessário para um módulo de massa 1969 é 654 + 216 + 70 + 21 + 5 = 966.
O combustível requisitado por um módulo de massa 100756 e seu combustível é: 33583 + 11192 + 3728 + 1240 + 411 + 135 + 43 + 12 + 2 = 50346.
Qual é a soma de combustível requerido para todos os módulos da sua espaçonave quando levamos em conta a massa do combustível adicionado? (Calcule o combustível requisitado para cada módulo separadamente, então some tudo no final.)