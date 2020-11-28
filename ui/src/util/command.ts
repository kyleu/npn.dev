export const clientCommands = {
  testbed: "testbed",
  saveProfile: "saveProfile",
  runURL: "runURL",

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
  transform: "transform",

  // Session
  getSession: "getSession",
  addSession: "addSession",
  saveSession: "saveSession",
  deleteSession: "deleteSession"
};

export const serverCommands = {
  log: "log",

  // Collection
  collections: "collections",
  collectionDetail: "collectionDetail",
  collectionNotFound: "collectionNotFound",
  collectionAdded: "collectionAdded",
  collectionUpdated: "collectionUpdated",
  collectionDeleted: "collectionDeleted",

  // Request
  requestDetail: "requestDetail",
  requestNotFound: "requestNotFound",
  requestAdded: "requestAdded",
  requestDeleted: "requestDeleted",
  callResult: "callResult",
  transformResult: "transformResult",

  // Session
  sessions: "sessions",
  sessionAdded: "sessionAdded",
  sessionDetail: "sessionDetail",
  sessionDeleted: "sessionDeleted",
  sessionNotFound: "sessionNotFound"
};
