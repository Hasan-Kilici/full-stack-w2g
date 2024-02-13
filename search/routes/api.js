const search = require("../utils/request")

async function routes (fastify, options) {
    fastify.get('/search/youtube', async (request, reply) => {
        return { result: await search.youtube(request.query.s)}
    })
}

module.exports = routes