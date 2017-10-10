/**
 * Exception that gets thrown when there's an error.
 */
class TeleIrcException {

  // ---------------- Constructor ----------------

  /**
   * Constructor.
   * @param errorCode - Unique number for this error.
   * @param message - Corresponding Message.
   */
  constructor(errorCode, message) {
    this.errorCode = errorCode;
    this.message = message;
  }

  toString() {
    return "TeleIrc Error " + this.errorCode + ": " + this.message;
  }
}

module.exports = TeleIrcException;
