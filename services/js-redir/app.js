/**
 * js-redir
 * Simple API to be redirected to the specified URL
 */

const app = require('express')();
const config = require('./config');

// the one and only route user can get
app.get('/', (req, res) => {
    // target exist test
    if (!req.query.target) {
        return res.json({ status: "ERROR", msg: 'No target URL was specified. Add ?target=some.url&protocol=https' })
    }

    // determine redirect protocol
    let protocol = req.query.protocol || 'http';
    let allowedProtocols = ['http', 'https', 'ftp'];
    if (!allowedProtocols.includes(protocol)) {
        return res.json({ status: "ERROR", msg: 'Only protocols in the "allowedProtocols" field are allowed.', allowedProtocols })
    }

    // do the redirection
    res.redirect(protocol + '://' + req.query.target);
});

// start listening
app.listen(config.port, (err) => {
    if (err) {
        console.error(err);
        process.exit(1);
    }

    console.log(`Listening on port ${config.port}.`);
});
