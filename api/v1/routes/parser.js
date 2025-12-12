const { console } = require('console');
const { resolve } = require('path');
const fs = require('fs');
const xml2js = require('xml2js');

const Parser = function (filePath) {
  this.filePath = filePath;
};

Parser.prototype.parse = function (callback) {
  const parser = new xml2js.Parser({ explicitArray: false });
  fs.readFile(this.filePath, (err, data) => {
    if (err) {
      console.error(err);
      return;
    }
    parser.parseString(data, (err, result) => {
      if (err) {
        console.error(err);
        return;
      }
      callback(null, result);
    });
  });
};

module.exports = Parser;