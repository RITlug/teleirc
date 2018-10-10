/**
 * Does basic cleanup if the user does not have a username on telegram.
 * Replaces with a firstname instead.
 * @param {*} from - The object that contains the Telegram-side username
 *                   information.
 * @returns The username, if the user has one, else the first name.
 */
module.exports.CleanUpUserName = function(from) {
    if(from.username === undefined) {
        return from.first_name;
    }
    
    return from.username;
};

/**
 * @param {*} from - The object that contains the Telegram-side username
 *                   information.
 * @returns The user's first name if the user does not have a username.
 *          Otherwise, it returns 'firstName (@userName)'.
 */
module.exports.GetFullUserName = function(from) {
    if (from.username === undefined) {
        return from.first_name;
    }

    return from.first_name + " (@" + from.username + ")";
};