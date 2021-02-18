class MongoInterface {
  constructor() {}

    /**
   * Appends a document to a running mongoDB instance and handles the errors.
   * @param {any}  schemaDocument  - a declared mongo.schema variable with the appropriate parameters
   * @param {bool} logError        - whether to console log the error message
   * @param {bool} logSuccess      - whether to console log the success message
   */
  addDocument  = async (schemaDocument: any, logError = false, logSuccess = false) => {
    await schemaDocument
      .save() // save to db
      .then((res: any) => {
        logSuccess && console.log("Sucessfully added to mongo db: ", res);
        return;
      })
      .catch((error: Error) => {
        logError && console.log("Unsucessful document add: ", error);
        return;
      });
  };

    /**
   * Delete a document from a running mongoDB instance and handles the errors.
   * @param {string} deleteType      - the type of deletion. Defaults to findOneAndDelete
   * @param {any}    schema          - a mongoose Schema model
   * @param {Object} deleteParams     - the parameters to search for the document within the collection. 
   * @param {bool}   logError        - whether to console log the error message
   * @param {bool}   logSuccess      - whether to console log the success message
   */
  deleteDocument = async (deleteType = "findAndDelete", schema: any, deleteParams: Object, logError = false, logSuccess = false) => {
    switch(deleteType){
      case "deleteFirst":
        await schema.deleteOne(deleteParams)
        .then((result: any) => {
          if (result.deletedCount > 0){
          console.log("Successfuly deleted document")
        } else {
          console.log("Unsucessful deleted document. Check if document exists or that delete parameters match document items.");
        }})
        .catch((error: Error) => console.log("Unexpected error when deleting document:", error))
        break

      case "deleteMany":
        break
      
      case "findAndDelete":
        await schema.findOneAndDelete(deleteParams)
        .then((res: any) => {
          if (logSuccess && res != null){
            console.log("Sucessfully delete document: ", res);
          } else if (logError && res == null){
            console.log("Unsucessful deleted document. Check if document exists or that delete parameters match document items.");
          }
          return;
        })
        .catch((error: Error) => {
          logError && console.log("Unexpected error when deleting document:", error);
          return;
        });
        break

      default: 
      break
    }
  }
}

export default MongoInterface;
