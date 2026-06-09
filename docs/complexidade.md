# Análise de Complexidade

## Trie

### Inserção — O(L)

```
para cada rune na palavra:
    se children[rune] == nil: alocar nó   ← O(1) inserção no map
    avançar ponteiro                       ← O(1)
marcar nó terminal
```

L = comprimento da palavra. Uma consulta ao map e uma possível alocação por caractere. N palavras → O(N·L) tempo total de construção.

**Espaço**: O(N·L) no pior caso (sem prefixos compartilhados). Na prática muito menos — prefixos comuns colapsam em nós compartilhados. `map[rune]*Node` aloca apenas as arestas de ramificação existentes, não um array de 26 posições por nó.

### Busca — O(L)

Mesmo percurso da inserção, sem alocação. Retorna `(freq, found)` no nó terminal.

---

## Busca Aproximada por DFS (Levenshtein com carry de linha)

### Ideia central

O DP de Levenshtein padrão constrói uma matriz `(m+1) × (n+1)` comparando duas strings completas. Custo: O(m·n).

Aqui, percorre-se a Trie por DFS carregando apenas a **linha atual do DP** (comprimento `len(query)+1`). A cada nó da Trie com caractere `c`:

```
currRow[0] = prevRow[0] + 1
para i em 1..len(query):
    insert  = currRow[i-1] + 1
    delete  = prevRow[i]   + 1
    replace = prevRow[i-1] + (0 se query[i-1]==c, senão 1)
    currRow[i] = min(insert, delete, replace)
```

Isso computa `Levenshtein(prefixo_atual_na_trie, query[:i])` de forma incremental.

### Condição de poda

```
se min(currRow) > maxDist: retornar   ← poda toda a subárvore
```

**Por que está correto**: cada caractere subsequente pode reduzir a distância de edição em no máximo 1 por passo. Se o valor mínimo da linha atual já excede `maxDist`, nenhum caminho descendente pode trazer a distância de volta ao intervalo permitido.

### Complexidade

| Caso | Tempo |
|------|-------|
| Pior caso (sem poda, trie densa) | O(N·L), onde N = tamanho do dicionário, L = comprimento médio das palavras |
| Caso médio (poda ativa) | O(b^maxDist · L), onde b = fator de ramificação por nó |
| Teste de estresse (200k palavras, maxDist=2) | Sub-segundo — a maioria das subárvores é podada cedo |

**Espaço**: O(profundidade · L) para a pilha de chamadas DFS carregando uma linha por nível. Profundidade ≤ comprimento máximo de palavra.

---

## MergeSort

```
MergeSort(items):
    se len <= 1: retornar items
    mid = len / 2
    esquerda  = MergeSort(items[:mid])
    direita   = MergeSort(items[mid:])
    retornar merge(esquerda, direita)

merge(esquerda, direita):
    enquanto ambos não vazios:
        anexar item de maior frequência   ← ordenação freq desc
    anexar restante
```

**Tempo**: O(N log N) — recorrência clássica de divisão e conquista T(N) = 2T(N/2) + O(N).

**Espaço**: O(N) — `merge` aloca um slice resultado de capacidade `len(esquerda)+len(direita)`. Pilha de recursão é O(log N).

**Estabilidade**: palavras com frequência igual preservam a ordem relativa de inserção (lado esquerdo vence empates via `>=`).

---

## Tabela Resumo

| Operação | Tempo | Espaço |
|----------|-------|--------|
| Trie Inserção | O(L) por palavra, O(N·L) total | O(N·L) pior caso |
| Trie Busca | O(L) | O(1) |
| Busca Aproximada DFS | O(N·L) pior caso, O(b^d·L) médio | O(profundidade·L) pilha |
| MergeSort | O(N log N) | O(N) |

N = tamanho do dicionário, L = comprimento da palavra/query, d = maxDist (2), b = fator de ramificação da Trie.
