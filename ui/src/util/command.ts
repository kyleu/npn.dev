export const clientCommands = {
  ping: "ping",
  connect: "connect",
  testbed: "testbed",

  // Collection
  getCollection: "getCollection",
  addCollection: "addCollection",
  saveCollection: "saveCollection",
  deleteCollection: "deleteCollection",
  addRequestURL: "addRequestURL",

  // Request
  getRequest: "getRequest",
  saveRequest: "saveRequest",
  deleteRequest: "deleteRequest",
  call: "call",
  transform: "transform"
};

export const serverCommands = {
  error: "error",

  pong: "pong",
  connected: "connected",

  collections: "collections",
  collectionDetail: "collectionDetail",
  collectionNotFound: "collectionNotFound",
  collectionAdded: "collectionAdded",
  collectionUpdated: "collectionUpdated",
  collectionDeleted: "collectionDeleted",

  requestDetail: "requestDetail",
  requestNotFound: "requestNotFound",
  requestAdded: "requestAdded",
  requestDeleted: "requestDeleted",
  callResult: "callResult",
  transformResult: "transformResult"
};
