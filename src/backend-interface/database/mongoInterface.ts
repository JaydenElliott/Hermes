class MongoInterface {
  constructor() {}

  /**
   * Appends a document to a running mongoDB instance and handles the errors.
   * @param {any}  schemaVariable  - a declared mongo.schema variable with the appropriate parameters
   * @param {bool} logError        - whether to console log the error message
   * @param {bool} logSuccess      - whether to console log the success message
   */
  addDocument = async (schemaVariable: any, logError = false, logSuccess) => {
    await schemaVariable
      .save() // save to db
      .then((res: any) => {
        logError && console.log("Sucessfully added to mongo db: ", res);
        return;
      })
      .catch((error: Error) => {
        logSuccess && console.log("Unsucessful document add: ", error);
        return;
      });
  };
}

export default MongoInterface;
