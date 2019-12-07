#include <stdio.h>
#include <stdlib.h>
#include <string.h>

void getSolution(int part) {
	char stringEq[7];
	int numberOfPasswords = 0;
	int isAscending = 1; //default to true
	int hasDouble = 0; //default to false
	for (int x = 278384; x < 824795; x++) {
		sprintf(stringEq, "%d", x);
		// reset flags
		isAscending = 1;
		hasDouble = 0;
		for (int i = 0; i < 6; i++) {
			// get part 1 solution if specified by param
			if (part == 1) {
				// check if number at current index is greater than the last
				// if we're past the first element
				if (i > 0) {
					if (stringEq[i] < stringEq[i-1]) {
						isAscending = 0;
					} else if (stringEq[i] == stringEq[i-1]) {
						hasDouble = 1;
					}
				}
			// get part 2 solution if specified by param
			} else if (part == 2) {
				if (i > 0) {
					if (stringEq[i] < stringEq[i-1]) {
						isAscending = 0;
					}

					// if iterating over index 1, simply check previous and next elem
					if (i == 1) {
						if (stringEq[i] == stringEq[i-1] && stringEq[i] != stringEq[i+1]) {
							hasDouble = 1;
						}
					// if in middle of array, must check previous two and next elem
					} else if (i > 1 && i < 5) {
						if (stringEq[i] == stringEq[i-1] && stringEq[i] != stringEq[i-2] && stringEq[i] != stringEq[i+1]) {
							hasDouble = 1;
						}
					// if at the last elem, only check last 2 elems
					} else if (i == 5) {
						if (stringEq[i] == stringEq[i-1] && stringEq[i] != stringEq[i-2]) {
							hasDouble = 1;
						}
					}
				}
			}
		}
		if (isAscending && hasDouble) {
			numberOfPasswords++;
		}
	}
	if (part == 1)
		printf("Part 1 solution: %d\n", numberOfPasswords);
	else if (part == 2)
		printf("Part 2 solution: %d\n", numberOfPasswords);
}

void main() {
	getSolution(1);
	getSolution(2);
}
