namespace services {
  export interface Service {
    readonly key: string;
    readonly title: string;
    readonly plural: string;
    readonly icon: string;
  }

  export const system: Service = { key: "system", title: "System", plural: "systems", icon: "close" };

  const allServices = [system];

  export function fromKey(key: string) {
    const ret = allServices.find(s => s.key === key);
    if (!ret) {
      throw `invalid service [${key}]`;
    }
    return ret;
  }
}
