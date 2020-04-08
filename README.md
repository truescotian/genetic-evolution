Adapted from https://natureofcode.com/book/chapter-9-the-evolution-of-code/ which uses Processing.

# Improvements
There are optimizations for this (need more variability/improve heredity) such as:

1. Use less ASCII characters.
2. Crossover to give each gene 50% chance of coming from parentA and 50% chance of coming from parentB
3. Replace calculated fitness scores with ordinals of scoring (their rank)
4. When retrieving two parents for mating, ensure they are not the same parent
