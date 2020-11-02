import {ref} from "@vue/composition-api";

export interface Breadcrumb {
  readonly title: string;
  readonly path: string;
}

export const breadcrumbsRef = ref<Breadcrumb[]>();
