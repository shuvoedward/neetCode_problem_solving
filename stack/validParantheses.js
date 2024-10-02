function isValid(s) {
    if (s.trim().length === 0) {
        return false;
    }
    const stack = [];
    const pairs = new Map([
        [")", "("],
        ["}", "{"],
        ["]", "["],
    ]);

    for (let c of s) {
        if (pairs.has(c)) {
            if (stack.pop() !== pairs.get(c)) return false;
        } else {
            stack.push(c);
        }
    }
    return true;
}

const s = "([{}])";

const r = "([}])";
console.log(isValid(s)); // true

console.log(isValid(r)); // false

console.log(isValid("()[]{}")); // true
console.log(isValid("([])")); // true
console.log(isValid("(]")); // false
console.log(isValid("     ")); // false
