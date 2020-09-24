namespace command {
  export const client = {
    ping: "ping",
    connect: "connect",
    getCollections: "getCollections",
    getCollection: "getCollection",
    requestCall: "requestCall"
  };

  export const server = {
    error: "error",

    pong: "pong",
    connected: "connected",

    collections: "collections",
    collectionDetail: "collectionDetail",

    callResult: "callResult"
  };
}
