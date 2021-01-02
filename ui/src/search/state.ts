import {ref} from "@vue/composition-api";

interface SearchResult {
  readonly msg: string;
}

export const searchResultsRef = ref<SearchResult[]>([]);

export function onSearchResults(sr: SearchResult[]): void {
  searchResultsRef.value = sr;
}
