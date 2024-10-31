function maxProfit(prices) {
    let maxProfit = 0;
    let minimum = prices[0];

    for (let i = 1; i < prices.length; i++) {
        minimum = Math.min(minimum, prices[i]);
        maxProfit = Math.max(maxProfit, prices[i] - minimum);
    }

    return maxProfit;
}

function maxProfit2(prices) {
    let minPrice = Infinity;
    let maxProfit = 0;

    for (let price of prices) {
        if (price < minPrice) {
            minPrice = price;
        } else if (price - minPrice > maxProfit) {
            maxProfit = price - minPrice;
        }
    }
    return maxProfit;
}

function slidingWindow(prices) {
    // From neetcode
    let min = prices[0];
    let max = min;
    let value = 0;
    for (let i = 0; i < prices.length; i++) {
        if (i != prices.length - 1 && prices[i] <= min) {
            max = min = prices[i];
        } else if (prices[i] > max) {
            max = prices[i];
        }
        value = max - min > value ? max - min : value;
    }
    return value;
}

console.log(maxProfit([10, 1, 5, 6, 7, 1]));

console.log(maxProfit([10, 8, 7, 5, 2]));
console.log(maxProfit([7, 1, 5, 3, 6, 4]));
