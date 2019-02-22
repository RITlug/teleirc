/**
 * Does basic cleanup if the user does not have a username on telegram.
 * Replaces with a firstname instead.
 * @param {*} from - The object that contains the Telegram-side username
 *                   information.
 * @returns The username, if the user has one, else the first name.
 */
module.exports.ResolveUserName = function(from) {
    if(from.username === undefined) {
        return from.first_name;
    }
    
    // Add ZWP to remove IRC highlighting when Telegram + IRC nicks are the same
    return from.username.substr(0,1) + "\u200B" + from.username.substr(1);
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
