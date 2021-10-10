const http = require('http');
const serveHandler = require('serve-handler');
const server = http.createServer();

const port = process.env.port || 8087;

server.on('request', (request, response) => {
   
    response.writeHead(200, {
        'Content-Type': 'application/json',
        'Access-Control-Allow-Origin': '*',
        'Access-Control-Allow-Headers': 'Origin, X-Requested-With, Content-Type, Accept'
    });

    let body = [];
    request.on('data', (chunk) => {
        body.push(chunk);
    }).on('end', () => {
        body = Buffer.concat(body).toString();

	console.log(`==== ${request.method} ${request.url}`);
	// console.log('> Headers');
    //     console.log(request.headers);

	console.log('> Body');
	console.log(body);
        response.end();
    });
}).listen(port, function(err) {
    if(err) throw err;
    console.log('Started on:' + port);
});

const serverStatic = http.createServer((req, res) => serveHandler(req, res, {
    public: './dist' // folder of files to serve
  }));
serverStatic.listen(8000, function(err) {
    if(err) throw err;
    console.log('Started on:' + 8000);
});
