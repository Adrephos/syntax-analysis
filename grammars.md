> No es ni LL(1) ni LR(0)
```
S -> VT
T -> bST | ε
V -> cV | c
```
---
> No es ni LL(1) ni LR(0)
```
S -> aAa
A -> aa | ε
```
---
> No es ni LL(1) ni LR(0)

```
S -> aSb | cSa | ε
```
---
> No es ni LL(1) pero es LR(0)

```
E -> E+T | T
T -> TxF | F
F -> (E) | i
```
---
> No es ni LL(1) ni LR(0)

```
S -> A | a
A -> a
```
---
> No es ni LL(1) pero es LR(0)

> Es recursiva por la izquerda, queda probar si al eliminarla sirve

```
S -> Ab | a
A -> S | ε
```
---
> No es ni LL(1) ni LR(0)
```
S -> iEtSF | a
F -> eS | ε
E -> b
```
---
<span color="red">No es ni LL(1) pero es LR(0)</span>
```
S -> Aa | b
A -> Bd | ε
B -> Sc
```
---
> No es ni LL(1) pero es LR(0)

> Es recursiva por la izquerda, queda probar si al eliminarla sirve

```
S -> Aa | b
A -> Sd | ε
```
---
> No es ni LL(1) pero es LR(0)

> Es recursiva por la izquerda, queda probar si al eliminarla sirve

```
S -> Aa | b
A -> Ac | Sd | ε
```
---
> No es ni LL(1) ni LR(0)

```
S -> ACB | CbB | Ba
A -> da | BC
B -> g | ε
C -> h | ε
```
---
> Es LL(1) no es LR(0)

```
S -> AaAb | BbBa
A -> ε
B -> ε
```
---
> Es LL(1) y LR(0)

```
S -> (L) | a
L -> SM
M -> ,SM | ε
```
---
> Es LL(1) y LR(0)
```
S -> A
A -> aBD
D -> dD | ε
B -> b
C -> g
```
---
> No es ni LL(1) pero es LR(0)

```
S -> A
A -> aB | Ad
B -> b
C -> g
```
---
> Es LL(1) y LR(0)

```
S -> aBDh
B -> cC
C -> bC | ε
D -> EF
E -> g | ε
F -> f | ε
```
---
> Es LL(1) y LR(0)
```
E -> TG
G -> +TG | ε
T -> FU
U -> *FU | ε
F -> (E) | i
```
