import {ref} from "@vue/composition-api";

export interface RequestTransformResult {
  coll: string;
  req:  string;
  fmt:  string;
  out:  string;
}

export interface CollectionTransformResult {
  coll: string;
  fmt:  string;
  out:  string;
}

export interface SessionTransformResult {
  key: string;
  out:  string;
}
