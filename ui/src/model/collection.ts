export interface Collection {
  key: string;
  title: string;
  description: string;
}

export const MockCollections: Collection[] = [
  {
    key: "test1",
    title: "Test 1",
    description: "Test Collection 1"
  },
  {
    key: "test2",
    title: "Test 2",
    description: "Test Collection 2"
  }
]
