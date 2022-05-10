// Disallowed: a list of infinite length with all elements being 1.
#List: {
	head: 1
	tail: #List
}

// Disallowed: another infinite structure (a:{b:{d:{b:{d:{...}}}}}, ...).
a: {
	b: c
}
c: {
	d: a
}

// #List defines a list of arbitrary length. Because the recursive reference
// is part of a disjunction, this does not result in a structural cycle.
#List: {
	head: _
	tail: null | #List
}

// Usage of #List. The value of tail in the most deeply nested element will
// be `null`: as the value of the disjunct referring to list is the only
// conjunct, all conjuncts are cyclic and the value is invalid and so
// eliminated from the disjunction.
MyList: #List & {head: 1, tail: {head: 2}}
