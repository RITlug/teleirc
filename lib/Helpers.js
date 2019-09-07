/*
The MIT License (MIT)

Copyright (c) 2016 RIT Linux Users Group

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/



/**
 * Checks to see if an element exists in an array
 * that matches the given condition.
 * @param {Array} array - The array to check
 * @param {Function} predicate - The condition in the form
 *                               of a function,
 *                               whose argument is an object, and
 *                               returns true.
 * @returns True if an object that matches the condition exists
 *          in the array, else false.  Also returns false if the array
 *          is null or undefined.
 */
module.exports.Exists = function Exists(array, predicate) {
    if ((array === undefined) || (array === null))
    {
        return false;
    }

    for(let i = 0; i < array.length; ++i)
    {
        if (predicate(array[i]))
        {
            return true;
        }
    }

    return false;
}

/**
 * Checks to see if a string exists in an array, while
 * ignoring casing.
 * @param {Array} array - The array of strings to check.
 * @param {String} str - The string to check and see if it exists.
 * @returns true if the string exists in the array, else false.
 *          Also returns fale if the array is null or undefined.
 */
module.exports.StringExistsIgnoreCase = function StringExistsIgnoreCase(array, str) {
    if((str===undefined) || (str === null)) {
        return false;
    }

    return module.exports.Exists(
        array,
        (s) => { return str.toLowerCase() === s.toLowerCase(); }
    );
}

module.exports.IsNullOrUndefined = function IsNullOrUndefined(obj) {
    return (obj === undefined) || (obj === null);
}
