const fs = require('fs');

// advent of code, 1 Dec: https://adventofcode.com/2018/day/1
fs.readFile('text.txt', 'utf8', function(err, contents) {
  if (err) throw err;
  let arr = contents.split('\n');
  arr = arr.map(val => parseInt(val));
  const accumulation = arr.reduce((i, j) => i + j);
  console.log('Assignment 1, Accumalation', accumulation);

  let calibrationSecond = undefined;
  let set = new Set();
  let sum = 0;
  while (calibrationSecond === undefined) {
    arr.forEach(val => {
      sum += val;
      if (set.has(sum) && calibrationSecond === undefined) {
        calibrationSecond = sum;
      } else {
        set.add(sum);
      }
    });
  }
  console.log('Assignment 2, Calibration', calibrationSecond);
});
