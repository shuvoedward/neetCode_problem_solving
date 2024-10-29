function search(nums, target) {
    let l = 0,
        r = nums.length - 1;

    while (l <= r) {
        let mid = Math.floor((l + r) / 2);
        if (nums[mid] === target) {
            return mid;
        } else if (nums[l] <= target && nums[mid] > target) {
            r = mid - 1;
        } else {
            l = mid + 1;
        }
    }
    return -1;
}

console.log(search([3, 4, 5, 6, 1, 2], 1)); // 4
console.log(search([3, 5, 6, 0, 1, 2], 1)); // 4
console.log(search([5, 1, 3], 5)); // 0
