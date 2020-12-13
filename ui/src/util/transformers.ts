// Request
export interface RequestTransformer {
  key: string;
  title: string;
  description: string;
}

const requestTestbed = {key: "testbed", title: "Testbed", description: "an internal transformer for testing request exports"};

export const RequestTransformers: RequestTransformer[] = [
  // start of generated request code
  {"key": "curl", "title": "CURL", "description": "TODO: curl"},
  {"key": "http", "title": "HTTP", "description": "TODO: http"},
  {"key": "json", "title": "JSON", "description": "TODO: json"},
  {"key": "openapi", "title": "OpenAPI", "description": "TODO: OpenAPI"},
  {"key": "postman", "title": "Postman", "description": "TODO: postman"},
  // end of generated request code
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
  // start of generated collection code
  {"key": "json", "title": "JSON", "description": "TODO: json"},
  {"key": "openapi", "title": "OpenAPI", "description": "TODO: OpenAPI"},
  {"key": "postman", "title": "Postman", "description": "TODO: postman"},
  // end of generated collection code
  collectionTestbed
];

export function getCollectionTransformer(fmt: string | undefined): CollectionTransformer {
  if (fmt === undefined) {
    fmt = "testbed";
  }
  return CollectionTransformers.find(x => x.key === fmt) || collectionTestbed;
}
