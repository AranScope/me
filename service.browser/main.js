const puppeteer = require('puppeteer');
const express = require('express');
const bodyParser = require('body-parser');
const app = express();
app.use(bodyParser.json());

const port = 3000;

const functions = {
    "open": async (page, params) => {
        await page.goto(params.url);
        return r => r;
    },
    "input": async (page, params) => {
        await page.waitForSelector(params.selector);
        await page.type(params.selector, params.text);
        return r => r;
    },
    "submit": async (page, params) => {
        await page.waitForSelector(params.selector);
        await Promise.all([
            page.click(params.selector),
            page.waitForNavigation({waitUntil: 'networkidle2'})
        ]);
        return r => r;
    },
    "save-text": async (page, params) => {
        await page.waitForSelector(params.selector);
        const text = await page.$eval(params.selector, e => e.innerText);
        return r => {
            r[params.response_key] = text;
            return r;
        }
    }
};

const runSteps = async (page, steps) => {
    let response = {};

    for(let step of steps) {
        let page_func = functions[step.function];
        let mutate_response = await page_func(page, step.params);
        response = mutate_response(response);
    }

    return response;
};

app.post('/scrape', (req, res) => {

    const body = req.body;
    (async () => {

        const browser = await puppeteer.launch({headless: true, args: [
                '--no-sandbox',
                '--disable-setuid-sandbox',
            ]});
        const page = await browser.newPage();
        await page.setRequestInterception(true);
        page.on('request', (request) => {
            if (['image', 'stylesheet', 'font'].indexOf(request.resourceType()) !== -1) {
                request.abort();
            } else {
                request.continue();
            }
        });

        const response = await runSteps(page, body.steps);
        res.send(response);
        await browser.close();
    })();
});

app.listen(port, () => {
    console.log(`app listening on port ${port}`);
});
