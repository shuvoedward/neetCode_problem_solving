/**
 * Duplicate Integer
Given an integer array nums, return true if any value appears more than once in the array, otherwise return false.

Example 1:

Input: nums = [1, 2, 3, 3]

Output: true
Example 2:

Input: nums = [1, 2, 3, 4]

Output: false
 */

// TODO: Using set to keep track of all the elements
function hasDuplicates(nums) {
    const track = new Set();

    for (let num of nums) {
        if (track.has(num)) return true;
        track.add(num);
    }
    return false;
    // O(n), set.has() Time Complexity
    // Space complexity  O(n), creating a new set. in the worst case size of set will equal to the size of the array
}

const nums = [1, 2, 3, 3];
console.log(hasDuplicates(nums));
