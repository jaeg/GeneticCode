# Genetic Code

## Genes
- Represents a token in a program.  This could range from "if" or "+" to the number 3 or 6.  Each type of gene inherits from a Gene interface and has its own Eval function.  If a gene requires more genes to run its Eval it'll call the Eval function of the gene after it.  This gene could in turn call others' eval function.  

## Computer
- Represents the computer running the program encoded by a creatures genes. 
- Starting at an index of 0 it'll eval each line of code using the returened index from "Eval" to know where to find the next line of code to execute.
## Creature

## Simulation


Code example in Genecode:
1.  if < 1 2  - evals as true
2.  if < + 1 2 1 - evals as false

Easier to read form:
1. if 1 < 2
2. if 1 + 2 < 1

Loop through genes.
Call eval on gene passing all the genes and the current index.
	- If gene needs other genes, increment the index and call eval, increment index (which can change based on the next eval call) call eval.
	- Do this until we run out of genes.
