const fs = require('fs');

//Jonathan Burdette
//Text Cloud Implementation
//I used an answer from this on line 45 https://stackoverflow.com/questions/37982476/how-to-sort-a-map-by-value-in-javascript

if(process.argv.length < 5) {
    console.log("Invalid syntax. To run program, enter: \"node textCloud.js <input file name> <exclude file name> <output file name>\"");
} else {
    var inputFile = process.argv[2];
    var excludeFile = process.argv[3];
    var outputFile = process.argv[4];

    //store words from exclude file in map (blank values)
    var excludeMap = new Map();
    fs.readFile(excludeFile, function(err, lines) {
        if(err) throw err;

        var wordArr = lines.toString().split(/[^A-Z'a-z]/);
        for(var i in wordArr) {
            var word = wordArr[i].toLowerCase();
            excludeMap.set(word, "");
        }
    })

    //store the words and their counts from input file
    var inputMap = new Map();
    fs.readFile(inputFile, function(err, lines) {
        if(err) throw err;

        var wordArr = lines.toString().split(/[^A-Z'a-z]/);
        for(var i in wordArr) {
            var word = wordArr[i].toLowerCase();
            if(!excludeMap.has(word) && word.match(/[A-Za-z]/) && word.length > 1) {
                var count = inputMap.get(word);
                if(count == null) {
                    inputMap.set(word, 1);
                } else {
                    inputMap.set(word, count + 1);
                }
            }
        }

        //sort the words based on count
        var sortedCounts = [...inputMap.entries()].sort(function(a, b){return b[1] - a[1]});

        //take the top fifty words based on count
        var topFifty = sortedCounts.slice(0, 50);

        //calculate info for html file based on top fifty words
        var range = (topFifty[0][1] - topFifty[49][1]);
        var sizeFactor = (1000 / range);

        //sort the top fifty in alphabetical order
        var sortedWords = topFifty.sort();

        //loop through the sorted words to create a string with html info
        var colors = ["#623224", "#FABE8B", "#EC1D23", "#0160B0"];
        var count = 0;
        var htmlLines = "";
        for(var i in sortedWords) {
            if(count == 4) { //colors restart after 4 iterations
                count = 0;
            }
            var fontSize = sortedWords[i][1] * sizeFactor;
            htmlLines = htmlLines + "<span style=\"font-size:"+fontSize+"%; color:"+colors[count]+";\">"+sortedWords[i][0]+"</span> &nbsp; &nbsp;\n";
            count++;
        }

        //write html file
        fs.writeFile(outputFile, htmlLines, function(err) {
            if(err) throw err;
        })
    })
}
