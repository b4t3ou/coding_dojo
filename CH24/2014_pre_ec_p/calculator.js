var startTime = +new Date(),
    argv      = process.argv;

if (3 > argv.length)
{
    console.log('Usage: node <filename.js> <input-file-id> [-v]');
    return;
}

var verbose = !!argv[3] && '-v' === argv[3],
    fs      = require('fs');

fs.readFile('inputs/' + argv[2] + '.in', processInput);

function processInput(err, data)
{
    if (err) { throw err; }

    var lines      = data.toString('utf-8').split("\n"),
        numCinemas = +lines.shift(),
        distance   = 0;

    for (
        var i = 1,
            coordinates1 = lines[0],
            idx1 = coordinates1.indexOf(' '),
            lat1 = toRad(+coordinates1.substring(0, idx1)),
            lon1 = toRad(+coordinates1.substring(idx1)),
            coordinates2, idx2, lat2, lon2;
        i <= numCinemas;
        i++
    ) {
        coordinates2 = lines[i];

        if (numCinemas === i)
        {
            coordinates2 = lines[0];
        }

        idx2 = coordinates2.indexOf(' ');
        lat2 = toRad(+coordinates2.substring(0, idx2));
        lon2 = toRad(+coordinates2.substring(idx2));

        distance += sphericalLawOfCosines(lat1, lon1, lat2, lon2);

        coordinates1 = coordinates2;
        idx1 = idx2;
        lat1 = lat2;
        lon1 = lon2;
    }

    processOutput(distance);
}

function processOutput(data)
{
    fs.writeFile('outputs/p' + argv[2] + '.out', data, function(err)
    {
        if (err) { throw err; }

        console.log(data + "\n");

        console.log('Elapsed time: ' + (+new Date() - startTime) + ' ms');
        console.log((process.memoryUsage().heapTotal / 1024 / 1024).toFixed(2) + ' MiB');
    });
}


function sphericalLawOfCosines(lat1, lon1, lat2, lon2)
{
    return Math.acos(Math.sin(lat1) * Math.sin(lat2) +
        Math.cos(lat1) * Math.cos(lat2) *
        Math.cos(lon2 - lon1)) * 6371;
}

function toRad(n)
{
    return n * Math.PI / 180;
}