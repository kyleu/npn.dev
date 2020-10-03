namespace command {
  export const client = {
    ping: "ping",
    connect: "connect",

    // Collection
    getCollection: "getCollection",
    addURL: "addURL",

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

    requestDetail: "requestDetail",
    callResult: "callResult",
    transformResult: "transformResult"
  };
}
