export const clientCommands = {
  ping: "ping",
  connect: "connect",

  // Collection
  getCollection: "getCollection",
  addCollection: "addCollection",
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
  collectionAdded: "collectionAdded",
  collectionDeleted: "collectionDeleted",

  requestDetail: "requestDetail",
  requestAdded: "requestAdded",
  requestDeleted: "requestDeleted",
  callResult: "callResult",
  transformResult: "transformResult"
};
