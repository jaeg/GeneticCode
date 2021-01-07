# Genetic Code


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
