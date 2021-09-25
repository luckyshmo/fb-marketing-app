const http = require('http');
const server = http.createServer();

const port = process.env.port || 8087;

server.on('request', (request, response) => {
    let body = [];
    request.on('data', (chunk) => {
        body.push(chunk);
    }).on('end', () => {
        body = Buffer.concat(body).toString();

	console.log(`==== ${request.method} ${request.url}`);
	console.log('> Headers');
        console.log(request.headers);

	console.log('> Body');
	console.log(body);
        response.end();
    });
}).listen(port, function(err) {
    if(err) throw err;
    console.log('Started on:' + port);
});