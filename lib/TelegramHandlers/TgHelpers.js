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
    // See https://github.com/42wim/matterbridge/issues/175 for inspiration
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

    // Add ZWP so users with same username across bridge don't ping themselves
    username = from.username.substr(0,1) + "\u200B" + from.username.substr(1);

    return from.first_name + " (@" + username + ")";
};
