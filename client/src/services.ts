namespace services {
  export interface Service {
    readonly key: string;
    readonly title: string;
    readonly plural: string;
    readonly icon: string;
  }

  export const system: Service = { key: "system", title: "System", plural: "systems", icon: "close" };
  export const collection: Service = { key: "collection", title: "Collection", plural: "Collections", icon: "folder" };
  export const request: Service = { key: "request", title: "Request", plural: "Requests", icon: "file-text" };

  const allServices = [system, collection];

  export function fromKey(key: string) {
    const ret = allServices.find(s => s.key === key);
    if (!ret) {
      throw `invalid service [${key}]`;
    }
    return ret;
  }
}
