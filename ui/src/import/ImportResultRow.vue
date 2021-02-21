<template>
  <tr>
    <td>{{ file.filename }}</td>
    <td>{{ file.type }}</td>

    <td v-if="file.value.coll" :title="file.value.coll.key">{{ file.value.coll.title }}</td>
    <td v-else>???</td>

    <td v-if="file.value.reqs">{{ file.value.reqs.length }}</td>
    <td v-else>???</td>

    <td v-if="file.value.sess.cookies">{{ file.value.sess.cookies.length }}</td>
    <td v-else>0</td>

    <td v-if="file.value.sess.variables">{{ file.value.sess.variables.length }}</td>
    <td v-else>0</td>
  </tr>
</template>

<script lang="ts">
import {Component, Prop, Vue} from "vue-property-decorator";
import Icon from "@/util/Icon.vue";
import {ImportResultFile} from "@/import/model";

@Component({ components: { Icon } })
export default class ImportResultRow extends Vue {
  @Prop() file!: ImportResultFile;
}
</script>
