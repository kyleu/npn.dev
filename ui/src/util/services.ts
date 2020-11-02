export interface Service {
  readonly key: string;
  readonly title: string;
  readonly plural: string;
  readonly icon: string;
}

export const systemService: Service = { key: "system", title: "System", plural: "systems", icon: "close" };
export const collectionService: Service = { key: "collection", title: "Collection", plural: "Collections", icon: "folder" };
export const requestService: Service = { key: "request", title: "Request", plural: "Requests", icon: "file-text" };
