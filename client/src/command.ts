namespace command {
  export const client = {
    ping: "ping",
    connect: "connect",

    // Collection
    getCollection: "getCollection",
    addCollection: "addCollection",
    addRequestURL: "addRequestURL",

    // Request
    getRequest: "getRequest",
    call: "call",
    transform: "transform"
  };

  export const server = {
    error: "error",

    pong: "pong",
    connected: "connected",

    collections: "collections",
    collectionDetail: "collectionDetail",
    collectionAdded: "collectionAdded",

    requestDetail: "requestDetail",
    requestAdded: "requestAdded",
    callResult: "callResult",
    transformResult: "transformResult"
  };
}
