function isValid(s) {
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
console.log(isValid(s));

console.log(isValid(r));

console.log(isValid("()[]{}"));
console.log(isValid("([])"));
console.log(isValid("(]"));
