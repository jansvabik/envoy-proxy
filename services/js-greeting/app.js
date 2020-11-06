/**
 * js-greeting
 * Simple API to get random greeting
 */

const app = require('express')();
const config = require('./config');

/**
 * Get random greeting
 * @returns {Object} Structurized data about one random greeting
 */
const greet = () => {
    const list = [
        ['Hello', 'english', true],
        ['Hi', 'english', false],
        ['Dobrý den', 'czech', true],
        ['Ahoj', 'czech', false],
        ['Bonjour', 'french', true],
        ['Salut', 'french', false],
        ['Buenos días', 'spanish', true],
        ['Hola', 'spanish', false],
        ['Aloha', 'hawaiian', true],
        ['Hui', 'hawaiian', false],
        ['Halló', 'icelandic', false],
        ['Hæ', 'icelandic', false],
        ['Päivää', 'finnish', true],
        ['Hei', 'finnish', false],
        ['Szia', 'hungarian', true],
    ];
    
    // get random greeting and return it structurized
    const greeting = list[Math.floor(Math.random() * list.length)];
    return {
        text: greeting[0],
        language: greeting[1],
        formal: greeting[2],
    };
};

// the one and only route user can get
app.get('/', (req, res) => {
    res.json({
        status: 'OK',
        data: greet(),
    });
});

// start listening
app.listen(config.port, (err) => {
    if (err) {
        console.error(err);
        process.exit(1);
    }

    console.log(`Listening on port ${config.port}.`);
});
