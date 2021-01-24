export interface Service {
  readonly key: string;
  readonly title: string;
  readonly plural: string;
}

export const systemService: Service = {
  key: "system",
  title: "System",
  plural: "systems"
};

export const sessionService: Service = {
  key: "session",
  title: "Session",
  plural: "Sessions"
};

export const collectionService: Service = {
  key: "collection",
  title: "Collection",
  plural: "Collections"
};

export const requestService: Service = {
  key: "request",
  title: "Request",
  plural: "Requests"
};

export const importService: Service = {
  key: "import",
  title: "Import",
  plural: "imports"
};
