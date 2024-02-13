const puppeteer = require('puppeteer');

module.exports = {
    youtube: async (search) => {
        const browser = await puppeteer.launch();
        const page = await browser.newPage();
    
        await page.goto(`https://youtube.com/results?search_query=${search}`);
    
        const links = await page.$$eval('a', links => links.map(link => link.href));
    
        const uniqueLinks = [];
    
        links.forEach(link => {
            if (link.includes('watch') && uniqueLinks.indexOf(link) === -1) {
                uniqueLinks.push(link);
            }
        });
    
        await browser.close();
        return uniqueLinks;
    }
}
