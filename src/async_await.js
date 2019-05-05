function wait(waitSec) {
  return new Promise(resolve => {
    setTimeout(() => {
      resolve('resolved');
    }, waitSec);
  });
}

async function asyncCall(num) {
  console.log('calling:' + num);
  wait(num).then((value) => {
    // expected output: 'resolved'
    console.log(`async [${num}]` + value);
  });
}

async function awaitAndAsyncCall(nums) {

  while ( (num = nums.shift()) !== undefined ) {
    console.log('await:' + num);
    const result = await wait(num);
    console.log(`await [${num}]` + result);
  }

  nums.forEach((n) => {
    // async
    asyncCall(n);
  });

  return await wait(3000);
}

const nums = [2000, 1000, 1500];
awaitAndAsyncCall(nums).then((msg) => {
  console.log(msg);
});
