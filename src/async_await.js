function wait(waitSec) {
  return new Promise(resolve => {
    setTimeout(_ => {
      resolve(`resolved:${waitSec}`);
    }, waitSec);
  });
}

async function awaitNumsWait(nums) {
  while ( (num = nums.shift()) !== undefined ) {
    const result = await wait(num);
    console.log(`await [${result}]`);
  }
}

async function asyncNumsWait(nums) {
  Promise.all(nums.map(n => {
    return wait(n);
  })).then(values => {
	return console.log('asyncDone::' +values.join(','));
  });
}

async function awaitAndAsyncCall(nums) {
  asyncNumsWait(nums);
  awaitNumsWait(nums);
  return 'awaitAndAsyncCall::done';
}

const nums = [2000, 1000, 1500];
awaitAndAsyncCall(nums).then((msg) => {
  console.log(msg);
});
