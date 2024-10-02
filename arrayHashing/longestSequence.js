nums = [2, 20, 4, 10, 3, 4, 5];

// sort nums = [2, 3, 4, 5, 10, 20]

// how to find consecutive sequence?

function longestConsecutive(nums) {
    const graph = new Map();
    const stack = [];
    for (let number of nums) {
        graph.set(number, [number - 1, number + 1]);
    }

    for (let value of graph) {
        while (stack.length > 0) {}
    }
}

// TODO: instead of going to the solution, from solution find the process
function solution(nums) {
    const set = new Set(nums);
    let longest = 0;
    for (let n of set) {
        const neighbor = n - 1;
        if (!set.has(neighbor)) {
            // check if n+1 exists
            let length = 1;
            while (set.has(n + length)) {
                length++;
            }
            longest = Math.max(length, longest);
        }
    }
    console.log(longest);
}
solution(nums);
