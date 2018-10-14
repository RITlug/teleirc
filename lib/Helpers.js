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
