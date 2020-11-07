/**
 * js-web
 * Very simple website running on Node.js
 */

const app = require('express')();
const config = require('./config');

// set up pug as a template engine
app.set('view engine', 'pug')

// routes
app.get('/', (req, res) => res.render('homepage'));
app.get('/about', (req, res) => res.render('about'));
app.get('/contact', (req, res) => res.render('contact'));

// start listening
app.listen(config.port, (err) => {
    if (err) {
        console.error(err);
        process.exit(1);
    }

    console.log(`Listening on port ${config.port}.`);
});
