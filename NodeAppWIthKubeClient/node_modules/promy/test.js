const test = require('ava');
const promy = require('.');

function asyncTest(ok, cb){
    setTimeout(() => ok ? cb(null, ok) : cb(new Error('error')), 10);
}

const promyTest = promy(asyncTest);

test.cb('Callback way: ok', t => {
    promyTest(true, (err, res) => {
        t.is(err, null);
        t.is(res, true);
        t.end();
    });
});

test.cb('Callback way: error', t => {
    promyTest(false, (err, res) => {
        t.is(err.message, 'error');
        t.is(res, undefined);
        t.end();
    });
});

test('Promise way', async t => {
    t.is(await promyTest(true), true);
    await t.throws(promyTest(false));
});
