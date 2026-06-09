# Projeto 6 — Corretor Ortográfico com Busca Aproximada

**Disciplina:** Estrutura de Dados Avançada  
**Aluno:** Marcelo Ramalho

---

## Descrição

Módulo preditivo de correção ortográfica com tolerância a erros de digitação. Implementa:

- **Trie** (`map[rune]*Node`) - carregamento do dicionário com compressão de prefixos
- **Busca aproximada por DFS** - Levenshtein com carry de linha e poda (≤ 2 edições)
- **MergeSort** - ordenação das sugestões por frequência decrescente (implementação manual)

---

## Pré-requisitos

- [Go](https://go.dev/doc/install) — versão 1.21 ou superior
- [Make](https://www.gnu.org/software/make/) — para execução via Makefile

---

## Compilação

```bash
make build
```

Gera os binários `bin/spell-checker` e `bin/gen-dict`.

---

## Execução

**Via Makefile:**

```bash
make run input=data/input_basico.json
make run input=data/input_avancado.json
make run input=data/input_estresse.json
```

**Via binário direto:**

```bash
bin/spell-checker -input data/input_basico.json
bin/spell-checker -input data/input_avancado.json
bin/spell-checker -input data/input_estresse.json
```

O arquivo de saída é gerado automaticamente em `out/output_<nivel>.json`.

---

## Testes

```bash
make test
```

Executa os testes unitários de Trie, busca aproximada e MergeSort.

---

## Geração do dicionário de estresse (tooling interno)

O arquivo `data/input_estresse.json` (200 mil palavras, 1 mil queries) foi gerado com:

```bash
make build
make gen
```

Isso executa `bin/gen-dict` e sobrescreve `data/input_estresse.json`.

---

## Estrutura do repositório

```
/
├── src/
│   ├── cmd/spell-checker/main.go
│   ├── trie/          # Trie + Node
│   ├── fuzzy/         # DFS + Levenshtein
│   ├── sort/          # MergeSort
│   └── io/            # codec JSON
├── data/
│   ├── input_basico.json
│   ├── input_avancado.json
│   ├── input_estresse.json
│   ├── output_esperado_basico.json
│   └── output_esperado_avancado.json
├── docs/
│   └── complexity.md
├── tools/
│   └── gen_dict.go
├── Makefile
└── README.md
```

---

## Nota sobre os gabaritos

`output_esperado_estresse.json` não está incluído no repositório (4,2 MB). Para reproduzi-lo:

```bash
make run input=data/input_estresse.json
# saída gerada em out/output_estresse.json
```

---

## Formato de entrada/saída

**Entrada:**
```json
{"dictionary": [{"word": "hello", "freq": 1500}], "queries": ["helo"]}
```

**Saída:**
```json
{"results": [{"query": "helo", "suggestions": ["hello"]}]}
```

Sugestões = palavras com distância de edição ≤ 2, ordenadas por frequência decrescente.
