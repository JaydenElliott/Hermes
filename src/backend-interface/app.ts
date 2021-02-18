import MongoInterface from "./database/mongoInterface";
import * as mongoose from "mongoose";
import Channel from "./database/models/channelSchema";

/** Mongoose Server for testing */
mongoose.connect(
  "mongodb://localhost:27017/arcstackTesting", //where the endpoint is the db name
  {
    useNewUrlParser: true,
    useUnifiedTopology: true,
    useCreateIndex: true,
  },
  () => {
    console.log("Connected to DB!");
  }
);

let test = new MongoInterface();

const channel1 = new Channel({ users: ["John"] });
console.log("test");
// test.addDocument(channel1);
test.deleteDocument("deleteFirst", Channel, {name: "abc"}, true, true)