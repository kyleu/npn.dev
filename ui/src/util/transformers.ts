// Request
export interface RequestTransformer {
  key: string;
  title: string;
  description: string;
}

const requestTestbed = {key: "testbed", title: "Testbed", description: "an internal transformer for testing request exports"};

export const RequestTransformers: RequestTransformer[] = [
  {key: "http", title: "HTTP", description: "exports this request as HTTP text"},
  {key: "json", title: "JSON", description: "exports this request in its native format"},
  {key: "curl", title: "curl", description: "exports this request in a format usable by curl"},
  {key: "postman", title: "Postman", description: "exports this request as a Postman collection"},
  requestTestbed
];

export function getRequestTransformer(fmt: string | undefined): RequestTransformer {
  if (fmt === undefined) {
    fmt = "testbed";
  }
  return RequestTransformers.find(x => x.key === fmt) || requestTestbed;
}

// Collection
export interface CollectionTransformer {
  key: string;
  title: string;
  description: string;
}

const collectionTestbed = {key: "testbed", title: "Testbed", description: "an internal transformer for testing collection exports"};

export const CollectionTransformers: CollectionTransformer[] = [
  {key: "json", title: "JSON", description: "exports this collection in its native format"},
  {key: "postman", title: "Postman", description: "exports all requests in this collection as a Postman collection"},
  collectionTestbed
];

export function getCollectionTransformer(fmt: string | undefined): CollectionTransformer {
  if (fmt === undefined) {
    fmt = "testbed";
  }
  return CollectionTransformers.find(x => x.key === fmt) || collectionTestbed;
}
