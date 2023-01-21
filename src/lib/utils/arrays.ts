export const getRandomElements = (arr: any[], n: number) => {
  var result = new Array(n);
  var len = arr.length;
  var taken = new Array(len);

  while (n--) {
    var x = Math.floor(Math.random() * len);
    result[n] = arr[x in taken ? taken[x] : x];
    taken[x] = --len in taken ? taken[len] : len;
  }

  return result;
};

export const shuffleElements = (arr: any[]) => {
  let currentIndex = arr.length, randomIndex;

  while (currentIndex != 0) {
    randomIndex = Math.floor(Math.random() * currentIndex);
    currentIndex--;

    [arr[currentIndex], arr[randomIndex]] = [arr[randomIndex], arr[currentIndex]];
  }

  return arr
}