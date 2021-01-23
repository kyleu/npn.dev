import {ref} from "@vue/composition-api";
import {socketRef} from "@/socket/socket";
import {systemService} from "@/util/services";
import {clientCommands} from "@/util/command";

export interface SearchResult {
  readonly t: string;
  readonly key: string;
  readonly prelude: string;
  readonly postlude: string;
  readonly loc: string;
}

class SearchQuery {
  q = "";
  o = 0;
  loaded = false;
  results: SearchResult[] = [];
}

export const searchQueryRef = ref<SearchQuery>(new SearchQuery());
export const searchResultsRef = ref<SearchResult[]>([]);

export function onSearchQuery(q: string, o: number): void {
  if (q !== searchQueryRef.value.q || o !== searchQueryRef.value.o) {
    searchQueryRef.value.q = q;
    searchQueryRef.value.o = o;

    if (q.length > 0 && socketRef.value) {
      socketRef.value.send({channel: systemService.key, cmd: clientCommands.search, param: {q, o}});
    }
  }
}

export function onSearchResults(sr: SearchResult[]): void {
  searchResultsRef.value = sr;
}
