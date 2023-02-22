# promy

Conditional promisify. Transform callback-based function to callback-and-promise-based one.

[![Build Status][travis-image]][travis-url]
[![NPM version][npm-image]][npm-url]


## Install

```bash
npm install promy
```

## Usage

If you have some async function getting calback as second parameter you can do something like this:

```js
const promy = require('promy');
const fn = promy(
    require('./someAsyncCallbackBasedFunction')
);
````

Now promisified function can be used in two ways. It can return promise if called with only one first argument, or it can run callback if it given as second argument.

```js
// callback way:

fn(arg, (err, res) => {
    console.log(err ? err : res);
});


//promise way:

fn(arg)
    .then((res) => console.log(res))
    .catch((err) => console.log(err));
```

## License

MIT

[npm-url]: https://npmjs.org/package/promy
[npm-image]: https://badge.fury.io/js/promy.svg
[travis-url]: https://travis-ci.org/astur/promy
[travis-image]: https://travis-ci.org/astur/promy.svg?branch=master