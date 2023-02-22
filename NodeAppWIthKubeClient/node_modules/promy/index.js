module.exports = fn => (obj, cb) => {
    if(cb && cb instanceof Function){
        return fn(obj, cb);
    }
    return new Promise((resolve, reject) => {
        fn(obj, (err, res) => {
            if(err){
                reject(err);
            } else {
                resolve(res);
            }
        });
    });
};
