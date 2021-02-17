const mongoose = require("mongoose");

/**
 *  Mongo schema to define type: Room.
 *  A document (channel) ID is automatically generated
 */
const ChannelSchema = mongoose.Schema({
  name: {
    type: String,
    required: false,
    default: "New Channel",
  },
  private: {
    type: Boolean,
    required: false,
    default: true,
  },
  users: {
    type: Array,
    required: true,
    default: [],
  },
});

module.exports = mongoose.model("Channel", ChannelSchema);
