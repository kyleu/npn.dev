export interface Collection {
  key: string;
  title: string;
  description?: string;
}

export interface CollectionCount {
  coll: Collection;
  count: number;
}
